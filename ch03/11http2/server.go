// https://localhost:8080
// curl --http2 --insecure https://localhost:8080
// 下記ページに --https オプションに対応する方法（へのリンク）が記載されています
// http://localhost/www.marlin-arms.com/support/goweb/links.html

package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	http2.ConfigureServer(&server, &http2.Server{})
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
