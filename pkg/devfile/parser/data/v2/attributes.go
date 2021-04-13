package v2

import (
	"fmt"

	attributes "github.com/maysunfaisal/api/v2/pkg/attributes"
)

func (d *DevfileV2) GetTopLevelAttributes() (attributes.Attributes, error) {

	// This feature was introduced in 210; so any version 210 and up should use the 210 implementation
	switch d.SchemaVersion {
	case "2.0.0":
		return attributes.Attributes{}, fmt.Errorf("top-level attributes is not supported in devfile schema version 2.0.0")
	default:
		return d.Attributes, nil
	}
}
