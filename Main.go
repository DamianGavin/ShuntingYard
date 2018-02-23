package main

import (
	"fmt"
	"golang.org/x/tools/go/gcimporter15/testdata"
)

func intopost(infix string) string {
	postfix := ""

	return postfix
}
func main() {
	//answer should be ab.c*
	fmt.Print("Infix:   ", "a.s.c*")
	fmt.Print("Postfix  ", intopost("a.b.c*"))

	//answer should be abd|.*
	fmt.Print("Infix:   ", "a.(b|d))*")
	fmt.Print("Postfix  ", intopost("(a.(b|d))*"))

	//answer should be abd|.c*.
	fmt.Print("Infix:   ", "a.(b|d).c*")
	fmt.Print("Postfix: ", intopost("a.(b|d).c*"))

	//answer should be abb.+.c.
	fmt.Print("Infix    ", "a.(b.b)+.c")
	fmt.Print("Postfix: ", intopost("a.(b.b)+.c"))

}
