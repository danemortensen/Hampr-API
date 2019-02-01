package server

import (
   "log"
   "net/http"

   "github.com/go-chi/chi"
)

type Server struct {
   router *chi.Mux
}

func NewServer() *Server {
   return &Server {
      router: chi.NewRouter(),
   }
}

func (s *Server) Start() {
   log.Println("Listening on port 8080")
   log.Println("TARK!")
   log.Fatal(http.ListenAndServe(":8080", s.router))
}
