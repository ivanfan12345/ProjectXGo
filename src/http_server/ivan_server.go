package main

import (
    "fmt"
    "net/http"
)

// define a type for the response
type Hello struct{}

// let that type implement the ServeHTTP method (defined in interface http.Handler)
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    for i := 0 ; i < 100 ; i ++ {
      //time.Sleep(100 * time.Millisecond)
      fmt.Fprint(w, "Hello World! Welcome to the 范志愷 server!")
    }
}

func main() {
    var h Hello
    //http.ListenAndServe("localhost:4000", h)
    http.ListenAndServe("localhost:4000", h)
}

// Here's the method signature of http.ServeHTTP:
// type Handler interface {
//     ServeHTTP(w http.ResponseWriter, r *http.Request)
// }
