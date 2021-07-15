package main

import (
	"fmt"
	"testing"
)

func TestDelete(t *testing.T) {
		a := []int{62,88,58,47,35,73,51,99,37,93}
		var startNode = new(*BiTNode)
		for _, b := range a{
			InsertBst(startNode, b)
		}
		InOrderTraverse(*startNode)
		fmt.Println("--------------------------")
		DeleteBst(*startNode, 62)
	InOrderTraverse(*startNode)

}
