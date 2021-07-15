package main

import (
	"errors"
	"fmt"
)

var etv,ltv []int
var stack2 []int
var top2 int

func CriticalToPoLogicalSort(gl ToPoGraphAdjList) error {
	var stack []int
	var top int
	var count int
	for i := 0; i < gl.numVertexes; i++ {
		if gl.adjList[i].in == 0{
			stack = append(stack, i)
			top++
		}
	}
	etv = make([]int, gl.numVertexes)
	for top != 0 {
		top--
		gettop := stack[top]
		count++
		top2++
		stack2 = append(stack2, gettop)
		for e := gl.adjList[gettop].firstEdge; e != nil; e = e.next {
			k := e.adjVex
			gl.adjList[k].in--
			if gl.adjList[k].in == 0 {
				top++
				stack[top] = k
			}
			if etv[gettop] + e.weight > etv[k] {
				etv[k] = etv[gettop] + e.weight
			}
		}
	}
	if count < gl.numVertexes {
		return errors.New("存在环")
	}
	return nil
}

func CriticalPath(gl ToPoGraphAdjList)  {
	err := CriticalToPoLogicalSort(gl)
	if err != nil {
		panic(err)
	}
	for i := 0; i < gl.numVertexes; i++ {
		ltv[i] = etv[gl.numVertexes - 1]
	}
	if top2 != 0 {
		top2--
		getTop := stack2[top2]
		for e := gl.adjList[getTop].firstEdge; e != nil; e = e.next {
			k := e.adjVex
			if ltv[k] - e.weight < ltv[getTop] {
				ltv[getTop] = ltv[k] - e.weight
			}
		}
	}
	for j := 0; j < gl.numVertexes; j++ {
		for e := gl.adjList[j].firstEdge; e != nil; e = e.next {
			k := e.adjVex
			ete := etv[j]
			lte := ltv[k] - e.weight
			if ete == lte {
				fmt.Println(gl.adjList[j].data, gl.adjList[k].data, e.weight)
			}
		}
	}
}