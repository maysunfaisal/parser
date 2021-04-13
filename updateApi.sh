#!/bin/bash

BLUE='\033[1;34m'
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'
BOLD='\033[1m'

set -e

CURRENT_DIR=$(pwd)
API_PKG="github.com/maysunfaisal/api/v2"
SCHEMA_URL_MASTER="https://raw.githubusercontent.com/devfile/api/main/schemas/latest/devfile.json"

# 2.0.0 devfile
SCHEMA_URL_200="https://raw.githubusercontent.com/devfile/api/2.0.x/schemas/latest/devfile.json"
PACKAGE_VERSION_200="version200"
JSON_SCHEMA_200="JsonSchema200"
FILE_PATH_200="./pkg/devfile/parser/data/v2/2.0.0/devfileJsonSchema200.go"

# 2.1.0 devfile
PACKAGE_VERSION_210="version210"
JSON_SCHEMA_210="JsonSchema210"
FILE_PATH_210="./pkg/devfile/parser/data/v2/2.1.0/devfileJsonSchema210.go"


onError() {
  cd "${CURRENT_DIR}"
}
trap 'onError' ERR


echo -e "${GREEN}Updating devfile/api in go.mod${NC}"
go get "${API_PKG}@main"

echo -e "${GREEN}Get latest schema${NC}"

case "${1}" in
   "2.0.0")
     SCHEMA_URL=${SCHEMA_URL_200}
     PACKAGE_VERSION=${PACKAGE_VERSION_200}
     JSON_SCHEMA=${JSON_SCHEMA_200}
     FILE_PATH=${FILE_PATH_200}
   ;;
   *)
     # default
     SCHEMA_URL=${SCHEMA_URL_MASTER}
     PACKAGE_VERSION=${PACKAGE_VERSION_210}
     JSON_SCHEMA=${JSON_SCHEMA_210}
     FILE_PATH=${FILE_PATH_210}
   ;;
esac

schema=$(curl -L "${SCHEMA_URL}")

#replace all ` with ' and write to schema file
echo -e "${GREEN}Write to go file${NC}"
go build *.go
./main updateSchema "${schema}" "${SCHEMA_URL}" "${PACKAGE_VERSION}" "${JSON_SCHEMA}" "${FILE_PATH}"
