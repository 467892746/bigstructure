package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Data interface{}
	next  *Node
}

func (c *Node) GetElem(i int) (*Node, error) {
	p := c.next

	j := 1
	for p != nil && j < i {
		p = p.next
		j++
	}
	if p == nil || j > i {
		return nil, errors.New("第i个元素不存在")
	}
	return p, nil
}

func (c *Node) ListInsert(i int, data interface{}) (*Node, error) {
	if c.next == nil && i == 0 {
		c.next = &Node{
			Data: data,
			next: nil,
		}
		return c.next, nil
	}
	iNode, err := c.GetElem(i)
	if err != nil {
		return nil, err
	}
	newNode := Node{
		Data: data,
		next: iNode.next,
	}
	iNode.next = &newNode
	return  &newNode, nil
}

func main() {

	headNode := Node{
		Data: nil,
		next: nil,
	}

	newNode, err := headNode.ListInsert(0, "哈哈")
	if err != nil {
		panic(err)
	}
	newNode, err = headNode.ListInsert(1, "哈哈11")
	if err != nil {
		panic(err)
	}
	fmt.Println(headNode.next.Data)
	fmt.Println(newNode.Data)
}
