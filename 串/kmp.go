package main

import "fmt"

func genNext(s string) []int {
	sLen := len(s)
	next := make([]int, sLen)
	c := 0
	d := -1
	next[0] = -1
	for c < sLen - 1 {
		if d == -1 || s[c] == s[d] {
			c++
			d++
			if s[c] == s[d] {
				next[c] = next[d]
			}else {
				next[c] = d
			}
		}else {
			d = next[d]
		}
	}
	return next
}

func Kmp(chang string, duan string) int {
	cLen := len(chang)
	dLen := len(duan)
	if dLen == 0 || cLen == 0 {
		return 0
	}
	next := genNext(duan)
	c := -1
	d := -1
	for c < cLen  && d < dLen  {
		if d == -1 || (chang[c] == duan[d])  {
			c++
			d++
		}else {
				d = next[d]
		}
	}
	if d == dLen {
		return c - dLen
	}
	return 0
}

func main() {
	fmt.Println(Kmp("sadsadsadbcv11", "bcv"))
}
