package service

import (
	"fmt"
	"github.com/majustman/utility-tree/tree"
)

func Run(address string, filesParam bool) error {
	t, err := tree.NewTree(address)
	if err != nil {
		return err
	}
	printTree(t, filesParam, 0)
	return nil
}

func printTree(t *tree.Tree, fileParam bool, level int) {
	if level == 0 {	fmt.Println(t.Name)	}
	if fileParam { printFiles(t.Files, level, len(t.Dirs)) }
	for i, d := range t.Dirs {
		printSpace(level)
		if i != len(t.Dirs)-1 {
			fmt.Printf("├── %v\n", d.Name)
		} else {
			fmt.Printf("└── %v\n", d.Name)
		}
		printTree(d, fileParam, level+1)
	}
}

func printSpace(level int) {
	for i := 1; i <= level; i++ {
		fmt.Print("│   ")
	}
}

func printFiles(files []string, level int, lenDir int) {
	for i, f := range files {
		printSpace(level)
		if i != len(files)-1 || lenDir > 0 {
			fmt.Printf("├── %v\n", f)
		} else {
			fmt.Printf("└── %v\n", f)
		}
	}
}
