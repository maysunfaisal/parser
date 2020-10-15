package v2

import (
	"fmt"
	"strings"

	v1 "github.com/devfile/api/pkg/apis/workspaces/v1alpha2"
	"github.com/devfile/parser/pkg/devfile/parser/data/common"
)

// GetCustomType gets the custom type
func (d *DevfileV2) GetCustomType() string {

	devfileV2Type, err := NewDevfileDataV2(d.SchemaVersion)
	if err != nil {
		fmt.Printf("Error on DevfileV2 GetCustomType %v\n", err)
	}

	return devfileV2Type.GetCustomType()
}

//SetSchemaVersion sets devfile schema version
func (d *DevfileV2) SetSchemaVersion(version string) {
	d.SchemaVersion = version
}

// GetMetadata returns the DevfileMetadata Object parsed from devfile
func (d *DevfileV2) GetMetadata() v1.DevfileMetadata {
	return d.Metadata
}

// SetMetadata sets the metadata for devfile
func (d *DevfileV2) SetMetadata(name, version string) {
	d.Metadata = v1.DevfileMetadata{
		Name:    name,
		Version: version,
	}
}

// GetParent returns the Parent object parsed from devfile
func (d *DevfileV2) GetParent() *v1.Parent {
	return d.Parent
}

// SetParent sets the parent for the devfile
func (d *DevfileV2) SetParent(parent *v1.Parent) {
	d.Parent = parent
}

// GetProjects returns the Project Object parsed from devfile
func (d *DevfileV2) GetProjects() []v1.Project {
	return d.Projects
}

// AddProjects adss the slice of Devfile projects to the Devfile's project list
// if a project is already defined, error out
func (d *DevfileV2) AddProjects(projects []v1.Project) error {
	projectsMap := make(map[string]bool)
	for _, project := range d.Projects {
		projectsMap[project.Name] = true
	}

	for _, project := range projects {
		if _, ok := projectsMap[project.Name]; !ok {
			d.Projects = append(d.Projects, project)
		} else {
			return &common.AlreadyExistError{Name: project.Name, Field: "project"}
		}
	}
	return nil
}

// UpdateProject updates the slice of Devfile projects parsed from the Devfile
func (d *DevfileV2) UpdateProject(project v1.Project) {
	for i := range d.Projects {
		if d.Projects[i].Name == strings.ToLower(project.Name) {
			d.Projects[i] = project
		}
	}
}

// GetComponents returns the slice of Component objects parsed from the Devfile
func (d *DevfileV2) GetComponents() []v1.Component {
	return d.Components
}

// GetAliasedComponents returns the slice of Component objects that each have an alias
func (d *DevfileV2) GetAliasedComponents() []v1.Component {
	// V2 has name required in jsonSchema
	return d.Components
}

// AddComponents adds the slice of Component objects to the devfile's components
// if a component is already defined, error out
func (d *DevfileV2) AddComponents(components []v1.Component) error {

	// different map for volume and container component as a volume and a container with same name
	// can exist in devfile
	containerMap := make(map[string]bool)
	volumeMap := make(map[string]bool)

	for _, component := range d.Components {
		if component.Volume != nil {
			volumeMap[component.Name] = true
		}
		if component.Container != nil {
			containerMap[component.Name] = true
		}
	}

	for _, component := range components {

		if component.Volume != nil {
			if _, ok := volumeMap[component.Name]; !ok {
				d.Components = append(d.Components, component)
			} else {
				return &common.AlreadyExistError{Name: component.Name, Field: "component"}
			}
		}

		if component.Container != nil {
			if _, ok := containerMap[component.Name]; !ok {
				d.Components = append(d.Components, component)
			} else {
				return &common.AlreadyExistError{Name: component.Name, Field: "component"}
			}
		}
	}
	return nil
}

// UpdateComponent updates the component with the given name
func (d *DevfileV2) UpdateComponent(component v1.Component) {
	index := -1
	for i := range d.Components {
		if d.Components[i].Name == strings.ToLower(component.Name) {
			index = i
			break
		}
	}
	if index != -1 {
		d.Components[index] = component
	}
}

// GetCommands returns the slice of Command objects parsed from the Devfile
func (d *DevfileV2) GetCommands() map[string]v1.Command {

	commands := make(map[string]v1.Command, len(d.Commands))

	for _, command := range d.Commands {
		// we convert devfile command id to lowercase so that we can handle
		// cases efficiently without being error prone
		// we also convert the odo push commands from build-command and run-command flags
		commands[common.SetIDToLower(&command)] = command
	}

	return commands
}

// AddCommands adds the slice of Command objects to the Devfile's commands
// if a command is already defined, error out
func (d *DevfileV2) AddCommands(commands ...v1.Command) error {
	commandsMap := d.GetCommands()

	for _, command := range commands {
		id := common.GetID(command)
		if _, ok := commandsMap[id]; !ok {
			d.Commands = append(d.Commands, command)
		} else {
			return &common.AlreadyExistError{Name: id, Field: "command"}
		}
	}
	return nil
}

// UpdateCommand updates the command with the given id
func (d *DevfileV2) UpdateCommand(command v1.Command) {
	id := strings.ToLower(common.GetID(command))
	for i := range d.Commands {
		if common.SetIDToLower(&d.Commands[i]) == id {
			d.Commands[i] = command
		}
	}
}

//GetStarterProjects returns the DevfileStarterProject parsed from devfile
func (d *DevfileV2) GetStarterProjects() []v1.StarterProject {
	return d.StarterProjects
}

// AddStarterProjects adds the slice of Devfile starter projects to the Devfile's starter project list
// if a starter project is already defined, error out
func (d *DevfileV2) AddStarterProjects(projects []v1.StarterProject) error {
	projectsMap := make(map[string]bool)
	for _, project := range d.StarterProjects {
		projectsMap[project.Name] = true
	}

	for _, project := range projects {
		if _, ok := projectsMap[project.Name]; !ok {
			d.StarterProjects = append(d.StarterProjects, project)
		} else {
			return &common.AlreadyExistError{Name: project.Name, Field: "starterProject"}
		}
	}
	return nil
}

// UpdateStarterProject updates the slice of Devfile starter projects parsed from the Devfile
func (d *DevfileV2) UpdateStarterProject(project v1.StarterProject) {
	for i := range d.StarterProjects {
		if d.StarterProjects[i].Name == strings.ToLower(project.Name) {
			d.StarterProjects[i] = project
		}
	}
}

// GetEvents returns the Events Object parsed from devfile
func (d *DevfileV2) GetEvents() v1.Events {
	if d.Events != nil {
		return *d.Events
	}
	return v1.Events{}
}

// AddEvents adds the Events Object to the devfile's events
// if the event is already defined in the devfile, error out
func (d *DevfileV2) AddEvents(events v1.Events) error {
	if len(events.PreStop) > 0 {
		if len(d.Events.PreStop) > 0 {
			return &common.AlreadyExistError{Field: "pre stop"}
		}
		d.Events.PreStop = events.PreStop
	}

	if len(events.PreStart) > 0 {
		if len(d.Events.PreStart) > 0 {
			return &common.AlreadyExistError{Field: "pre start"}
		}
		d.Events.PreStart = events.PreStart
	}

	if len(events.PostStop) > 0 {
		if len(d.Events.PostStop) > 0 {
			return &common.AlreadyExistError{Field: "post stop"}
		}
		d.Events.PostStop = events.PostStop
	}

	if len(events.PostStart) > 0 {
		if len(d.Events.PostStart) > 0 {
			return &common.AlreadyExistError{Field: "post start"}
		}
		d.Events.PostStart = events.PostStart
	}

	return nil
}

// UpdateEvents updates the devfile's events
// it only updates the events passed to it
func (d *DevfileV2) UpdateEvents(postStart, postStop, preStart, preStop []string) {
	if len(postStart) != 0 {
		d.Events.PostStart = postStart
	}
	if len(postStop) != 0 {
		d.Events.PostStop = postStop
	}
	if len(preStart) != 0 {
		d.Events.PreStart = preStart
	}
	if len(preStop) != 0 {
		d.Events.PreStop = preStop
	}
}

// AddVolume adds the volume to the devFile and mounts it to all the container components
func (d *DevfileV2) AddVolume(volumeComponent v1.Component, path string) error {
	volumeExists := false
	var pathErrorContainers []string
	for _, component := range d.Components {
		if component.Container != nil {
			for _, volumeMount := range component.Container.VolumeMounts {
				if volumeMount.Path == path {
					var err = fmt.Errorf("another volume, %s, is mounted to the same path: %s, on the container: %s", volumeMount.Name, path, component.Name)
					pathErrorContainers = append(pathErrorContainers, err.Error())
				}
			}
			component.Container.VolumeMounts = append(component.Container.VolumeMounts, v1.VolumeMount{
				Name: volumeComponent.Name,
				Path: path,
			})
		} else if component.Volume != nil && component.Name == volumeComponent.Name {
			volumeExists = true
			break
		}
	}

	if volumeExists {
		return &common.AlreadyExistError{
			Field: "volume",
			Name:  volumeComponent.Name,
		}
	}

	if len(pathErrorContainers) > 0 {
		return fmt.Errorf("errors while creating volume:\n%s", strings.Join(pathErrorContainers, "\n"))
	}

	d.Components = append(d.Components, volumeComponent)

	return nil
}

// DeleteVolume removes the volume from the devFile and removes all the related volume mounts
func (d *DevfileV2) DeleteVolume(name string) error {
	found := false
	for i := len(d.Components) - 1; i >= 0; i-- {
		if d.Components[i].Container != nil {
			var tmp []v1.VolumeMount
			for _, volumeMount := range d.Components[i].Container.VolumeMounts {
				if volumeMount.Name != name {
					tmp = append(tmp, volumeMount)
				}
			}
			d.Components[i].Container.VolumeMounts = tmp
		} else if d.Components[i].Volume != nil {
			if d.Components[i].Name == name {
				found = true
				d.Components = append(d.Components[:i], d.Components[i+1:]...)
			}
		}
	}

	if !found {
		return &common.NotFoundError{
			Field: "volume",
			Name:  name,
		}
	}

	return nil
}

// GetVolumeMountPath gets the mount path of the required volume
func (d *DevfileV2) GetVolumeMountPath(name string) (string, error) {
	volumeFound := false
	mountFound := false
	path := ""

	for _, component := range d.Components {
		if component.Container != nil {
			for _, volumeMount := range component.Container.VolumeMounts {
				if volumeMount.Name == name {
					mountFound = true
					path = volumeMount.Path
				}
			}
		} else if component.Volume != nil {
			volumeFound = true
		}
	}
	if volumeFound && mountFound {
		return path, nil
	} else if !mountFound && volumeFound {
		return "", fmt.Errorf("volume not mounted to any component")
	}
	return "", &common.NotFoundError{
		Field: "volume",
		Name:  "name",
	}
}
