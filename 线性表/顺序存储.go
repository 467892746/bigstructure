package main

import (
	"errors"
	"fmt"
)

type Array struct {
	array []int
	len int
	cap int
}

func  Init(len int, cap int) (*Array, error) {
	if len > cap {
		return nil, errors.New("len不能大于cap")
	}
	return &Array{
		array: make([]int, len, cap),
		len:len,
		cap:cap,
	},nil

}

func (c *Array) ListInsert(i int, e int) error {
	if c.cap == c.len {
		return errors.New("线性表已满")
	}
	if i < 1 || i > c.len + 1 {
		return errors.New("i不在范围")
	}
	if i <= c.len  {
		c.array = append(c.array, c.array[c.len - 1])
		for k := c.len -2; k >= i - 1; k-- {
			c.array[k + 1] = c.array[k]
		}
		c.array[i-1] = e
	}else {
		c.array = append(c.array, e)
	}
	c.len ++
	return nil
}

func main() {
	 arr ,err  := Init(9, 10)
	if err != nil {
		panic(err)
	}
	err = arr.ListInsert(1, 20)
	if err != nil {
		panic(err)
	}
	fmt.Println(arr.array)
}
