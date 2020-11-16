package main

import (
	"io"
)

type byteCounter struct {
	writer  io.Writer
	written int64
}

func (b *byteCounter) Write(p []byte) (int, error) {
	n, err := b.writer.Write(p)
	if err != nil {
		return n, err
	}
	b.written += int64(n)
	return n, nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := byteCounter{w, 0}
	return c.writer, &c.written
}
