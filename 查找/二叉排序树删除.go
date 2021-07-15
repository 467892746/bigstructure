package main

func DeleteBst(t *BiTNode, key int)  bool {
	if t == nil {
		return false
	}else {
		if key == t.data {
			return Delete(&t)
		}else if key < t.data {
			return DeleteBst(t.lChid, key)
		}else  {
			return DeleteBst(t.rChid, key)
		}
	}


}

func Delete(p **BiTNode) bool {
	q := *p
	if q.rChid == nil {
		p = &q.lChid
	}else if q.lChid == nil {
		p = &q.rChid
	}else {
		s := q.lChid
		for s.rChid != nil {
			q = s
			s = s.rChid
		}
		(*p).data = s.data
		if q != *p {
			q.rChid = s.lChid
		}else {
			q.lChid = s.lChid
		}
	}
	return true
}


