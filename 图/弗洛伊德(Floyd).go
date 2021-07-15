package main

func ShortestPathFloyd(g MGraph)  {
	var p, d [MAXVEX][MAXVEX]int
	for v := 0; v < g.numVertexes; v ++ {
		for w := 0; w < g.numVertexes; w++ {
			d[v][w] = g.arc[v][w]
			p[v][w] = w
		}
	}

	for k := 0; k < g.numVertexes; k++ {
		for v := 0; v < g.numVertexes; v++ {
			for w := 0; w < g.numVertexes; w++ {
				if d[v][w] > d[v][k] + d[k][w] {
					d[v][w] = d[v][k] + d[k][w]
					p[v][w] = p[v][k]
				}
			}
		}
	}
}