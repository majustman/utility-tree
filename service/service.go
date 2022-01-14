package service

import (
	"fmt"
	"github.com/majustman/utility-tree/tree"
	"strings"
)

func Run(address string, filesParam bool) error {
	t, err := tree.NewTree(address)
	if err != nil {
		return err
	}
	var b strings.Builder
	writeTreeIntoBuilder(t, filesParam, 0, &b)
	fmt.Println(b.String())
	return nil
}

func writeTreeIntoBuilder(t *tree.Tree, fileParam bool, level int, b *strings.Builder) {
	if level == 0 {	fmt.Fprintln(b, t.Name)	}
	if fileParam {
		writeFiles(t.Files, level, len(t.Dirs), b)
	}
	for i, d := range t.Dirs {
		writeSpace(level, b)
		if i != len(t.Dirs)-1 {
			fmt.Fprintf(b,"├── %v\n", d.Name)
		} else {
			fmt.Fprintf(b,"└── %v\n", d.Name)
		}
		writeTreeIntoBuilder(d, fileParam, level+1, b)
	}
}

func writeSpace(level int, b *strings.Builder) {
	b.Grow(level*4)
	for i := 1; i <= level; i++ {
		fmt.Fprint(b,"│   ")
	}
}

func writeFiles(files []string, level int, lenDir int, b *strings.Builder) {
	for i, f := range files {
		writeSpace(level, b)
		b.Grow(len(f)+4)
		if i != len(files)-1 || lenDir > 0 {
			fmt.Fprintf(b,"├── %v\n", f)
		} else {
			fmt.Fprintf(b,"└── %v\n", f)
		}
	}
}
