package main

import "fmt"

func InOrderTraverse(t *BiTNode)  {
	if t == nil {
		return
	}
	InOrderTraverse(t.lChild)
	fmt.Println(t.data)
	InOrderTraverse(t.rChild)
}

func main() {
	
}
