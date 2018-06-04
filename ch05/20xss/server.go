// リスト5.20
package main

import (
  "net/http"
  "html/template"
)

func process(w http.ResponseWriter, r *http.Request) {  
//  w.Header().Set("X-XSS-Protection", "0")
  t, _ := template.ParseFiles("ch05/20xss/tmpl.html")
  t.Execute(w, r.FormValue("comment"))
//   t.Execute(w, template.HTML(r.FormValue("comment")))
}

func form(w http.ResponseWriter, r *http.Request) {  
  t, _ := template.ParseFiles("ch05/20xss/form.html")
  t.Execute(w, nil)  
}

func main() {
  server := http.Server{
    Addr: "127.0.0.1:8080",
  }
  http.HandleFunc("/process", process)
  http.HandleFunc("/", form)
  server.ListenAndServe()
}
