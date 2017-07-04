package echo

import (
    "net/http"
    "net/http/httputil"
    "fmt"
    "log"
)

func handler(w http.ResponseWriter, r *http.Request) {
    dump, err := httputil.DumpRequest(r, true)
    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }
    fmt.Println(dump)

    w.Header().Add("Set-Cookie", "VISIT=TRUE")
    if _, ok := r.Header["Cookie"]; ok {
        fmt.Fprintf(w, "<http><body>2回目の訪問</body></http>¥n")
    } else {
        fmt.Fprintf(w, "<http><body>初訪問</body></http>¥n")
    }
}

func main() {
    var httpServer http.Server
    http.HandleFunc("/", handler)
    log.Println("start http listening 18888")
    httpServer.Addr = ":18888"

    log.Println(httpServer.ListenAndServe())
}
