package main

//ShortestPathDijkstra
//v0代表从Vo顶点出发
func ShortestPathDijkstra(g MGraph, v0 int)  {
	var final [MAXVEX]bool // 表示求得v0到vw的最短路径
	var d [MAXVEX]int // 用于存储到各点最短路径的权值和 d[v]表示V0到V的最短路径长度和
	var p [MAXVEX]int // 用于存储最短路径下标的数组 p[v]的值为前驱顶点下标
	var k int
	for v := 0; v < g.numVertexes; v++ {
		d[v] = g.arc[v0][v]
	}
	final[v0] = true
	for v := 1; v < g.numVertexes; v++ {
		min := 65535
		for w := 0; w < g.numVertexes; w++ {
			if !final[w] && d[w] < min {
				k = w
				min = d[w]
			}
		}
		final[k] = true
		for w := 0; w < g.numVertexes; w++ {
			if !final[w] && (min + g.arc[k][w] < d[w]) {
				d[w] = min + g.arc[k][w]
				p[w] = k
			}
		}
	}
}