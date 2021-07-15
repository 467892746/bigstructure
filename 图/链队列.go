// 方便使用

package main

import (
	"errors"
	"fmt"
)

type QNode struct {
	data interface{}
	next *QNode
}

type LinkQueue struct {
	front *QNode
	rear *QNode
	
}

func NewLinkQueue() LinkQueue {
	node := &QNode{ // 头结点
		data: nil,
		next: nil,
	}
		return LinkQueue{
			front: node,
			rear:  node,
		}
}

func (c *LinkQueue) EnQueue(data interface{})  {
	node :=  &QNode{
		data: data,
		next: nil,
	}
	c.rear.next = node
	c.rear = node
}

func (c *LinkQueue) DeQueue() ( interface{}, error) {
	if c.rear == c.front {
		return nil, errors.New("链队列为空")
	}
	p := c.front.next
	data := p.data
	c.front.next = p.next
	if c.rear == p {
		c.rear = c.front
	}
	return data, nil
}


func main() {
	l := NewLinkQueue()
	l.EnQueue("11111111111")
	l.EnQueue("22222222222")
	fmt.Println(l.DeQueue())
	fmt.Println(l.DeQueue())
}
