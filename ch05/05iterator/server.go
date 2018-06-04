// リスト5.5
package main

import (
	"html/template"
	"net/http"
	"os"
	"fmt"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("ch05/05iterator/tmpl.html")
	// daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	daysOfWeek := []string{"月", "火", "水", "木", "金", "土", "日"}
	t.Execute(w, daysOfWeek)
}

func main() {
	p, _ := os.Getwd()
	fmt.Println(p)
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
