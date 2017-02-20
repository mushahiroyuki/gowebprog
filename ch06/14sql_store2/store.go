// リスト6.14
/*
 1. Postgresを起動
 2. psql -U gwp -f setup.sql -d gwp
   （一番最初は ERRORが表示されるが、そのままでOK)
 3. go run store.go
*/

package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

type Post struct {
	Id       int
	Content  string
	Author   string
	Comments []Comment
}

type Comment struct {
	Id      int
	Content string
	Author  string
	Post    *Post
}

var Db *sql.DB

// connect to the Db
func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func (comment *Comment) Create() (err error) {
	if comment.Post == nil {
 		err = errors.New("投稿が見つかりません")
		// err = errors.New("Post not found")
		return
	}
	err = Db.QueryRow("insert into comments (content, author, post_id) values ($1, $2, $3) returning id", comment.Content, comment.Author, comment.Post.Id).Scan(&comment.Id)
	return
}

// Get a single post
func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Comments = []Comment{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)

	rows, err := Db.Query("select id, content, author from comments where post_id = $1", id)
	if err != nil {
		return
	}
	for rows.Next() {
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
	return
}

// Create a new post
func (post *Post) Create() (err error) {
	err = Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id", post.Content, post.Author).Scan(&post.Id)
	return
}

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	post.Create()

	// Add a comment
	// comment := Comment{Content: "Good post!", Author: "Joe", Post: &post}
 	comment := Comment{Content: "いい投稿だね！", Author: "Joe", Post: &post}
	comment.Create()
	readPost, _ := GetPost(post.Id)

	fmt.Println(readPost)    // {1 Hello World! Sau Sheong [{1 いい投稿だね！ Joe 0xc420018800}]}
	fmt.Println(readPost.Comments)         // [{1 いい投稿だね！ Joe 0xc420018800}]
	fmt.Println(readPost.Comments[0].Post) // &{1 Hello World! Sau Sheong [{1 いい投稿だね！ Joe 0xc420018800}]}
}
