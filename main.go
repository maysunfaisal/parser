package main

import (
	"fmt"
	"reflect"

	devfilepkg "github.com/devfile/parser/pkg/devfile"
	"github.com/devfile/parser/pkg/devfile/parser"
	v2 "github.com/devfile/parser/pkg/devfile/parser/data/v2"
)

func main() {
	devfile, err := ParseDevfile("devfile.yaml")
	if err != nil {
		fmt.Println(err)
	} else {
		devdata := devfile.Data
		if (reflect.TypeOf(devdata) == reflect.TypeOf(&v2.DevfileV2{})) {
			d := devdata.(*v2.DevfileV2)
			fmt.Println(d.SchemaVersion)
		}

		for _, component := range devfile.Data.GetComponents() {
			/*
				if component.Dockerfile != nil {
							fmt.Println(component.Dockerfile.DockerfileLocation)
						}
			*/
			if component.Container != nil {
				fmt.Println(component.Container.Image)
			}
		}

		for _, command := range devfile.Data.GetCommands() {
			if command.Exec != nil {
				fmt.Println(command.Exec.Group.Kind)
			}
		}

		fmt.Println(devdata.GetCustomType())
	}

}

//ParseDevfile to parse devfile from library
func ParseDevfile(devfileLocation string) (devfileoj parser.DevfileObj, err error) {

	var devfile parser.DevfileObj
	devfile, err = devfilepkg.ParseAndValidate(devfileLocation)
	return devfile, err

	// return
}
