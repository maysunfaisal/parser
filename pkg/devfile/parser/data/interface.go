package data

import (
	v1 "github.com/devfile/api/pkg/apis/workspaces/v1alpha2"
	// v2 "github.com/devfile/parser/pkg/devfile/parser/data/v2"
)

// DevfileData is an interface that defines functions for Devfile data operations
type DevfileData interface {
	SetSchemaVersion(version string)
	GetMetadata() v1.DevfileMetadata
	SetMetadata(name, version string)

	// parent related methods
	GetParent() *v1.Parent
	SetParent(parent *v1.Parent)

	// event related methods
	GetEvents() v1.Events
	AddEvents(events v1.Events) error
	UpdateEvents(postStart, postStop, preStart, preStop []string)

	// component related methods
	GetComponents() []v1.Component
	AddComponents(components []v1.Component) error
	UpdateComponent(component v1.Component)
	GetAliasedComponents() []v1.Component

	// project related methods
	GetProjects() []v1.Project
	AddProjects(projects []v1.Project) error
	UpdateProject(project v1.Project)

	// starter projects related commands
	GetStarterProjects() []v1.StarterProject
	AddStarterProjects(projects []v1.StarterProject) error
	UpdateStarterProject(project v1.StarterProject)

	// command related methods
	GetCommands() map[string]v1.Command
	AddCommands(commands ...v1.Command) error
	UpdateCommand(command v1.Command)

	// volume related methods
	AddVolume(volume v1.Component, path string) error
	DeleteVolume(name string) error
	GetVolumeMountPath(name string) (string, error)

	GetCustomType() string

	// v2.DevfileDataV2
}
