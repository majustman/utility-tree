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
	writeTreeIntoBuilder(&b, t, filesParam, 0, []bool{true})
	fmt.Println(b.String())
	return nil
}

// The parameter levelSpaces is need to define what we should print on each level: just a spaces
// or '|'
func writeTreeIntoBuilder(b *strings.Builder, t *tree.Tree, fileParam bool, level int, levelSpaces []bool) {
	if level == 0 {	fmt.Fprintln(b, t.Name)	}
	if fileParam {
		writeFiles(b, t.Files, level, len(t.Dirs), levelSpaces)
	}
	for i, d := range t.Dirs {
		writeSpace(b, level, levelSpaces)
		newLevelSpaces := levelSpaces
		if i != len(t.Dirs)-1 {
			fmt.Fprintf(b,"├── %v\n", d.Name)
			newLevelSpaces = append(newLevelSpaces, true)
		} else {
			fmt.Fprintf(b,"└── %v\n", d.Name)
			newLevelSpaces = append(newLevelSpaces, false)
		}
		writeTreeIntoBuilder(b, d, fileParam, level+1, newLevelSpaces)
	}
}

// The parameter levelSpaces is need to define what we should print on each level: just a spaces
// or '|'
func writeSpace(b *strings.Builder, level int, levelSpaces []bool) {
	b.Grow(level*4)
	for i := 1; i <= level; i++ {
		if levelSpaces[i] {
			fmt.Fprint(b,"│   ")
		} else {
			fmt.Fprint(b,"    ")
		}
	}
}

// The parameter levelSpaces is need to define what we should print on each level: just a spaces
// or '|'
func writeFiles(b *strings.Builder, files []string, level int, lenDir int, levelSpaces []bool) {
	for i, f := range files {
		writeSpace(b, level, levelSpaces)
		b.Grow(len(f)+4)
		if i != len(files)-1 || lenDir > 0 {
			fmt.Fprintf(b,"├── %v\n", f)
		} else {
			fmt.Fprintf(b,"└── %v\n", f)
		}
	}
}
