package funcs

import (
	"fmt"
	"os"

	"github.com/haya14busa/offset"
)

func Pos() PosFuncs {
	return PosFuncs{}
}

type PosFuncs struct{}

func (PosFuncs) Offset(file string, byteOffset interface{}) offset.Position {
	p, err := offset.FromFilename(file, toInt(byteOffset))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return p
}
