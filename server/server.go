package server

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
	"github.com/seanmcgary/minweb/router"
)

type HTTPHandler struct {
	res http.ResponseWriter
	req *http.Request
}

type HTTPServer struct {
	middlewares []func()
	routes map[string]router.Route
}

func (h HTTPServer) Start(){
	http.ListenAndServe(":8000", nil)
}

func (h HTTPServer) UseMiddleware(m func(h HTTPHandler, next func())){
	//fmt.Println(len(h.middlewares))
}

func (h HTTPHandler) Send(str string){
	io.WriteString(h.res, str)
}

func (h HTTPHandler) SendJSON(j map[string]interface{}){
	if jsonString, err := json.Marshal(j); err == nil {
		h.res.Write(jsonString)
	}
}

func Create() (h HTTPServer){
	h = HTTPServer{}

	h.routes = make(map[string]router.Route)

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request){
		fmt.Println(req.URL.Path);
		fmt.Println(h.routes);

		found := false
		for _, value := range h.routes {
			if(value.Match(req.URL.Path)){
				fmt.Println("MATCHES!")
				return
			}
		}
	})

	return h
}

type RouteHandler func(h HTTPHandler, next func())

var traverseHandlers func()

func (h HTTPServer) Route(uri string, handlers ...RouteHandler){

	route := router.CreateRoute(uri)

	h.routes[route.UrlPattern] = route


	/*
	http.HandleFunc(uri, func(res http.ResponseWriter, req *http.Request) {
		current := 0
		handler := HTTPHandler{res, req}

		fmt.Println(handlers)

		traverseHandlers = func() {
			handleThis := handlers[current]
			handleThis(handler, func() {
				current++
				if len(handlers) > current {
					fmt.Println("next handler")
					traverseHandlers()
				}
			})
		}

		if len(handlers) > 0 {
			traverseHandlers()
		}
	})
	*/
}

