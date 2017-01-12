package main

import (
  "encoding/json"
  "fmt"
  "os"
)

type Post struct {
  Id       int       `json:"id"`
  Content  string    `json:"content"`
  Author   Author    `json:"author"`
  Comments []Comment `json:"comments"`
}

type Author struct {
  Id   int    `json:"id"`
  Name string `json:"name"`
}

type Comment struct {
  Id      int    `json:"id"`
  Content string `json:"content"`
  Author  string `json:"author"`
}

func decode(filename string) (post Post, err error) {  
  jsonFile, err := os.Open(filename)
  if err != nil {
    fmt.Println("Error opening JSON file:", err)
    return
  }
  defer jsonFile.Close()

  decoder := json.NewDecoder(jsonFile)
  err = decoder.Decode(&post)
  if err != nil {
    fmt.Println("Error decoding JSON:", err)
    return
  }
  return
}

func main() {
  _, err := decode("post.json")    
  if err != nil {
    fmt.Println("Error:", err)
  }
}
