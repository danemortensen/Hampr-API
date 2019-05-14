package server

import (
    "encoding/json"
    "net/http"

    "github.com/danemortensen/Hampr-API/pkg"

    "github.com/go-chi/chi"
    "gopkg.in/mgo.v2/bson"
    "log"
)

/**

Use the following names for all service operations:
Insert
Find
Update
Delete

**/

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
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request body")
        log.Println("Invalid request body for insertGarmentHandler")
        return
    }
    err = gr.garmentService.InsertGarment(authId, &garment)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Upload error")
        log.Println("Unable to insert garment into db")
        return
    }
    respond(w, http.StatusOK, nil)
}

func (gr *garmentRouter) deleteGarmentHandler(w http.ResponseWriter, r *http.Request) {
    var garmentId string
    authId := r.Header.Get("authId")
    err := json.NewDecoder(r.Body).Decode(&garmentId)
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request body")
        log.Println("Invalid request body for deleteGarmentHandler")
        return
    }
    err = gr.garmentService.DeleteGarment(authId, garmentId)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Upload error")
        log.Println("Unable to delete garment from db")
        return
    }
    respond(w, http.StatusOK, nil)
}
