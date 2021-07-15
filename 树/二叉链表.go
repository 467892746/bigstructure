package main

import "github.com/labstack/echo/v4/middleware"

type BiTNode struct {
	data interface{}
	lChild *BiTNode
	rChild *BiTNode
}

func main() {
	middleware.Gzip()
}
