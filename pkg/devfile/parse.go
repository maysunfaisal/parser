package devfile

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/devfile/library/pkg/devfile/parser"
	"github.com/devfile/library/pkg/devfile/validate"
	v2Variables "github.com/maysunfaisal/api/v2/pkg/validation/variables"
)

// ParseFromURLAndValidate func parses the devfile data from the url
// and validates the devfile integrity with the schema
// and validates the devfile data.
// Creates devfile context and runtime objects.
// Deprecated, use ParseDevfileAndValidate() instead
func ParseFromURLAndValidate(url string) (d parser.DevfileObj, err error) {

	// read and parse devfile from the given URL
	d, err = parser.ParseFromURL(url)
	if err != nil {
		return d, err
	}

	// generic validation on devfile content
	err = validate.ValidateDevfileData(d.Data)
	if err != nil {
		return d, err
	}

	return d, err
}

// ParseFromDataAndValidate func parses the devfile data
// and validates the devfile integrity with the schema
// and validates the devfile data.
// Creates devfile context and runtime objects.
// Deprecated, use ParseDevfileAndValidate() instead
func ParseFromDataAndValidate(data []byte) (d parser.DevfileObj, err error) {
	// read and parse devfile from the given bytes
	d, err = parser.ParseFromData(data)
	if err != nil {
		return d, err
	}
	// generic validation on devfile content
	err = validate.ValidateDevfileData(d.Data)
	if err != nil {
		return d, err
	}

	return d, err
}

// ParseAndValidate func parses the devfile data
// and validates the devfile integrity with the schema
// and validates the devfile data.
// Creates devfile context and runtime objects.
// Deprecated, use ParseDevfileAndValidate() instead
func ParseAndValidate(path string) (d parser.DevfileObj, err error) {

	// read and parse devfile from given path
	d, err = parser.Parse(path)
	if err != nil {
		return d, err
	}

	// generic validation on devfile content
	err = validate.ValidateDevfileData(d.Data)
	if err != nil {
		return d, err
	}

	return d, err
}

// ParseDevfileAndValidate func parses the devfile data
// and validates the devfile integrity with the schema
// and validates the devfile data.
// Creates devfile context and runtime objects.
func ParseDevfileAndValidate(args parser.ParserArgs) (d parser.DevfileObj, err error) {
	d, err = parser.ParseDevfile(args)
	if err != nil {
		return d, err
	}

	if d.Data.GetSchemaVersion() == "2.1.0" {
		warning := v2Variables.ValidateAndReplaceGlobalVariable(d.Data.GetDevfileWorkspaceSpec())
		if !reflect.DeepEqual(warning, v2Variables.VariableWarning{}) {
			if len(warning.Commands) > 0 {
				fmt.Printf("top-level variable warning - the following commands reference invalid variables:\n")
				for command, keys := range warning.Commands {
					fmt.Printf("%s: %s\n", command, strings.Join(keys, ","))
				}
			}

			if len(warning.Components) > 0 {
				fmt.Printf("top-level variable warning - the following components reference invalid variables:\n")
				for component, keys := range warning.Components {
					fmt.Printf("%s: %s\n", component, strings.Join(keys, ","))
				}
			}

			if len(warning.Projects) > 0 {
				fmt.Printf("top-level variable warning - the following projects reference invalid variables:\n")
				for project, keys := range warning.Projects {
					fmt.Printf("%s: %s\n", project, strings.Join(keys, ","))
				}
			}

			if len(warning.StarterProjects) > 0 {
				fmt.Printf("top-level variable warning - the following starter projects reference invalid variables:\n")
				for starterProject, keys := range warning.StarterProjects {
					fmt.Printf("%s: %s\n", starterProject, strings.Join(keys, ","))
				}
			}

			fmt.Printf("\n")
		}
	}

	// generic validation on devfile content
	err = validate.ValidateDevfileData(d.Data)
	if err != nil {
		return d, err
	}

	return d, err
}
