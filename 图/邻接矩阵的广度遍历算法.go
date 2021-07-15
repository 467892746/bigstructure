package main

import "fmt"

func mBfsTraverse(g MGraph)  {
	q := NewLinkQueue()
	for i := 0; i < g.numVertexes; i++ {
		if !visited[i] {
			visited[i] = true
			fmt.Println(g.vexs[i])
			q.EnQueue(i) // 防止不是连通图也能全遍历完
			for q.front != q.rear{
				data, _ := q.DeQueue()
				i := data.(int)
				for j := 0; j < g.numVertexes; j ++ {
					if g.arc[i][j] == 1 && !visited[j] {
						visited[j] = true
						fmt.Println(g.arc[j])
						q.EnQueue(j)
					}
				}
			}
		}
	}
}
