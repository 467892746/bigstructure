package main

import (
	"errors"
	"fmt"
)

const  maxSize = 10

type sqStack struct {
	data [maxSize]interface{}
	top int
}

func (c *sqStack) push(e interface{}) error {
	if c.top == maxSize - 1 {
		return errors.New("栈满")
	}
	c.top++
	c.data[c.top] = e
	return nil
}

func (c *sqStack) pop() (interface{}, error) {
	if c.top == -1 {
		return nil, errors.New("栈为空")
	}
	c.top--
	return c.data[c.top], nil
}

type SqDoubleStack struct {
	data [maxSize]interface{}
	top1 int
	top2 int
}

func NewSqDoubleStack() SqDoubleStack {
	return SqDoubleStack{
		data: [10]interface{}{},
		top1: -1,
		top2: maxSize ,
	}
}

func (c *SqDoubleStack) push(e interface{}, stackNumber int) error {
	if c.top1+1 == c.top2 {
		return errors.New("栈已满")
	}
	if stackNumber ==  1 {
		c.top1++
		c.data[c.top1] = e
	}else if stackNumber == 2 {
		c.top2--
		c.data[c.top2] = e
	}else {
		return  errors.New("没有此栈")
	}
	return nil
}

func (c *SqDoubleStack) pop(stackNumber int) (interface{}, error) {
	switch stackNumber {
	case 1:
		if c.top1 == -1 {
			return nil, errors.New("此栈已无元素")
		}
		defer func() {
			c.top1--
		}()
		return c.data[c.top1], nil
	case 2:
		if c.top1 == maxSize {
			return nil, errors.New("此栈已无元素")
		}
		defer func() {
			c.top2++
		}()
		return c.data[c.top2], nil
	default:
		return nil, errors.New("没有此栈")
	}
}

func main() {
	sqDoubleStack := NewSqDoubleStack()
	sqDoubleStack.push(5, 1)
	sqDoubleStack.push(5, 1)
	sqDoubleStack.push(5, 1)
	sqDoubleStack.push(5, 1)
	sqDoubleStack.push(5, 1)
	sqDoubleStack.push(5, 1)
	sqDoubleStack.push(5, 1)
	sqDoubleStack.push(5, 1)
	sqDoubleStack.push(5, 1)
	sqDoubleStack.push(5, 1)
	fmt.Println(sqDoubleStack.pop(1))
	fmt.Println(sqDoubleStack.pop(1))
	fmt.Println(sqDoubleStack.pop(1))
	fmt.Println(sqDoubleStack.pop(1))
	fmt.Println(sqDoubleStack.pop(1))
	fmt.Println(sqDoubleStack.pop(1))
	fmt.Println(sqDoubleStack.pop(1))
	fmt.Println(sqDoubleStack.pop(1))
	fmt.Println(sqDoubleStack.pop(1))
	fmt.Println(sqDoubleStack.pop(1))
	fmt.Println(sqDoubleStack.pop(1))
}
