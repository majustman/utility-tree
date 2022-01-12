package main

import (
	"fmt"
	"github.com/majustman/utility-tree/service"
	"os"
)

func main() {
	args := os.Args[1:]
	l := len(args)
	if l == 0 {
		fmt.Println(fmt.Errorf("not enough arguments: %v", len(args)))
		return
	}
	if l > 2 {
		fmt.Println(fmt.Errorf("too much arguments: %v", len(args)))
		return
	}
	if l == 2 && args[0] == "-f" {
		err := service.Run(args[1], true)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	if l == 2 && args[0] != "-f" {
		fmt.Println(fmt.Errorf("unknown %v option", args[0]))
		return
	}
	err := service.Run(args[0], false)
	if err != nil {
		fmt.Println(err)
	}
}
