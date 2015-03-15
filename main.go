package main

import (
    "fmt"
    "net/http"
    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
//    "github.com/zenazn/goji/web/middleware"
//    "github.com/iyoo14/gowebapi/model/member"
)

func list(c web.C, w http.ResponseWriter, r *http.Request) {
    v := r.FormValue("login")
    fmt.Fprintf(w, "Request /foo/list : %s\n", v)
}

func add(c web.C, w http.ResponseWriter, r *http.Request) {
    //v := r.FormValue("login")
    v := r.PostForm
    /*
    m := member.New()
    m.SetVal(v)
    lst := m.List()
    */
    fmt.Fprintf(w, "Request /foo/add : %s\n", v)
}

func admin(c web.C, w http.ResponseWriter, r *http.Request) {
    v := r.FormValue("login")
    fmt.Fprintf(w, "Request /admin/* : %s\n", v)
}

type MyHandler struct{}
type JsonHandler struct{}

func (h *JsonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    var js string
    switch r.Method {
        case "POST":
        n := "n"
        v := r.FormValue(n)
        js = `{"res":"` + v + `"}`
        fmt.Fprintf(w, js)
    case "GET":
        js = `{"res":"get"}`
        fmt.Fprintf(w, js)
    }
}

func (my *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Request /bar")
}

func main() {
    goji.Get("/foo/list", list)
    goji.Post("/foo/add",  add)
    goji.Use(ApplicationJson)
    goji.Use(SuperSecure)
/*
    adminMux := web.New()
    goji.Handle("/*", adminMux)
    adminMux.Use(middleware.SubRouter)
    adminMux.Use(SuperSecure)
    adminMux.Post("/login", admin)
*/
    goji.Serve()

}

