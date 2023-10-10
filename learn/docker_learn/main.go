package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	server := &http.Server{
		Addr: ":8888",
	}
	fmt.Println("server startup...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("server startup failed...", err)
	}
}

func hello(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("hello sfy docker learn"))
}
