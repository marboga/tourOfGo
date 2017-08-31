package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(s []byte) (int, error) {
	len, e := rr.r.Read(s)

	for i := 0; i < len; i++ {
		if s[i] > 'a' && s[i] <= 'z' {
			s[i] += 13
			if s[i] > 'z' {
				s[i] -= 26
			}
		} else if s[i] > 'A' && s[i] <= 'Z' {
			s[i] += 13
			if s[i] > 'Z' {
				s[i] -= 26
			}
		}
	}
	return len, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
