package main

func InsertBst(t **BiTNode, key int) bool {
	if p, ok := SearchBST(*t, nil,  key); !ok {
		var s = new(BiTNode)
		s.data = key
		if p == nil {
			*t = s
		}else if key < p.data {
			p.lChid = s
		}else {
			p.rChid = s
		}
		return true
	}
	return false
}


