package main

import (
	"fmt"
	"github.com/seanmcgary/minweb/server"
)

func testMiddleware(h server.HTTPHandler, next func()){
	fmt.Println("LOL some middleware");
	next()
}

func main() {
	s := server.Create()

	s.UseMiddleware(func(h server.HTTPHandler, next func()){
		fmt.Println("middleware for all")
	})

	
	server.Route("/test/:foo", testMiddleware, func(h server.HTTPHandler, next func()){
		fmt.Println("test")
		h.Send("text from /test")
	})

	//server.Start()
	

	//reg, _ := regexp.Compile(`(\/[^/]+)`);
	//fmt.Println(reg.FindAllString("/foobar/:test", -1))

	/*
	reg, _ := regexp.Compile(`\/`)

	url := "/foo/:test/:balls"
	keys := make([]string, 0, 0)
	source := url

	url = reg.ReplaceAllString(url, "\\/")

	reg, _ = regexp.Compile(`\.`)
	url = reg.ReplaceAllString(url, `\\.?`)

	reg, _ = regexp.Compile(`\*`)
	url = reg.ReplaceAllString(url, `.+`)

	reg, _ = regexp.Compile(`:(\w+)(?:\(([^\)]+)\))?(\?)?`)

	url = reg.ReplaceAllStringFunc(url, func(str string) string{
		keys = append(keys, str[1:])
		return `([^\/]+)`
	})

	reg, _ = regexp.Compile(`\\\/\(\[\^\/\]\*\)`)
	url = reg.ReplaceAllString(url, `(?:\\/(\\w*))?`)

	url = "^" + url + `\/?$`

	fmt.Println(url)
	fmt.Println(keys)
	fmt.Println(source)
	


	reg, _ = regexp.Compile(url);

	match := reg.FindAllStringSubmatch("/foo/key/", -1)
	if len(match) > 0 {
		fmt.Println(match)	
	} else {
		fmt.Println("no match")
	}
	*/
	
}
