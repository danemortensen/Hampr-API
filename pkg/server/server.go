package server

import (
    "log"
    "net/http"

    "github.com/go-chi/chi"

    "github.com/danemortensen/Hampr-API/pkg/config"
)

type Server struct {
    config *config.ServerConfig
    router *chi.Mux
}

func NewServer(config *config.ServerConfig) *Server {
    r := chi.NewRouter()
    r.Mount("/garment", newGarmentRouter())

    return &Server {
        config: config,
        router: r,
    }
}

func (s *Server) Start() {
    port := s.config.Port
    log.Printf("Listening on port %s\n", port)
    log.Fatal(http.ListenAndServe(port, s.router))
}
