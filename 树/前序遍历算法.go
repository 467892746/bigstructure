package main

import "fmt"

func PreOrderTraverse(t *BiTNode)  {
	if t == nil {
		return
	}
	fmt.Println(t.data)
	PreOrderTraverse(t.lChild)
	PreOrderTraverse(t.rChild)
}

func main() {
	
}
