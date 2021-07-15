package main

import "fmt"

type MultipleEdgeNode struct {
	iVex int
	iLink *MultipleEdgeNode
	jVex int
	jLink *MultipleEdgeNode
}

type MultipleVertexNode struct {
	data interface{}
	firstEdge *MultipleEdgeNode
}

type GraphMultipleAdjList struct {
	adjList []MultipleVertexNode
	numVertexes, numEdges int
}

func CreateMultipleALGraph(g *GraphMultipleAdjList)  {
	var i,j int

	fmt.Println("输入顶点数和边数")
	_, err := fmt.Scan(&g.numVertexes, &g.numEdges)
	if err != nil {
		panic(err)
	}
	g.adjList = make([]MultipleVertexNode, g.numVertexes)
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
		_, err := fmt.Scan(&i,&j)
		if err != nil {
			panic(err)
		}
		e := &MultipleEdgeNode{
			iVex:i,
			jVex:j,
		}
		gIFirst := g.adjList[i].firstEdge

		if gIFirst == nil {
			g.adjList[i].firstEdge = e
		}else {
			e.iLink = gIFirst
			g.adjList[i].firstEdge = e
		}

		gJFirst := g.adjList[j].firstEdge
		if gJFirst == nil {
			g.adjList[j].firstEdge = e
		}else {
			e.jLink = gJFirst
			g.adjList[j].firstEdge = e
		}

	}
}

func main() {
	
}
