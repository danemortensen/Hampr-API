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
    userService root.UserService
    router *chi.Mux
}

func NewServer(gs root.GarmentService, us root.UserService, config *config.ServerConfig) *Server {
    r := chi.NewRouter()
    s := Server {
        config: config,
        garmentService: gs,
        userService: us,
        router: r,
    }
    return &s
}

func (s *Server) Start() {
    port := s.config.Port
    log.Printf("Listening on port %s\n", port)
    log.Printf("%s\n", s.config.Auth.AppSecret)

    s.router.Mount("/auth", s.newAuthRouter())
    s.router.Mount("/api", s.newApiRouter())

    log.Fatal(http.ListenAndServe(port, s.router))
}

func (s *Server) newApiRouter() *chi.Mux {
    apiRouter := chi.NewRouter()
    apiRouter.Use(s.authMiddleware)
    apiRouter.Mount("/user", s.newUserRouter())
    apiRouter.Mount("/garments", s.newGarmentRouter())
    return apiRouter
}
