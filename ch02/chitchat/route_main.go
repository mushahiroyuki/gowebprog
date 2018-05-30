package main

import (
 // 	"github.com/sausheong/gwp/Chapter_2_Go_ChitChat/chitchat/data"
	"github.com/tseno/gowebprog/ch02/chitchat/data"
	"net/http"
)

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	// データの取得
	threads, err := data.Threads()
	if err != nil {
		error_message(writer, request, "Cannot get threads")
	} else {
		// ログインしているか判定する
		_, err := session(writer, request)
		if err != nil {
			// 一般用のナビバー
			generateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			// 会員用のナビバー
			generateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}
