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

    s := Server {
        config: config,
        router: r,
    }
    s.registerHandlers()
    return &s
}

func (s *Server) Start() {
    port := s.config.Port
    log.Printf("Listening on port %s\n", port)
    log.Fatal(http.ListenAndServe(port, s.router))
}

func (s *Server) registerHandlers() {
    s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("hello"))
    })
}
