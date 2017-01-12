package main              

import (
  "testing"
)

func TestDecode(t *testing.T) {
  post, err := decode("post.json")

  if err != nil {                      
    t.Error(err)
  }
  if post.Id != 1 {                                    
    t.Error("Wrong id, was expecting 1 but got", post.Id)
  }
  if post.Content != "Hello World!" {                         
    t.Error("Wrong content, was expecting 'Hello World!' but got", post.Content)
  }
}

func TestEncode(t *testing.T) {
  t.Skip("Skipping encoding for now")                  
}
