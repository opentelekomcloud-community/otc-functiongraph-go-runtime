// main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful"
)

func registerServer() {
	fmt.Println("Running a Go Http server at localhost:8000/")

	ws := new(restful.WebService)
	ws.Path("/")

	ws.Route(ws.GET("/hello").To(Hello))
	c := restful.DefaultContainer
	c.Add(ws)
	fmt.Println(http.ListenAndServe(":8000", c))
}

func Hello(req *restful.Request, resp *restful.Response) {
	resp.Write([]byte("nice to meet you"))
}

func main() {
	registerServer()
}
