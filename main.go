package main

import (
  "fmt"
  "http"
  "io/ioutil"
)

type result struct {
  id int
  bytes int64
  info string
  content []byte
}

func main() {
  c := make(chan *result, 100)

  req := []string{
    "http://www.amazon.com/gp/offer-listing/B003U8HTMG",
  }

  for i, _ := range req {
    go geturl(i, req[i], c)
    fmt.Printf("[%d] %s geturl\n", i, req[i])
  }

  for i, _ := range req {
    res := <-c
    fmt.Printf("[%d] %s (%d bytes): %s\n", i, req[res.id], res.bytes, res.content)
  }

}

func geturl(num int, url string, c chan *result) {
  response, err := http.Get(url)
  defer response.Body.Close()

  if err != nil {
    c <- &result{num, 0, err.String(), []byte{}}
    return
  }

  content, err := ioutil.ReadAll(response.Body)
  if err != nil {
    c <- &result{num, 0, err.String(), []byte{}}
  }

  c <- &result{num, response.ContentLength, response.Status, content}

}
