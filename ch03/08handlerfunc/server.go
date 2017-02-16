package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
		// Handlerは指定しない -> DefaultServeMuxをハンドラとして利用
	}
	http.HandleFunc("/hello", hello) // 関数をハンドラに変換して、DefaultServeMuxに登録
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}
