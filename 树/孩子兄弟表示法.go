package main

import (
	"fmt"
	"unsafe"
)

type CSNode struct {
	data interface{}
	firstChild *CSNode
	rightSib *CSNode
}

func main() {
	var a *string = nil
	fmt.Println(unsafe.Sizeof(a))
}
