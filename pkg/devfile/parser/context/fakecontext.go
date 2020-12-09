package parser

import "github.com/maysunfaisal/parser/pkg/testingutil/filesystem"

func FakeContext(fs filesystem.Filesystem, absPath string) DevfileCtx {
	return DevfileCtx{
		fs:      fs,
		absPath: absPath,
	}
}
