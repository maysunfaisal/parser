package common

import (
	"reflect"
	"testing"

	v1 "github.com/maysunfaisal/api/v2/pkg/apis/workspaces/v1alpha2"
)

func TestGetGroup(t *testing.T) {

	tests := []struct {
		name    string
		command v1.Command
		want    *v1.CommandGroup
	}{
		{
			name: "Exec command group",
			command: v1.Command{
				Id: "exec1",
				CommandUnion: v1.CommandUnion{
					Exec: &v1.ExecCommand{
						LabeledCommand: v1.LabeledCommand{
							BaseCommand: v1.BaseCommand{
								Group: &v1.CommandGroup{
									IsDefault: true,
									Kind:      v1.RunCommandGroupKind,
								},
							},
						},
					},
				},
			},
			want: &v1.CommandGroup{
				IsDefault: true,
				Kind:      v1.RunCommandGroupKind,
			},
		},
		{
			name: "Composite command group",
			command: v1.Command{
				Id: "composite1",
				CommandUnion: v1.CommandUnion{
					Composite: &v1.CompositeCommand{
						LabeledCommand: v1.LabeledCommand{
							BaseCommand: v1.BaseCommand{
								Group: &v1.CommandGroup{
									IsDefault: true,
									Kind:      v1.BuildCommandGroupKind,
								},
							},
						},
					},
				},
			},
			want: &v1.CommandGroup{
				IsDefault: true,
				Kind:      v1.BuildCommandGroupKind,
			},
		},
		{
			name:    "Empty command",
			command: v1.Command{},
			want:    nil,
		},
		{
			name: "Apply command group",
			command: v1.Command{
				Id: "apply1",
				CommandUnion: v1.CommandUnion{
					Apply: &v1.ApplyCommand{
						LabeledCommand: v1.LabeledCommand{
							BaseCommand: v1.BaseCommand{
								Group: &v1.CommandGroup{
									IsDefault: true,
									Kind:      v1.TestCommandGroupKind,
								},
							},
						},
					},
				},
			},
			want: &v1.CommandGroup{
				IsDefault: true,
				Kind:      v1.TestCommandGroupKind,
			},
		},
		{
			name: "Custom command group",
			command: v1.Command{
				Id: "custom1",
				CommandUnion: v1.CommandUnion{
					Custom: &v1.CustomCommand{
						LabeledCommand: v1.LabeledCommand{
							BaseCommand: v1.BaseCommand{
								Group: &v1.CommandGroup{
									IsDefault: true,
									Kind:      v1.BuildCommandGroupKind,
								},
							},
						},
					},
				},
			},
			want: &v1.CommandGroup{
				IsDefault: true,
				Kind:      v1.BuildCommandGroupKind,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			commandGroup := GetGroup(tt.command)
			if !reflect.DeepEqual(commandGroup, tt.want) {
				t.Errorf("expected %v, actual %v", tt.want, commandGroup)
			}
		})
	}

}

func TestGetExecComponent(t *testing.T) {

	tests := []struct {
		name    string
		command v1.Command
		want    string
	}{
		{
			name: "Case 1: Exec component present",
			command: v1.Command{
				Id: "exec1",
				CommandUnion: v1.CommandUnion{
					Exec: &v1.ExecCommand{
						Component: "component1",
					},
				},
			},
			want: "component1",
		},
		{
			name: "Case 2: Exec component absent",
			command: v1.Command{
				Id: "exec1",
				CommandUnion: v1.CommandUnion{
					Exec: &v1.ExecCommand{},
				},
			},
			want: "",
		},
		{
			name:    "Case 3: Empty command",
			command: v1.Command{},
			want:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := GetExecComponent(tt.command)
			if component != tt.want {
				t.Errorf("expected %v, actual %v", tt.want, component)
			}
		})
	}

}

func TestGetExecCommandLine(t *testing.T) {

	tests := []struct {
		name    string
		command v1.Command
		want    string
	}{
		{
			name: "Case 1: Exec command line present",
			command: v1.Command{
				Id: "exec1",
				CommandUnion: v1.CommandUnion{
					Exec: &v1.ExecCommand{
						CommandLine: "commandline1",
					},
				},
			},
			want: "commandline1",
		},
		{
			name: "Case 2: Exec command line absent",
			command: v1.Command{
				Id: "exec1",
				CommandUnion: v1.CommandUnion{
					Exec: &v1.ExecCommand{},
				},
			},
			want: "",
		},
		{
			name:    "Case 3: Empty command",
			command: v1.Command{},
			want:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			commandLine := GetExecCommandLine(tt.command)
			if commandLine != tt.want {
				t.Errorf("expected %v, actual %v", tt.want, commandLine)
			}
		})
	}

}

func TestGetExecWorkingDir(t *testing.T) {

	tests := []struct {
		name    string
		command v1.Command
		want    string
	}{
		{
			name: "Case 1: Exec working dir present",
			command: v1.Command{
				Id: "exec1",
				CommandUnion: v1.CommandUnion{
					Exec: &v1.ExecCommand{
						WorkingDir: "workingdir1",
					},
				},
			},
			want: "workingdir1",
		},
		{
			name: "Case 2: Exec working dir absent",
			command: v1.Command{
				Id: "exec1",
				CommandUnion: v1.CommandUnion{
					Exec: &v1.ExecCommand{},
				},
			},
			want: "",
		},
		{
			name:    "Case 3: Empty command",
			command: v1.Command{},
			want:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			workingDir := GetExecWorkingDir(tt.command)
			if workingDir != tt.want {
				t.Errorf("expected %v, actual %v", tt.want, workingDir)
			}
		})
	}

}

func TestGetCommandType(t *testing.T) {

	tests := []struct {
		name        string
		command     v1.Command
		wantErr     bool
		commandType v1.CommandType
	}{
		{
			name: "Exec command",
			command: v1.Command{
				Id: "exec1",
				CommandUnion: v1.CommandUnion{
					Exec: &v1.ExecCommand{},
				},
			},
			commandType: v1.ExecCommandType,
			wantErr:     false,
		},
		{
			name: "Composite command",
			command: v1.Command{
				Id: "comp1",
				CommandUnion: v1.CommandUnion{
					Composite: &v1.CompositeCommand{},
				},
			},
			commandType: v1.CompositeCommandType,
			wantErr:     false,
		},
		{
			name: "Apply command",
			command: v1.Command{
				Id: "apply1",
				CommandUnion: v1.CommandUnion{
					Apply: &v1.ApplyCommand{},
				},
			},
			commandType: v1.ApplyCommandType,
			wantErr:     false,
		},
		{
			name: "Custom command",
			command: v1.Command{
				Id: "custom",
				CommandUnion: v1.CommandUnion{
					Custom: &v1.CustomCommand{},
				},
			},
			commandType: v1.CustomCommandType,
			wantErr:     false,
		},
		{
			name: "Unknown command",
			command: v1.Command{
				Id:           "unknown",
				CommandUnion: v1.CommandUnion{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCommandType(tt.command)
			// Unexpected error
			if (err != nil) != tt.wantErr {
				t.Errorf("TestGetCommandType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.commandType {
				t.Errorf("TestGetCommandType error: command type mismatch, expected: %v got: %v", tt.commandType, got)
			}
		})
	}

}
