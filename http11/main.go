package main

import (
    "net/http"
    "io/ioutil"
    "log"
)

func main() {
    resp, err := http.Get("https://localhost:18888")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    log.Println(string(body))
}
