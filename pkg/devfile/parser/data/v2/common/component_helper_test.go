package common

import (
	"testing"

	v1 "github.com/maysunfaisal/api/v2/pkg/apis/workspaces/v1alpha2"
)

func TestIsContainer(t *testing.T) {

	tests := []struct {
		name            string
		component       v1.Component
		wantIsSupported bool
	}{
		{
			name: "Case 1: Container component",
			component: v1.Component{
				Name: "name",
				ComponentUnion: v1.ComponentUnion{
					Container: &v1.ContainerComponent{},
				},
			},
			wantIsSupported: true,
		},
		{
			name: "Case 2: Not a container component",
			component: v1.Component{
				Name: "name",
				ComponentUnion: v1.ComponentUnion{
					Openshift: &v1.OpenshiftComponent{},
				},
			},
			wantIsSupported: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isSupported := IsContainer(tt.component)
			if isSupported != tt.wantIsSupported {
				t.Errorf("TestIsContainer error: component support mismatch, expected: %v got: %v", tt.wantIsSupported, isSupported)
			}
		})
	}

}

func TestIsVolume(t *testing.T) {

	tests := []struct {
		name            string
		component       v1.Component
		wantIsSupported bool
	}{
		{
			name: "Case 1: Volume component",
			component: v1.Component{
				Name: "name",
				ComponentUnion: v1.ComponentUnion{
					Volume: &v1.VolumeComponent{
						Volume: v1.Volume{
							Size: "size",
						},
					},
				},
			},
			wantIsSupported: true,
		},
		{
			name: "Case 2: Not a volume component",
			component: v1.Component{
				Name: "name",
				ComponentUnion: v1.ComponentUnion{
					Openshift: &v1.OpenshiftComponent{},
				},
			},
			wantIsSupported: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isSupported := IsVolume(tt.component)
			if isSupported != tt.wantIsSupported {
				t.Errorf("TestIsVolume error: component support mismatch, expected: %v got: %v", tt.wantIsSupported, isSupported)
			}
		})
	}

}

func TestGetComponentType(t *testing.T) {

	tests := []struct {
		name          string
		component     v1.Component
		wantErr       bool
		componentType v1.ComponentType
	}{
		{
			name: "Volume component",
			component: v1.Component{
				Name: "name",
				ComponentUnion: v1.ComponentUnion{
					Volume: &v1.VolumeComponent{
						Volume: v1.Volume{},
					},
				},
			},
			componentType: v1.VolumeComponentType,
			wantErr:       false,
		},
		{
			name: "Openshift component",
			component: v1.Component{
				Name: "name",
				ComponentUnion: v1.ComponentUnion{
					Openshift: &v1.OpenshiftComponent{},
				},
			},
			componentType: v1.OpenshiftComponentType,
			wantErr:       false,
		},
		{
			name: "Kubernetes component",
			component: v1.Component{
				Name: "name",
				ComponentUnion: v1.ComponentUnion{
					Kubernetes: &v1.KubernetesComponent{},
				},
			},
			componentType: v1.KubernetesComponentType,
			wantErr:       false,
		},
		{
			name: "Container component",
			component: v1.Component{
				Name: "name",
				ComponentUnion: v1.ComponentUnion{
					Container: &v1.ContainerComponent{},
				},
			},
			componentType: v1.ContainerComponentType,
			wantErr:       false,
		},
		{
			name: "Plugin component",
			component: v1.Component{
				Name: "name",
				ComponentUnion: v1.ComponentUnion{
					Plugin: &v1.PluginComponent{},
				},
			},
			componentType: v1.PluginComponentType,
			wantErr:       false,
		},
		{
			name: "Custom component",
			component: v1.Component{
				Name: "name",
				ComponentUnion: v1.ComponentUnion{
					Custom: &v1.CustomComponent{},
				},
			},
			componentType: v1.CustomComponentType,
			wantErr:       false,
		},
		{
			name: "Unknown component",
			component: v1.Component{
				Name:           "name",
				ComponentUnion: v1.ComponentUnion{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetComponentType(tt.component)
			// Unexpected error
			if (err != nil) != tt.wantErr {
				t.Errorf("TestGetComponentType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.componentType {
				t.Errorf("TestGetComponentType error: component type mismatch, expected: %v got: %v", tt.componentType, got)
			}
		})
	}

}
