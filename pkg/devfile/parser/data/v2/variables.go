package v2

import "fmt"

func (d *DevfileV2) GetTopLevelVariables() (map[string]string, error) {

	// This feature was introduced in 210; so any version 210 and up should use the 210 implementation
	switch d.SchemaVersion {
	case "2.0.0":
		return nil, fmt.Errorf("top-level variables is not supported in devfile schema version 2.0.0")
	default:
		return d.Variables, nil
	}
}

func (d *DevfileV2) UpdateTopLevelVariables(variables map[string]string) error {

	// This feature was introduced in 210; so any version 210 and up should use the 210 implementation
	switch d.SchemaVersion {
	case "2.0.0":
		return fmt.Errorf("top-level variables is not supported in devfile schema version 2.0.0")
	default:
		for k, v := range variables {
			d.Variables[k] = v
		}
		return nil
	}
}
