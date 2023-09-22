package handlers

import (
    "github.com/go-chi/chi"
    chimiddle "github.com/go-chi/chi/middleware"
    "go-coin/internal/middleware"
)

func Handler(r *chi.Mux) {
    r.User(chimiddle.StripSlashes)

    r.Route("/account", func(router chi.Router) {
        router.Use(middleware.Authorization)

        router.Get("/coins", GetCoinBalance)
    })
}



