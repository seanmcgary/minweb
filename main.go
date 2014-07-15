package main

import (
	"fmt"
	"github.com/seanmcgary/servit/server"
)

func testMiddleware(h server.HTTPHandler, next func()){
	fmt.Println("LOL some middleware");
	next()
}

func main() {
	server.Route("/test", testMiddleware, func(h server.HTTPHandler, next func()){
		fmt.Println("test")
		h.Send("text from /test")
	})

	server.Start()
}
