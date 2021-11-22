package ehtml

import (
	"golang.org/x/net/html"
	"io"
)

type reader struct {
	s string // raw read data
	i int    // reading offset
}

func (r *reader) Read(p []byte) (int, error) {
	remain := len(r.s) - r.i
	if remain == 0 {
		return 0, io.EOF
	}
	if remain <= len(p) {
		copy(p, r.s[r.i:])
		r.i += remain
		return remain, nil
	}

	size := len(p)
	copy(p, r.s[r.i:r.i+size])
	r.i += size
	return size, nil
}

func NewReader(p string) *reader {
	return &reader{s: p, i: 0}
}

type limitReader struct {
	io.Reader
	remain int64
}

func (r *limitReader) Read(p []byte) (int, error) {
	if r.remain <= 0 {
		return 0, io.EOF
	}

	if r.remain < int64(len(p)) {
		p = p[0:r.remain]
	}
	n, err := r.Read(p)
	r.remain -= int64(n)
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{Reader: r, remain: n}
}

func Parse(r io.Reader) (*html.Node, error) {
	return html.Parse(r)
}
