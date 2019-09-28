package main

import (
      "fmt"
      "net/http"
)

func main() {
      resp, err := http.Get("https://github.com")
      if err != nil {
              panic(err)
      }
      defer resp.Body.Close()
      for k, v := range resp.Header {
              fmt.Print(k)
              fmt.Print(" : ")
              fmt.Println(v)
      }
}

