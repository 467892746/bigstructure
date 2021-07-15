package main

import "fmt"

func GlBfsTraverse(gl GraphAdjList)  {

	q := NewLinkQueue()
	for i := 0; i < gl.numVertexes; i++  {
		if !visited[i] {
			visited[i] = true
			fmt.Println(gl.adjList[i].data)
			q.EnQueue(i)
			for q.front != q.rear {
				data, _ := q.DeQueue()
				i := data.(int)
				p := gl.adjList[i].firstEdge
				for p != nil {
					if !visited[p.adjVex] {
						visited[p.adjVex] = true
						fmt.Println(gl.adjList[p.adjVex].data)
						q.EnQueue(p.adjVex)
					}
					p = p.next
				}
			}
		}
	}
}
