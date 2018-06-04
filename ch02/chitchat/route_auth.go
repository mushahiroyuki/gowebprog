package main

import (
//	"github.com/sausheong/gwp/Chapter_2_Go_ChitChat/chitchat/data"
	"github.com/tseno/gowebprog/ch02/chitchat/data"
	"net/http"
)

// GET /login
// Show the login page
func login(writer http.ResponseWriter, request *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(writer, nil)
}

// GET /signup
// Show the signup page
func signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

// POST /signup
// Create the user account
func signupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
	}
	user := data.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		danger(err, "Cannot create user")
	}
	http.Redirect(writer, request, "/login", 302)
}

// POST /authenticate
// Authenticate the user given the email and password
func authenticate(writer http.ResponseWriter, request *http.Request) {
	// フォームの内容をパースする。
	err := request.ParseForm()
	// 構造体Userを作成して、入力されたemailをキーにDBからデータを入れる。
	user, err := data.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		danger(err, "Cannot find user")
	}
	// ハッシュ化されたpasswordと比較する
	if user.Password == data.Encrypt(request.PostFormValue("password")) {
		// DBのsessionテーブルにinsertする
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}
		// セッションクッキーの作成。有効期限は設定しないことで、ブラウザが終了する際に自動的に消去される。
		// HttpOnlyで、HTTP,HTTPSに限定される。Javascript等のからはアクセスできない。
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}

}

// GET /logout
// Logs the user out
func logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		warning(err, "Failed to get cookie")
		session := data.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(writer, request, "/", 302)
}
