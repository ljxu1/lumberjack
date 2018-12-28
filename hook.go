package lumberjack

import (
	"os"
	"io"
)

type OnFileOpenHook func(f *os.File) io.WriteCloser
