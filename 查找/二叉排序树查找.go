package main

type BiTNode struct {
	data int
	lChid *BiTNode
	rChid *BiTNode
}

func SearchBST(t *BiTNode, f *BiTNode,  key int) (*BiTNode, bool) {
	if t == nil {
		return f, false
	}else if key == t.data {
		return t, true
	}else if key < t.data {
		return SearchBST(t.lChid, t,  key)
	} else  {
		return SearchBST(t.rChid, t,  key)
	}
}