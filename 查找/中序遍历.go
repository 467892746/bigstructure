package main

import "fmt"

func InOrderTraverse(t *BiTNode)  {
	if t == nil {
		return
	}
	InOrderTraverse(t.lChid)
	fmt.Println(t.data)
	InOrderTraverse(t.rChid)
}


