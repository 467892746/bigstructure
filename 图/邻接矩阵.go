package main

import "fmt"

const (
	MAXVEX  = 100
	INFINITY = 65535
)


type MGraph struct {
	vexs [MAXVEX]int
	arc  [MAXVEX][MAXVEX]int
	numVertexes, numEdges int
}

func CreateMGraph(g *MGraph)  {
	var i,j,k,w int
	fmt.Println("输入顶点数和边数:\n")
	_, err := fmt.Scanf("%d %d", &g.numVertexes, &g.numEdges)
	if err != nil {
		panic(err)
	}
	for i := 0; i < g.numVertexes; i++ {
		_, err = fmt.Scan(&g.vexs[i])
		if err != nil {
			panic(err)
		}
	}
	for i := 0; i < g.numVertexes; i++ {
		for j := 0; j < g.numVertexes; j++ {

			g.arc[i][j] = INFINITY
		}
	}
	for k = 0; k < g.numEdges; i++ {
		fmt.Println("输入边(vi,vj)上的下标i, 下标j和权w:")
		_, err = fmt.Scanf("%d %d %d", &i, &j, &w)
		if err != nil {
			panic(err)
		}
		g.arc[i][j] = w
		g.arc[j][i] = w // 因为是无向图，矩阵对称
	}

}

func main() {
	var a MGraph
	CreateMGraph(&a)
}
