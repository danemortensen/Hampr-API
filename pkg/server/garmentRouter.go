package server

import (
    "encoding/json"
    "net/http"

    "github.com/danemortensen/Hampr-API/pkg"

    "github.com/go-chi/chi"
    "gopkg.in/mgo.v2/bson"
    "log"
)

type garmentRouter struct {
    garmentService root.GarmentService
}

func (s *Server) newGarmentRouter() *chi.Mux {
    subrouter := chi.NewRouter()
    garmentRouter := &garmentRouter {
        garmentService: s.garmentService,
    }
    subrouter.Post("/insert", garmentRouter.insertGarmentHandler)
    return subrouter
}

func (gr *garmentRouter) insertGarmentHandler(w http.ResponseWriter, r *http.Request) {
    var garment bson.M
    authId := r.Header.Get("authId")
    err := json.NewDecoder(r.Body).Decode(&garment)
    log.Println("handler hit")
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request body")
        log.Println("Invalid request body for insertGarmentHandler")
        return
    }
    err = gr.garmentService.InsertGarment(authId, &garment)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Uploading error")
        log.Println("Unable to insert garment into db")
        return
    }
    respond(w, http.StatusOK, nil)
}
