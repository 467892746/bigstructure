package main

import "fmt"

func MiniSpanTreePrim(g MGraph)  {
	var adjvex [MAXVEX]int 
	var lowcost [MAXVEX]int
	for i := 1; i < g.numVertexes; i++ {
		lowcost[i] = g.arc[0][i]
	}
	for i := 1; i < g.numVertexes; i++ {
		min := 65535
		j := 1
		k := 0
		for j < g.numVertexes {
			if lowcost[j] != 0 && lowcost[j] < min {
				min = lowcost[j]
				k = j
			}
			j ++
		}
		fmt.Println(adjvex[k], k)
		lowcost[k] = 0
		for j := 1; j < g.numVertexes; j++ {
			if lowcost[j] != 0 && g.arc[i][j] < lowcost[j] {
				lowcost[j] = g.arc[i][j] // 将较小权值存入lowcost
				adjvex[j] = k // 将下标为k的顶点存入adjvex
			}
		}
	}
}
