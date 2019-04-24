package server

import (
    "log"
    "net/http"

    "github.com/go-chi/chi"

    "github.com/danemortensen/Hampr-API/pkg"
    "github.com/danemortensen/Hampr-API/pkg/config"
)

type Server struct {
    config *config.ServerConfig
    garmentService root.GarmentService
    router *chi.Mux
}

func NewServer(gs root.GarmentService, config *config.ServerConfig) *Server {
    r := chi.NewRouter()

    s := Server {
        config: config,
        garmentService: gs,
        router: r,
    }

    return &s
}

func (s *Server) Start() {
    port := s.config.Port
    log.Printf("Listening on port %s\n", port)
    log.Printf("%s\n", s.config.Auth.Secret)
    s.router.Get("/auth/code", s.handleAuthCode)
    // s.router.Route("/garment", )
    // s.router.Mount("/garment", s.newGarmentRouter())
    newGarmentRouter(s.garmentService, s.router)
    log.Fatal(http.ListenAndServe(port, s.router))
}
