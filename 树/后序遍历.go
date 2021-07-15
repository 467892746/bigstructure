package main

import "fmt"

func PostOrderTraverse(b *BiTNode)  {
	if b == nil {
		return
	}
	PostOrderTraverse(b.lChild)
	PostOrderTraverse(b.rChild)
	fmt.Println(b.data)
}

func main() {
	
}
