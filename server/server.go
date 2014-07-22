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
	return h
}

type RouteHandler func(h HTTPHandler, next func())

var traverseHandlers func()

func Route(uri string, handlers ...RouteHandler){

	fmt.Println(router.CreateRoute(uri))

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
}

