package main

import "io"

type WrapperWriter struct {
	io.Writer
	counts int64
}

func (w *WrapperWriter) Write(p []byte) (n int, err error) {
	w.counts += int64(len(p))
	return w.Write(p)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var wrapper = WrapperWriter{Writer: w}
	return &wrapper, &(wrapper.counts)
}
