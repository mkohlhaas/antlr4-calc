package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	fmt.Println("vim-go")
	antlr.NewFileStream("hello.txt")
}
