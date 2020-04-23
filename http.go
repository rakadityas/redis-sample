package main

import (
	"fmt"
	"net/http"
)

func initHandler() {
	http.HandleFunc("/ping", handlePing)
	// TODO: Adding more APIs

	fmt.Println("serving in 8181 ")
	http.ListenAndServe(":8181", nil)
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
