package server

import (

    "github.com/go-chi/chi"
    "github.com/danemortensen/Hampr-API/pkg/db"

    //"log"
)

type Outfit struct {
    Id string
    Name string
    Photo string
    Garments []string
    Dates []string
}

type outfitRouter struct {
    session *db.Session
    router *chi.Mux
}

func newOutfitRouter(session *db.Session) *outfitRouter {
    r := chi.NewRouter()
    o := outfitRouter {
        session: session,
        router: r,
    }
    // r.Post("/new", addOutfitHandler)
    // r.Post("/user", o.getUser)

    return &o
}
