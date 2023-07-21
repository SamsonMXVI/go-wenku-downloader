package util

import (
	"io"

	"gopkg.in/cheggaaa/pb.v2"
)

type ProgressWriter struct {
	Writer       io.Writer
	ProgressBar  *pb.ProgressBar
	CurrentBytes int64
}

func (pw *ProgressWriter) Write(p []byte) (int, error) {
	n, err := pw.Writer.Write(p)
	if n > 0 {
		pw.CurrentBytes += int64(n)
		pw.ProgressBar.SetCurrent(pw.CurrentBytes)
	}
	return n, err
}
