package server

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

type HTTPHandler struct {
	res http.ResponseWriter
	req *http.Request
}

func (h HTTPHandler) Send(str string){
	io.WriteString(h.res, str)
}

func (h HTTPHandler) SendJSON(j map[string]interface{}){
	if jsonString, err := json.Marshal(j); err == nil {
		h.res.Write(jsonString)
	}
}

type RouteHandler func(h HTTPHandler, next func())

var traverseHandlers func()

func Route(uri string, handlers ...RouteHandler){

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

func Start(){
	http.ListenAndServe(":8000", nil)
}

