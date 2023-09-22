package main

import (
    "fmt"
    "net/http"

    "github.com/go-chi/chi"
    "go-coin/internal/handlers"
    log "github.com/sirupsen/logrus"
)

func main() {
    log.SetReportCaller(true)
    var r *chi.Mux = chi.NewRouter()
    handler.Handler(r)

    fmt.Println("starting GO API service...")
    err := http.ListenAndServe("localhost:8000", r)
    if err != nil {
        log.Error(err)
    }
}
