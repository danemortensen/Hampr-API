package server

import (
    "encoding/json"
    "net/http"

    "github.com/danemortensen/Hampr-API/pkg"

    "github.com/go-chi/chi"
)

type garmentRouter struct {
    garmentService root.GarmentService
}

func (s *Server) newGarmentRouter() *chi.Mux {
    subrouter := chi.NewRouter()
    garmentRouter := &garmentRouter {
        garmentService: s.garmentService,
    }
    subrouter.Post("/create", garmentRouter.createGarmentHandler)
    return subrouter
}

func (gr *garmentRouter) createGarmentHandler(w http.ResponseWriter,
        r *http.Request) {
    var garment root.Garment
    err := json.NewDecoder(r.Body).Decode(&garment)
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    err = gr.garmentService.InsertGarment(&garment)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respond(w, http.StatusOK, nil)
}
