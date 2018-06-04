// リスト5.28
package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("ch05/28nested2/layout.html", "ch05/28nested2/red_hello.html")
	} else {
		t, _ = template.ParseFiles("ch05/28nested2/layout.html", "ch05/28nested2/blue_hello.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
