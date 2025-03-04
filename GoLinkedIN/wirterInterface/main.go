package main

import (
	"fmt"
	"io"
	"os"
)

type Capper struct {
	wtr io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
	diff := byte('a' - 'A')
	for i := range p {
		if p[i] >= 'a' && p[i] <= 'z' {
			p[i] -= diff
		}
	}
	fmt.Println(p)
	return c.wtr.Write(p)
}

func main() {
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello There")
}
