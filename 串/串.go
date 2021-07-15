package main

import "fmt"

func Index(s string, t string, pos int) int {
	i := pos
	j := 0
	n := len(s)
	m := len(t)
	for i < n && j < m  {
		if s[i]  == t[j]{
			i++
			j++
		}else {
				i = i-j + 1
				j = 0
		}
	}
	if j >= m {
		return i - m
	}else {
		return 0
	}
}

func main() {
	fmt.Println(Index("213", "13", 0))
}
