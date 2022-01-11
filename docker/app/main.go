package main

import (
	"fmt"
	"net/http"
)

type Server struct {
}

func (Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	http.ListenAndServe(":8080", Server{})
	fmt.Println("Demo App Running at port:8080")
}
