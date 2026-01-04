package runtime

import "io"

type IOStreams struct {
	In  io.Reader
	Out io.Writer
	Err io.Writer
}
