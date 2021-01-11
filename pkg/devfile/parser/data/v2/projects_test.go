package v2

import (
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	v1 "github.com/maysunfaisal/api/pkg/apis/workspaces/v1alpha2"
)

func TestDevfile200_AddProjects(t *testing.T) {
	currentProject := []v1.Project{
		{
			Name:      "java-starter",
			ClonePath: "/project",
		},
		{
			Name:      "quarkus-starter",
			ClonePath: "/test",
		},
	}

	d := &DevfileV2{
		v1.Devfile{
			DevWorkspaceTemplateSpec: v1.DevWorkspaceTemplateSpec{
				DevWorkspaceTemplateSpecContent: v1.DevWorkspaceTemplateSpecContent{
					Projects: currentProject,
				},
			},
		},
	}

	tests := []struct {
		name    string
		wantErr bool
		args    []v1.Project
	}{
		{
			name:    "case:1 It should add project",
			wantErr: false,
			args: []v1.Project{
				{
					Name: "nodejs",
				},
				{
					Name: "spring-pet-clinic",
				},
			},
		},

		{
			name:    "case:2 It should give error if tried to add already present starter project",
			wantErr: true,
			args: []v1.Project{
				{
					Name: "quarkus-starter",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := d.AddProjects(tt.args)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("unexpected error: %v", err)
				}
				return
			}
			if tt.wantErr {
				t.Errorf("expected error, got %v", err)
				return
			}
			wantProjects := append(currentProject, tt.args...)

			if !reflect.DeepEqual(d.Projects, wantProjects) {
				t.Errorf("wanted: %v, got: %v, difference at %v", wantProjects, d.Projects, pretty.Compare(wantProjects, d.Projects))
			}
		})
	}

}

func TestDevfile200_UpdateProject(t *testing.T) {
	tests := []struct {
		name              string
		args              v1.Project
		devfilev2         *DevfileV2
		expectedDevfilev2 *DevfileV2
	}{
		{
			name: "case:1 It should update project for existing project",
			args: v1.Project{
				Name:      "nodejs",
				ClonePath: "/test",
			},
			devfilev2: &DevfileV2{
				v1.Devfile{
					DevWorkspaceTemplateSpec: v1.DevWorkspaceTemplateSpec{
						DevWorkspaceTemplateSpecContent: v1.DevWorkspaceTemplateSpecContent{
							Projects: []v1.Project{
								{
									Name:      "nodejs",
									ClonePath: "/project",
								},
								{
									Name:      "java-starter",
									ClonePath: "/project",
								},
							},
						},
					},
				},
			},
			expectedDevfilev2: &DevfileV2{
				v1.Devfile{
					DevWorkspaceTemplateSpec: v1.DevWorkspaceTemplateSpec{
						DevWorkspaceTemplateSpecContent: v1.DevWorkspaceTemplateSpecContent{
							Projects: []v1.Project{
								{
									Name:      "nodejs",
									ClonePath: "/test",
								},
								{
									Name:      "java-starter",
									ClonePath: "/project",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "case:2 It should not update project for non existing project",
			args: v1.Project{
				Name:      "quarkus-starter",
				ClonePath: "/project",
			},
			devfilev2: &DevfileV2{
				v1.Devfile{
					DevWorkspaceTemplateSpec: v1.DevWorkspaceTemplateSpec{
						DevWorkspaceTemplateSpecContent: v1.DevWorkspaceTemplateSpecContent{
							Projects: []v1.Project{
								{
									Name:      "nodejs",
									ClonePath: "/project",
								},
								{
									Name:      "java-starter",
									ClonePath: "/project",
								},
							},
						},
					},
				},
			},
			expectedDevfilev2: &DevfileV2{
				v1.Devfile{
					DevWorkspaceTemplateSpec: v1.DevWorkspaceTemplateSpec{
						DevWorkspaceTemplateSpecContent: v1.DevWorkspaceTemplateSpecContent{
							Projects: []v1.Project{
								{
									Name:      "nodejs",
									ClonePath: "/project",
								},
								{
									Name:      "java-starter",
									ClonePath: "/project",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.devfilev2.UpdateProject(tt.args)

			if !reflect.DeepEqual(tt.devfilev2, tt.expectedDevfilev2) {
				t.Errorf("wanted: %v, got: %v, difference at %v", tt.expectedDevfilev2, tt.devfilev2, pretty.Compare(tt.expectedDevfilev2, tt.devfilev2))
			}
		})
	}
}

func TestDevfile200_AddStarterProjects(t *testing.T) {
	currentProject := []v1.StarterProject{
		{
			Name:        "java-starter",
			Description: "starter project for java",
		},
		{
			Name:        "quarkus-starter",
			Description: "starter project for quarkus",
		},
	}

	d := &DevfileV2{
		v1.Devfile{
			DevWorkspaceTemplateSpec: v1.DevWorkspaceTemplateSpec{
				DevWorkspaceTemplateSpecContent: v1.DevWorkspaceTemplateSpecContent{
					StarterProjects: currentProject,
				},
			},
		},
	}

	tests := []struct {
		name    string
		wantErr bool
		args    []v1.StarterProject
	}{
		{
			name:    "case:1 It should add starter project",
			wantErr: false,
			args: []v1.StarterProject{
				{
					Name:        "nodejs",
					Description: "starter project for nodejs",
				},
				{
					Name:        "spring-pet-clinic",
					Description: "starter project for springboot",
				},
			},
		},

		{
			name:    "case:2 It should give error if tried to add already present starter project",
			wantErr: true,
			args: []v1.StarterProject{
				{
					Name:        "quarkus-starter",
					Description: "starter project for quarkus",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := d.AddStarterProjects(tt.args)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("unexpected error: %v", err)
				}
				return
			}
			if tt.wantErr {
				t.Errorf("expected error, got %v", err)
				return
			}
			wantProjects := append(currentProject, tt.args...)

			if !reflect.DeepEqual(d.StarterProjects, wantProjects) {
				t.Errorf("wanted: %v, got: %v, difference at %v", wantProjects, d.StarterProjects, pretty.Compare(wantProjects, d.StarterProjects))
			}
		})
	}

}

func TestDevfile200_UpdateStarterProject(t *testing.T) {
	tests := []struct {
		name              string
		args              v1.StarterProject
		devfilev2         *DevfileV2
		expectedDevfilev2 *DevfileV2
	}{
		{
			name: "case:1 It should update project for existing project",
			args: v1.StarterProject{
				Name:   "nodejs",
				SubDir: "/test",
			},
			devfilev2: &DevfileV2{
				v1.Devfile{
					DevWorkspaceTemplateSpec: v1.DevWorkspaceTemplateSpec{
						DevWorkspaceTemplateSpecContent: v1.DevWorkspaceTemplateSpecContent{
							StarterProjects: []v1.StarterProject{
								{
									Name:   "nodejs",
									SubDir: "/project",
								},
								{
									Name:   "java-starter",
									SubDir: "/project",
								},
							},
						},
					},
				},
			},
			expectedDevfilev2: &DevfileV2{
				v1.Devfile{
					DevWorkspaceTemplateSpec: v1.DevWorkspaceTemplateSpec{
						DevWorkspaceTemplateSpecContent: v1.DevWorkspaceTemplateSpecContent{
							StarterProjects: []v1.StarterProject{
								{
									Name:   "nodejs",
									SubDir: "/test",
								},
								{
									Name:   "java-starter",
									SubDir: "/project",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "case:2 It should not update project for non existing project",
			args: v1.StarterProject{
				Name:   "quarkus-starter",
				SubDir: "/project",
			},
			devfilev2: &DevfileV2{
				v1.Devfile{
					DevWorkspaceTemplateSpec: v1.DevWorkspaceTemplateSpec{
						DevWorkspaceTemplateSpecContent: v1.DevWorkspaceTemplateSpecContent{
							StarterProjects: []v1.StarterProject{
								{
									Name:   "nodejs",
									SubDir: "/project",
								},
								{
									Name:   "java-starter",
									SubDir: "/project",
								},
							},
						},
					},
				},
			},
			expectedDevfilev2: &DevfileV2{
				v1.Devfile{
					DevWorkspaceTemplateSpec: v1.DevWorkspaceTemplateSpec{
						DevWorkspaceTemplateSpecContent: v1.DevWorkspaceTemplateSpecContent{
							StarterProjects: []v1.StarterProject{
								{
									Name:   "nodejs",
									SubDir: "/project",
								},
								{
									Name:   "java-starter",
									SubDir: "/project",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.devfilev2.UpdateStarterProject(tt.args)

			if !reflect.DeepEqual(tt.devfilev2, tt.expectedDevfilev2) {
				t.Errorf("wanted: %v, got: %v, difference at %v", tt.expectedDevfilev2, tt.devfilev2, pretty.Compare(tt.expectedDevfilev2, tt.devfilev2))
			}
		})
	}
}
