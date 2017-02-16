package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	hello := HelloHandler{} // helloはハンドラ（http.Handler）。ServeHTTPを持っているので
	world := WorldHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
		// Handlerは指定しない -> DefaultServeMuxをハンドラとして利用
	}

	http.Handle("/hello", &hello) // ハンドラhelloをDefaultServeMuxに登録
	http.Handle("/world", &world)

	server.ListenAndServe()
}
