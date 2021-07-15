package main

import "fmt"
//Edge
//！！！ begin一定要比end小才行
type Edge struct {
	begin int
	end int
	weight int
}

func find(parent [MAXVEX]int, f int) int {
	for parent[f] > 0 {
		f = parent[f]
	}
	return f
}

func MinSpanTreeKruskal(g MGraph)  {
	var parent [MAXVEX]int
	var edges [MAXVEX]Edge
	for i := 0; i < g.numEdges; i ++ {
		n := find(parent, edges[i].begin)
		m := find(parent, edges[i].end)
		if n != m {
			parent[n] = m
			fmt.Println(edges[i].begin, edges[i].end, edges[i].weight)
		}
	}
}
