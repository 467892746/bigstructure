package main

import "fmt"

type EdgeNode struct {
	adjVex int
	weight int
	next *EdgeNode
}

type VertexNode struct {
	data interface{}
	firstEdge *EdgeNode
}

type GraphAdjList struct {
	adjList []VertexNode
	numVertexes, numEdges int
}

func CreateALGraph(g *GraphAdjList)  {
	var i,j int

	fmt.Println("输入顶点数和边数")
	_, err := fmt.Scan(&g.numVertexes, &g.numEdges)
	if err != nil {
		panic(err)
	}
	g.adjList = make([]VertexNode, g.numVertexes)
	// 读入顶点信息
	for i := 0; i < g.numVertexes; i++ {
		_, err = fmt.Scan(&g.adjList[i].data)
		if err != nil {
			panic(err)
		}

	}
	// 建立边表
	for k := 0; k < g.numEdges; k++ {
		fmt.Println("输入边(vi, vj)上的顶点序号:")
		fmt.Scan(&i,&j)
		e := EdgeNode{
			adjVex: j,
			next:   nil,
		}
		e.next = g.adjList[i].firstEdge
		g.adjList[i].firstEdge = &e

	}
}

func main() {
	
}
