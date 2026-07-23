// main.go
package main

import (
	"fmt"
	"net/http"

	"io"

	"github.com/emicklei/go-restful"
)

// PostHello handles HTTP POST requests
func PostHello(req *restful.Request, resp *restful.Response) {
	var body []byte
	body, err := io.ReadAll(req.Request.Body)
	if err != nil {
		fmt.Println("Read body failed")
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	var requestId string = req.HeaderParameter("X-CFF-Request-Id")

	req.ReadEntity(&body)
	fmt.Printf("Received request [%s] body: %s\n", requestId, string(body))

	resp.Write([]byte("Hello from Go HTTP Function! Your request ID is: " + requestId + "\n"))
}

// GetHello handles HTTP GET requests
func GetHello(req *restful.Request, resp *restful.Response) {
	var requestId string = req.HeaderParameter("X-CFF-Request-Id")
	fmt.Printf("GET request [%s]\n", requestId)

	resp.Write([]byte("Hello from Go HTTP Function GET! Your request ID is: " + requestId + "\n"))
}

// registerServer sets up the HTTP server and routes
func registerServer() {
	fmt.Println("Running a Go Http server at localhost:8000/")

	ws := new(restful.WebService)
	ws.Path("/")

	ws.Route(ws.POST("/hello").To(PostHello))

	ws.Route(ws.GET("/hello").To(GetHello))

	c := restful.DefaultContainer
	c.Add(ws)
	fmt.Println(http.ListenAndServe(":8000", c))
}

// main function
func main() {
	registerServer()
}
