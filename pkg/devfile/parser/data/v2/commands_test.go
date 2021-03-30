package v2

import (
	"reflect"
	"testing"

	v1 "github.com/maysunfaisal/api/v2/pkg/apis/workspaces/v1alpha2"
	"github.com/maysunfaisal/api/v2/pkg/attributes"
	"github.com/maysunfaisal/parser/pkg/devfile/parser/data/v2/common"
	"github.com/stretchr/testify/assert"
)

func TestDevfile200_GetCommands(t *testing.T) {

	tests := []struct {
		name            string
		currentCommands []v1.Command
		filterOptions   common.DevfileOptions
		wantCommands    []string
		wantErr         bool
	}{
		{
			name: "Get all the commands",
			currentCommands: []v1.Command{
				{
					Id: "command1",
					CommandUnion: v1.CommandUnion{
						Exec: &v1.ExecCommand{},
					},
				},
				{
					Id: "command2",
					CommandUnion: v1.CommandUnion{
						Composite: &v1.CompositeCommand{},
					},
				},
			},
			wantCommands: []string{"command1", "command2"},
			wantErr:      false,
		},
		{
			name: "Get the filtered commands",
			currentCommands: []v1.Command{
				{
					Id: "command1",
					Attributes: attributes.Attributes{}.FromStringMap(map[string]string{
						"firstString":  "firstStringValue",
						"secondString": "secondStringValue",
					}),
					CommandUnion: v1.CommandUnion{
						Exec: &v1.ExecCommand{
							LabeledCommand: v1.LabeledCommand{
								BaseCommand: v1.BaseCommand{
									Group: &v1.CommandGroup{
										Kind: v1.BuildCommandGroupKind,
									},
								},
							},
						},
					},
				},
				{
					Id: "command2",
					Attributes: attributes.Attributes{}.FromStringMap(map[string]string{
						"firstString": "firstStringValue",
						"thirdString": "thirdStringValue",
					}),
					CommandUnion: v1.CommandUnion{
						Composite: &v1.CompositeCommand{},
					},
				},
				{
					Id: "command3",
					Attributes: attributes.Attributes{}.FromStringMap(map[string]string{
						"firstString": "firstStringValue",
						"thirdString": "thirdStringValue",
					}),
					CommandUnion: v1.CommandUnion{
						Composite: &v1.CompositeCommand{
							LabeledCommand: v1.LabeledCommand{
								BaseCommand: v1.BaseCommand{
									Group: &v1.CommandGroup{
										Kind: v1.BuildCommandGroupKind,
									},
								},
							},
						},
					},
				},
				{
					Id: "command4",
					Attributes: attributes.Attributes{}.FromStringMap(map[string]string{
						"thirdString": "thirdStringValue",
					}),
					CommandUnion: v1.CommandUnion{
						Apply: &v1.ApplyCommand{
							LabeledCommand: v1.LabeledCommand{
								BaseCommand: v1.BaseCommand{
									Group: &v1.CommandGroup{
										Kind: v1.BuildCommandGroupKind,
									},
								},
							},
						},
					},
				},
				{
					Id: "command5",
					Attributes: attributes.Attributes{}.FromStringMap(map[string]string{
						"firstString": "firstStringValue",
						"thirdString": "thirdStringValue",
					}),
					CommandUnion: v1.CommandUnion{
						Composite: &v1.CompositeCommand{
							LabeledCommand: v1.LabeledCommand{
								BaseCommand: v1.BaseCommand{
									Group: &v1.CommandGroup{
										Kind: v1.RunCommandGroupKind,
									},
								},
							},
						},
					},
				},
			},
			filterOptions: common.DevfileOptions{
				Filter: map[string]interface{}{
					"firstString": "firstStringValue",
				},
				CommandOptions: common.CommandOptions{
					CommandGroupKind: v1.BuildCommandGroupKind,
					CommandType:      v1.CompositeCommandType,
				},
			},
			wantCommands: []string{"command3"},
			wantErr:      false,
		},
		{
			name: "Wrong filter for commands",
			currentCommands: []v1.Command{
				{
					Id: "command1",
					Attributes: attributes.Attributes{}.FromStringMap(map[string]string{
						"firstString":  "firstStringValue",
						"secondString": "secondStringValue",
					}),
					CommandUnion: v1.CommandUnion{
						Exec: &v1.ExecCommand{},
					},
				},
				{
					Id: "command2",
					Attributes: attributes.Attributes{}.FromStringMap(map[string]string{
						"firstString": "firstStringValue",
						"thirdString": "thirdStringValue",
					}),
					CommandUnion: v1.CommandUnion{
						Composite: &v1.CompositeCommand{},
					},
				},
			},
			filterOptions: common.DevfileOptions{
				Filter: map[string]interface{}{
					"firstStringIsWrong": "firstStringValue",
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid command type",
			currentCommands: []v1.Command{
				{
					Id: "command1",
					Attributes: attributes.Attributes{}.FromStringMap(map[string]string{
						"firstString": "firstStringValue",
					}),
					CommandUnion: v1.CommandUnion{},
				},
			},
			filterOptions: common.DevfileOptions{
				Filter: map[string]interface{}{
					"firstString": "firstStringValue",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DevfileV2{
				v1.Devfile{
					DevWorkspaceTemplateSpec: v1.DevWorkspaceTemplateSpec{
						DevWorkspaceTemplateSpecContent: v1.DevWorkspaceTemplateSpecContent{
							Commands: tt.currentCommands,
						},
					},
				},
			}

			commands, err := d.GetCommands(tt.filterOptions)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestDevfile200_GetCommands() error = %v, wantErr %v", err, tt.wantErr)
			} else if err == nil {
				// confirm the length of actual vs expected
				if len(commands) != len(tt.wantCommands) {
					t.Errorf("TestDevfile200_GetCommands() error - length of expected commands is not the same as the length of actual commands")
					return
				}

				// compare the command slices for content
				for _, wantCommand := range tt.wantCommands {
					matched := false
					for _, command := range commands {
						if wantCommand == command.Id {
							matched = true
						}
					}

					if !matched {
						t.Errorf("TestDevfile200_GetCommands() error - command %s not found in the devfile", wantCommand)
					}
				}
			}
		})
	}
}

func TestDevfile200_AddCommands(t *testing.T) {

	type args struct {
		name string
	}
	tests := []struct {
		name            string
		currentCommands []v1.Command
		newCommands     []v1.Command
		wantErr         bool
	}{
		{
			name: "case 1: Command does not exist",
			currentCommands: []v1.Command{
				{
					Id: "command1",
					CommandUnion: v1.CommandUnion{
						Exec: &v1.ExecCommand{},
					},
				},
			},
			newCommands: []v1.Command{
				{
					Id: "command2",
					CommandUnion: v1.CommandUnion{
						Exec: &v1.ExecCommand{},
					},
				},
				{
					Id: "command3",
					CommandUnion: v1.CommandUnion{
						Exec: &v1.ExecCommand{},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "case 2: Command does exist",
			currentCommands: []v1.Command{
				{
					Id: "command1",
					CommandUnion: v1.CommandUnion{
						Exec: &v1.ExecCommand{},
					},
				},
			},
			newCommands: []v1.Command{
				{
					Id: "command1",
					CommandUnion: v1.CommandUnion{
						Exec: &v1.ExecCommand{},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DevfileV2{
				v1.Devfile{
					DevWorkspaceTemplateSpec: v1.DevWorkspaceTemplateSpec{
						DevWorkspaceTemplateSpecContent: v1.DevWorkspaceTemplateSpecContent{
							Commands: tt.currentCommands,
						},
					},
				},
			}

			got := d.AddCommands(tt.newCommands)
			if !tt.wantErr && got != nil {
				t.Errorf("TestDevfile200_AddCommands() unexpected error - %v", got)
			} else if tt.wantErr && got == nil {
				t.Errorf("TestDevfile200_AddCommands() wanted an error but got nil")
			}
		})
	}
}

func TestDevfile200_UpdateCommands(t *testing.T) {

	type args struct {
		name string
	}
	tests := []struct {
		name            string
		currentCommands []v1.Command
		newCommand      v1.Command
	}{
		{
			name: "case 1: update the command",
			currentCommands: []v1.Command{
				{
					Id: "command1",
					CommandUnion: v1.CommandUnion{
						Exec: &v1.ExecCommand{
							Component: "component1",
						},
					},
				},
				{
					Id: "command2",
					CommandUnion: v1.CommandUnion{
						Composite: &v1.CompositeCommand{},
					},
				},
			},
			newCommand: v1.Command{
				Id: "command1",
				CommandUnion: v1.CommandUnion{
					Exec: &v1.ExecCommand{
						Component: "component1new",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DevfileV2{
				v1.Devfile{
					DevWorkspaceTemplateSpec: v1.DevWorkspaceTemplateSpec{
						DevWorkspaceTemplateSpecContent: v1.DevWorkspaceTemplateSpecContent{
							Commands: tt.currentCommands,
						},
					},
				},
			}

			d.UpdateCommand(tt.newCommand)

			commands, err := d.GetCommands(common.DevfileOptions{})
			if err != nil {
				t.Errorf("TestDevfile200_UpdateCommands() unxpected error %v", err)
				return
			}

			matched := false
			for _, devfileCommand := range commands {
				if tt.newCommand.Id == devfileCommand.Id {
					matched = true
					if !reflect.DeepEqual(devfileCommand, tt.newCommand) {
						t.Errorf("TestDevfile200_UpdateCommands() command mismatch - wanted %+v, got %+v", tt.newCommand, devfileCommand)
					}
				}
			}

			if !matched {
				t.Errorf("TestDevfile200_UpdateCommands() command mismatch - did not find command with id %s", tt.newCommand.Id)
			}
		})
	}
}

func TestDeleteCommands(t *testing.T) {

	d := &DevfileV2{
		v1.Devfile{
			DevWorkspaceTemplateSpec: v1.DevWorkspaceTemplateSpec{
				DevWorkspaceTemplateSpecContent: v1.DevWorkspaceTemplateSpecContent{
					Commands: []v1.Command{
						{
							Id: "command1",
							CommandUnion: v1.CommandUnion{
								Exec: &v1.ExecCommand{},
							},
						},
						{
							Id: "command2",
							CommandUnion: v1.CommandUnion{
								Exec: &v1.ExecCommand{},
							},
						},
						{
							Id: "command3",
							CommandUnion: v1.CommandUnion{
								Composite: &v1.CompositeCommand{
									Commands: []string{"command1", "command2", "command1"},
								},
							},
						},
					},
				},
			},
		},
	}

	tests := []struct {
		name            string
		commandToDelete string
		wantCommands    []v1.Command
		wantErr         bool
	}{
		{
			name:            "Successfully delete command",
			commandToDelete: "command1",
			wantCommands: []v1.Command{
				{
					Id: "command2",
					CommandUnion: v1.CommandUnion{
						Exec: &v1.ExecCommand{},
					},
				},
				{
					Id: "command3",
					CommandUnion: v1.CommandUnion{
						Composite: &v1.CompositeCommand{
							Commands: []string{"command1", "command2", "command1"},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name:            "Missing Command",
			commandToDelete: "command34",
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := d.DeleteCommand(tt.commandToDelete)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err == nil {
				assert.Equal(t, tt.wantCommands, d.Commands, "The two values should be the same.")
			}
		})
	}

}
