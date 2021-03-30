package v2

import (
	v1 "github.com/maysunfaisal/api/v2/pkg/apis/workspaces/v1alpha2"
)

// GetDevfileWorkspaceContent returns the workspace content for the devfile
func (d *DevfileV2) GetDevfileWorkspaceContent() *v1.DevWorkspaceTemplateSpecContent {

	return &d.DevWorkspaceTemplateSpecContent
}

// SetDevfileWorkspaceContent sets the workspace content
func (d *DevfileV2) SetDevfileWorkspaceContent(content v1.DevWorkspaceTemplateSpecContent) {
	d.DevWorkspaceTemplateSpecContent = content
}

// GetDevfileWorkspace returns the workspace content for the devfile
func (d *DevfileV2) GetDevfileWorkspace() *v1.DevWorkspaceTemplateSpec {

	return &d.DevWorkspaceTemplateSpec
}

// SetDevfileWorkspace sets the workspace content
func (d *DevfileV2) SetDevfileWorkspace(spec v1.DevWorkspaceTemplateSpec) {
	d.DevWorkspaceTemplateSpec = spec
}
