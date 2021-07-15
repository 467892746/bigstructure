package main

import (
	"errors"
	"fmt"
)

const LoopQueueMaxSize  = 2

type LoopQueue struct {
	front int // 头指针
	rear int // 尾指针
	data [LoopQueueMaxSize]int
}

func NewLoopQueue() *LoopQueue {
	return &LoopQueue{
		front: 0,
		rear:  0,
	}
}

func (l *LoopQueue) IsFull() bool {
	return (l.rear + 1) % LoopQueueMaxSize == l.front
}

func (l *LoopQueue) Length() int {
	return (l.rear + LoopQueueMaxSize - l.front) % LoopQueueMaxSize
}

func (l *LoopQueue) EnQueue(data int) error {
	if l.IsFull() {
		return errors.New("队列已满")
	}
	l.data[l.rear] = data
	l.rear = (l.rear + 1) %LoopQueueMaxSize
	return nil
}

func (l *LoopQueue) DeQueue() (int, error) {
	if l.front == l.rear {
		return 0, errors.New("队列为空")
	}
	data := l.data[l.front]
	l.front = (l.front + 1 ) %LoopQueueMaxSize
	return data, nil
}

func main() {
	queue := NewLoopQueue()
	err := queue.EnQueue(111111111)
	if err != nil {
		panic(err)
	}

	data, err := queue.DeQueue()
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	fmt.Println(queue.data)
}
