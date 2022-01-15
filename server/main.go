package main

import (
    "encoding/json"
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
)


func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
        res, _ := json.Marshal(map[string]string{"message": "hello"})
        w.WriteHeader(http.StatusOK)
        _, _ = w.Write(res)
    })
    svr := &http.Server{Addr: ":8080", Handler: r}
    //svr.SetKeepAlivesEnabled(false)
    _ = svr.ListenAndServe()
}
