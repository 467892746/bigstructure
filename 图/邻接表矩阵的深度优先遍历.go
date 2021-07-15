package main

import (
	"fmt"

)

func adjDfs(gl GraphAdjList, i int)  {
	visited[i] = true
	fmt.Println(gl.adjList[i].data)
	p := gl.adjList[i].firstEdge
	for p != nil {
		if !visited[p.adjVex] {
			adjDfs(gl, p.adjVex)
		}
		p = p.next
	}
}

func adjDfsTraverse(gl GraphAdjList)  {
	for i := 0; i < gl.numVertexes; i++ {
		if !visited[i]{
			adjDfs(gl, i)
		}
	}

}