package main

import "fmt"

type ToPoEdgeNode struct {
	adjVex int
	weight int
	next *ToPoEdgeNode
}

type ToPoVertexNode struct {
	in int
	data int
	firstEdge *ToPoEdgeNode
}

type ToPoGraphAdjList struct {
	adjList [MAXVEX]ToPoVertexNode
	numVertexes,numEdge int
}

func ToPoLogicalSort(gl ToPoGraphAdjList)  {
	var linkStack = linkStack{}


	for i := 0; i < gl.numVertexes; i ++ {
		if gl.adjList[i].in == 0 {
			linkStack.push(i)
		}
	}
	var count int
	for  j, err := linkStack.pop(); err == nil; {
		gettop := j.(int)
		fmt.Println(gettop)
		count++
		for e := gl.adjList[j].firstEdge; e != nil; e = e.next {
			k := e.adjVex
			gl.adjList[k].in --
			if  gl.adjList[k].in == 0 {
				linkStack.push(k)
			}
		}
	}
}
