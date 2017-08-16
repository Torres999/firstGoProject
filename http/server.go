package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"firstGoProject/httpService"
)

func main() {
	/**
	=========================
	方式1：启动一个18080端口服务
	=========================
	*/
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		log.Println("receive a request,remote id address is :", r.RemoteAddr)
	})

	http.HandleFunc("/call", createHandler)

	//http.ListenAndServe(":18080", new(selfHandler))//18080服务对应的mapping：localhost:18080/*

	http.ListenAndServe(":18080", nil) //18080服务对应的mapping：localhost:18080/bar、localhost:18080/call

	/**
	=========================
	方式2：启动一个18081端口服务
	=========================
	*/
	server := &http.Server{Addr: ":18081", Handler: new(controller.Controller)}
	server.ListenAndServe()
}

func createHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Init request handler")
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

type selfHandler struct {
}

func (t *selfHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("Init a self handler")
	w.Write([]byte("This is an self server.\n"))
}
