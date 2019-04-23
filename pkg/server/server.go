package server

import (
    "log"
    "net/http"

    "github.com/go-chi/chi"

    "github.com/danemortensen/Hampr-API/pkg/config"
    "github.com/danemortensen/Hampr-API/pkg/db"
)

type Server struct {

    router *chi.Mux
    session *db.Session
    config *config.ServerConfig
}

func NewServer(config *config.ServerConfig, session *db.Session) *Server {
    o := newOutfitRouter(session)
    r := chi.NewRouter()
    r.Mount("/garment", newGarmentRouter())
    r.Mount("/outfit", o.router)
    s := Server {

        router: r,
        session: session,
        config: config,
    }

    return &s
}

func (s *Server) Start() {
    port := s.config.Port
    log.Printf("Listening on port %s\n", port)
    log.Fatal(http.ListenAndServe(port, s.router))
}
