package main

import (
    "net/http"
    "fmt"
    "log"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("URL: %s\n", r.URL.String())
    fmt.Printf("Query: %v\n", r.URL.Query())
}

func main() {
    var httpServer http.Server
    http.HandleFunc("/", handler)
    log.Println("start http listening 18888")
    httpServer.Addr = ":18888"

    log.Println(httpServer.ListenAndServe())
}
