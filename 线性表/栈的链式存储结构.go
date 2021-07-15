package main

import (
	"errors"
	"fmt"
)

type stackNode struct {
	data interface{}
	next *stackNode
}
type linkStack struct {
	top *stackNode
	count int
}

func (c *linkStack) push(data interface{}) {
	var stackNode = stackNode{
		data:      data,
		next: c.top,
	}
	c.top = &stackNode
	c.count++
}

func (c *linkStack) pop() (interface{}, error) {
	if c.top == nil {
		return nil, errors.New("无数据")
	}
	p := c.top
	c.top = c.top.next
	c.count++
	return p.data, nil
}


func main() {
	var linkStack = linkStack{}
	linkStack.push("6666")
	linkStack.push("77878")
	data, err := linkStack.pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

}
