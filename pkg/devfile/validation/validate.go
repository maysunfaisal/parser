package validation

import (
	"fmt"

	v2 "github.com/devfile/library/pkg/devfile/parser/data/v2"
	devfilev1 "github.com/maysunfaisal/api/pkg/apis/workspaces/v1alpha2"
	"github.com/maysunfaisal/api/pkg/validation"
)

// ValidateDevfileData validates whether sections of devfile are odo compatible
func ValidateDevfileData(data interface{}) error {
	var components []devfilev1.Component
	// var commandsMap map[string]devfilev1.Command
	// var events devfilev1.Events

	switch d := data.(type) {
	case *v2.DevfileV2:
		components = d.Components
		// commandsMap = d.GetCommands()
		// events = d.GetEvents()

		// Validate all the devfile components before validating commands
		if err := validation.ValidateComponents(components); err != nil {
			return err
		}

		// // Validate all the devfile commands before validating events
		// if err := validateCommands(d.Commands, commandsMap, components); err != nil {
		// 	return err
		// }

		// // Validate all the events after validating the commands
		// if err := validateEvents(events, commandsMap); err != nil {
		// 	return err
		// }
	default:
		fmt.Printf("data %+v\n\n", data.(*v2.DevfileV2))
		return fmt.Errorf("unknown devfile type %T", d)
	}

	// Successful
	return nil

}
