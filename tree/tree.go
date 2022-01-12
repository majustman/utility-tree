package tree

import (
	"os"
	"strings"
)

type Tree struct {
	Name string
	Address string
	Dirs []*Tree
	Files []string
}

func NewTree(address string) (*Tree, error) {
	var t Tree
	tmpS := strings.Split(address, "/")
	t.Name = tmpS[len(tmpS)-1]
	dirEntries, err := os.ReadDir(address)
	if err != nil {
		return nil, err
	}
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			tmp, err := NewTree(address + "/" + dirEntry.Name())
			if err == nil { t.Dirs = append(t.Dirs, tmp) }
		} else {
			t.Files = append(t.Files, dirEntry.Name())
		}
	}
	return &t, nil
}
