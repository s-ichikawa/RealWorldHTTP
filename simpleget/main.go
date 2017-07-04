package main

import (
    "net/http"
    "io/ioutil"
    "log"
    "net/url"
)

func main() {
    values := url.Values{
        "query": {"hello world"},
    }
    resp, err := http.Get("http://localhost:18888?" + values.Encode())
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    log.Println("Status: ", resp.Status)

    log.Println("StatusCode: ", resp.StatusCode)

    log.Println(string(body))

    log.Println("Headers: ", resp.Header)
}
