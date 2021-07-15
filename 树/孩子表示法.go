package main



type CTNode struct { // 孩子结点
	child int
	next *CTNode
}

type CTBox struct { // 表头结构
	data interface{}
	firstChild *CTNode
	parent int // 双亲位置 根结点为-1
}

type CTree struct { // 数结构
	CTBox [MAX_TREE_SIZE]CTBox
	r,n int // 根的位置和结点数
}

func main() {

}
