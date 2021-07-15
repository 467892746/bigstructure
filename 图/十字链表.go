package main

import "fmt"

type CrossEdgeNode struct {
	tailVex int // 尾顶点
	headVex int // 头顶点
	headLink *CrossEdgeNode // 头顶点链表
	tailLink *CrossEdgeNode // 尾顶点链表
}

type CrossVertexNode struct {
	firstIn *CrossEdgeNode // 头顶点链表头部
	firstOut *CrossEdgeNode // 尾顶点链表尾部
	data interface{}
}

type CrossGraph struct {
	CrossList []CrossVertexNode
	numVertexes, numEdges int
}

func CreateCrossGraph(g *CrossGraph)  {
	var i,j int

	fmt.Println("输入顶点数和边数")
	_, err := fmt.Scan(&g.numVertexes, &g.numEdges)
	if err != nil {
		panic(err)
	}
	g.CrossList = make([]CrossVertexNode, g.numVertexes)
	// 读入顶点信息
	for i := 0; i < g.numVertexes; i++ {
		_, err = fmt.Scan(&g.CrossList[i].data)
		if err != nil {
			panic(err)
		}

	}
	// 建立边表
	for k := 0; k < g.numEdges; k++ {
		fmt.Println("输入边(vi, vj)上的顶点序号:")
		_, err = fmt.Scan(&i,&j)
		if err != nil {
			panic(err)
		}
		e := CrossEdgeNode{
			tailVex: i,
			headVex:j,
		}
		e.tailLink = g.CrossList[i].firstOut
		g.CrossList[i].firstOut = &e

		e.headLink = g.CrossList[j].firstIn

		g.CrossList[j].firstIn = &e


	}
}


func main() {
	
}
