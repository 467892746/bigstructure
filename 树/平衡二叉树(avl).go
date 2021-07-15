package main

import "fmt"

type biTNode struct {
	data  int
	bf    int
	left  *biTNode
	right *biTNode
}

func rRotate(p *biTNode) *biTNode {
	l := p.left
	p.left = l.right
	l.right = p
	return l
}

func lRotate(p *biTNode) *biTNode {
	r := p.right
	p.right = r.left
	r.left = p
	return r
}

const (
	LH = 1  // 左高
	EH = 0  // 等高
	RH = -1 // 右高
)

func LeftBalance(t *biTNode) *biTNode {
	l := t.left
	switch l.bf {
	case LH:
		t.bf = EH
		l.bf = EH
		return rRotate(t)
	case RH:
		lr := l.right
		switch lr.bf {
		case LH:
			t.bf = RH
			l.bf = EH
		case EH:
			t.bf = EH
			l.bf = EH
		case RH:
			t.bf = EH
			l.bf = LH
		}
		lr.bf = EH
		t.left = lRotate(t.left)
		return rRotate(t)
	default:
		return nil
	}
}

func RightBalance(t *biTNode) *biTNode {
	r := t.right
	switch r.bf {
	case LH:
		rl := r.left
		switch rl.bf {
		case LH:
			t.bf = EH
			r.bf = RH
		case EH:
			t.bf = EH
			r.bf = EH
		case RH:
			r.bf = EH
			t.bf = LH
		}
		rl.bf = EH
		t.right = rRotate(t.right)
		return lRotate(t)
	case RH:
		t.bf = EH
		r.bf = EH
		return lRotate(t)
	default:
		return nil
	}
}

func InsertAvl(t *biTNode, e int) (*biTNode, bool) {
	var isInsert bool
	if t == nil {
		t = new(biTNode)
		t.data = e
		//fmt.Println(e)
		t.bf = EH
		return t, true
	}
	if t.data == e {
		return t, false
	} else if e < t.data {
		if t.left, isInsert = InsertAvl(t.left, e); !isInsert { // 未插入
			return t, false
		} else { // 已插入到T的左子树中且左子树"长高"
			switch t.bf {
			case LH: // 原本左子树比右子树高，需要做左平衡处理
				t = LeftBalance(t)
				return t, false
			case EH: // 原本左右子树等高，现因左子树增高而树增高
				t.bf = LH
				return t, true
			case RH: // 原本右子树比左子树高，现左右子树登高
				t.bf = EH
				return t, false
			}
		}
	}else {
		if t.right, isInsert = InsertAvl(t.right, e); !isInsert {
			return t, false
		}else {
			switch t.bf {
			case LH:
				t.bf = EH
				return t, false
			case EH:
				t.bf = RH
				return t, true
			case RH:
				t = RightBalance(t)
				return t, false
			}
		}
	}
	return t, true
}
func avlPreOrderTraverse(t *biTNode)  {
	if t == nil {
		return
	}
	fmt.Println(t.data)
	avlPreOrderTraverse(t.left)
	avlPreOrderTraverse(t.right)
}

func search(t *biTNode, data int) bool {
	if t == nil {
		return false
	}
	if t.data == data {
		return true
	}else if data > t.data {
		return search(t.right, data)
	}else {
		return search(t.left, data)
	}
}

func main() {
	a := []int{3,2,1,4,5,6,7,10,9,8}
	var t *biTNode
	for _, b := range a{
		t, _ = InsertAvl(t, b)
	}
	fmt.Println(search(t, 1))
}
