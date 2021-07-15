package main

import "fmt"

type BiThrNode struct {
	data interface{}
	lChild *BiThrNode
	rChild *BiThrNode
	LTag bool // false 表示左孩子 true表示前驱
	RTag bool // false 表示右孩子 true表示后驱
}



func inOrderCreateBiTree(data []int, i *int) *BiThrNode { // 中序遍历生成二叉树
	BiTree :=  &BiThrNode{
		data:   nil,
		lChild: nil,
		rChild: nil,
		LTag:   false,
		RTag:   false,
	}
	*i++
	if *i >= len(data)   {
		return nil
	}
	if data[*i] != 0  {
		t := *i
		BiTree.lChild = inOrderCreateBiTree(data, i)
		BiTree.data = data[t]
		BiTree.rChild = inOrderCreateBiTree(data, i)
	}else {
		return nil
	}
	return BiTree
}

func InOrderTraverseBiTree(t *BiThrNode)  { // 中序遍历二叉树
	if t == nil  {
		return
	}
	if !t.LTag {
		InOrderTraverseBiTree(t.lChild)
	}
	fmt.Println(t.data)
	if !t.RTag {
		InOrderTraverseBiTree(t.rChild)
	}
}

var Pre *BiThrNode
func InThreading(biThrTree *BiThrNode)  { // 中序遍历生成线索二叉树
	if biThrTree == nil {
		return
	}
	InThreading(biThrTree.lChild)
	if biThrTree.lChild == nil  {
		biThrTree.LTag = true
		biThrTree.lChild = Pre
	}
	if Pre != nil && Pre.rChild == nil {
		Pre.RTag = true
		Pre.rChild = biThrTree
	}
	Pre = biThrTree
	InThreading(biThrTree.rChild)
}

func InOrderTraverseThr(rootBiTree *BiThrNode)  { // 中序遍历线索二叉树
	for rootBiTree != nil {
		for !rootBiTree.LTag && rootBiTree.lChild != nil {
			rootBiTree = rootBiTree.lChild
		}
		fmt.Println(rootBiTree.data)
		for rootBiTree.RTag && rootBiTree.rChild != nil  {
			rootBiTree = rootBiTree.rChild
			fmt.Println(rootBiTree.data)
		}
		rootBiTree = rootBiTree.rChild
	}
}

func main() {
	data := []int{0,2,0,4,0,1,0,3,0}
	var i = 0
	var biTreeHeadNode = &BiThrNode{}
	biThr := inOrderCreateBiTree(data, &i)
	biTreeHeadNode.lChild = biThr

	InThreading(biThr)
	InOrderTraverseThr(biThr)
}
