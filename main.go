package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	devfilepkg "github.com/devfile/library/pkg/devfile"
	"github.com/devfile/library/pkg/devfile/parser"
	v2 "github.com/devfile/library/pkg/devfile/parser/data/v2"
	"github.com/devfile/library/pkg/devfile/parser/data/v2/common"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "updateSchema" {
		ReplaceSchemaFile()
	} else {
		parserTest()
	}
}

func parserTest() {
	var args parser.ParserArgs
	if len(os.Args) > 1 {
		if strings.HasPrefix(os.Args[1], "http") {
			args = parser.ParserArgs{
				URL: os.Args[1],
			}
		} else {
			args = parser.ParserArgs{
				Path: os.Args[1],
			}
		}
		fmt.Printf("parsing devfile from %s\n\n", os.Args[1])

	} else {
		args = parser.ParserArgs{
			Path: "devfile.yaml",
		}
		fmt.Printf("parsing devfile from ./devfile.yaml\n\n")
	}
	devfile, err := devfilepkg.ParseDevfileAndValidate(args)
	if err != nil {
		fmt.Println(err)
	} else {
		devdata := devfile.Data
		if (reflect.TypeOf(devdata) == reflect.TypeOf(&v2.DevfileV2{})) {
			d := devdata.(*v2.DevfileV2)
			fmt.Printf("devfile schema version: %s\n\n", d.SchemaVersion)
		}

		variables, varErr := devfile.Data.GetTopLevelVariables()
		if varErr != nil {
			fmt.Printf("err: %v\n", varErr)
		} else {
			fmt.Printf("My top level variable keys are: ")
			for k, _ := range variables {
				fmt.Printf("%s ", k)
			}
		}

		fmt.Printf("\n\n")

		components, e := devfile.Data.GetComponents(common.DevfileOptions{})
		if e != nil {
			fmt.Printf("err: %v\n", err)
		}
		fmt.Printf("All devfile container components \n")
		for _, component := range components {
			if component.Container != nil {
				fmt.Printf("component name: %s\n", component.Name)
				fmt.Printf("component image: %s\n", component.Container.Image)
				fmt.Printf("component mem limit: %s\n", component.Container.MemoryLimit)
			}
		}
	}

}
