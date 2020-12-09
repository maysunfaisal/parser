package parser

import "github.com/maysunfaisal/parser/pkg/testingutil/filesystem"

// GetFs returns the filesystem object
func (d *DevfileCtx) GetFs() filesystem.Filesystem {
	return d.fs
}
