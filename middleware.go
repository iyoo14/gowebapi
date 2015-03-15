package main

import (
    "net/http"
    "github.com/zenazn/goji/web"
    "log"
)

func ApplicationJson(h http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Content-Type", "application/json")
            h.ServeHTTP(w, r)
    }
    return http.HandlerFunc(fn)
}

func SuperSecure(c *web.C, h http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        auth := r.FormValue("login")
        if auth != "super" {
            log.Printf("no password")
            badRequest(w)
            return
        }
        log.Printf("yes password")
        h.ServeHTTP(w, r)
    }
    return http.HandlerFunc(fn)
}

func badRequest(w http.ResponseWriter) {
    w.WriteHeader(404)
    w.Write([]byte("BAD PASSWORD!\n"))
}
