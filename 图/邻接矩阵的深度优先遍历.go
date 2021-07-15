package main

import "fmt"

var visited  [MAXVEX]bool

func mDfs(g MGraph, i int)  {
	visited[i] = true
	fmt.Println(g.vexs[i])
	for j := 0; j < g.numVertexes; j ++{
		if g.arc[i][j] == 1 && !visited[j] {
			mDfs(g, j)
		}
	}
}

func mDfsTraverse(g MGraph)  {
	for i := 0; i < g.numVertexes; i ++ {
		if !visited[i] {
			mDfs(g, i)
		}
	}
}