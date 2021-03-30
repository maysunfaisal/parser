# Devfile Library

## About

The Devfile Parser library is a Golang module that:
1. parses the devfile.yaml as specified by the [api](https://devfile.github.io/maysunfaisal/api-reference.html) & [schema](https://github.com/maysunfaisal/api/tree/master/schemas/latest).
2. writes to the devfile.yaml with the updated data.
3. generates Kubernetes objects for the various devfile resources.
4. defines util functions for the devfile.

## Usage

The function documentation can be accessed via [pkg.go.dev](https://pkg.go.dev/github.com/maysunfaisal/parser). 
1. To parse a devfile, visit pkg/devfile/parse.go 
   ```
   // Parses the devfile and validates the devfile data
   devfile, err := devfilePkg.ParseAndValidate(devfileLocation)

   // To get all the components from the devfile
   components, err := devfile.Data.GetComponents(DevfileOptions{})

   // To get all the components from the devfile with attributes tagged - tool: console-import
   // & import: {strategy: Dockerfile}
   components, err := devfile.Data.GetComponents(DevfileOptions{
      Filter: map[string]interface{}{
			"tool": "console-import",
			"import": map[string]interface{}{
				"strategy": "Dockerfile",
			},
		},
   })

   // To get all the volume components
   components, err := devfile.Data.GetComponents(DevfileOptions{
		ComponentOptions: ComponentOptions{
			ComponentType: v1.VolumeComponentType,
		},
   })

   // To get all the exec commands that belong to the build group
   commands, err := devfile.Data.GetCommands(DevfileOptions{
		CommandOptions: CommandOptions{
			CommandType: v1.ExecCommandType,
			CommandGroupKind: v1.BuildCommandGroupKind,
		},
   })
   ```
2. To get the Kubernetes objects from the devfile, visit pkg/devfile/generator/generators.go
   ```
    // To get a slice of Kubernetes containers of type corev1.Container from the devfile component containers
    containers, err := generator.GetContainers(devfile)

    // To generate a Kubernetes deployment of type v1.Deployment
    deployParams := generator.DeploymentParams{
		TypeMeta:          generator.GetTypeMeta(deploymentKind, deploymentAPIVersion),
		ObjectMeta:        generator.GetObjectMeta(name, namespace, labels, annotations),
		InitContainers:    initContainers,
		Containers:        containers,
		Volumes:           volumes,
		PodSelectorLabels: labels,
	}
	deployment := generator.GetDeployment(deployParams)
   ```

## Updating Library Schema

Run `updateApi.sh` can update to use latest `github.com/maysunfaisal/api` and update the schema saved under `pkg/devfile/parser/data`

The script also accepts version number as an argument to update devfile schema for a specific devfile version.
For example, run the following command will update devfile schema for 2.0.0
```
./updateApi.sh 2.0.0
```
Running the script with no arguments will default to update the latest devfile version

## Projects using maysunfaisal/parser

The following projects are consuming this library as a Golang dependency

* [odo](https://github.com/openshift/odo)
* [OpenShift Console](https://github.com/openshift/console)

In the future, [Workspace Operator](https://github.com/devfile/devworkspace-operator) will be the next consumer of maysunfaisal/parser.

## Issues

Issues are tracked in the [maysunfaisal/api](https://github.com/maysunfaisal/api) repo with the label [area/library](https://github.com/maysunfaisal/api/issues?q=is%3Aopen+is%3Aissue+label%3Aarea%2Flibrary) 

## Releases

For maysunfaisal/parser releases, please check the release [page](https://github.com/maysunfaisal/parser/releases).
