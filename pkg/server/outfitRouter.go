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

type outfitRouter struct {
    outfitService root.OutfitService
}

func (s *Server) newOutfitRouter() *chi.Mux {
    subrouter := chi.NewRouter()
    outfitRouter := &outfitRouter {
        outfitService: s.outfitService,
    }
    subrouter.Post("/insert", outfitRouter.insertOutfitHandler)
    subrouter.Delete("/delete", outfitRouter.deleteOutfitHandler)
    return subrouter
}

func (gr *outfitRouter) insertOutfitHandler(w http.ResponseWriter, r *http.Request) {
    var outfit bson.M
    authId := r.Header.Get("authId")
    err := json.NewDecoder(r.Body).Decode(&outfit)
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request body")
        log.Println("Invalid request body for insertOutfitHandler")
        return
    }
    err = gr.outfitService.InsertOutfit(authId, &outfit)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Upload error")
        log.Println("Unable to insert outfit into db")
        return
    }
    respond(w, http.StatusOK, nil)
}

func (gr *outfitRouter) deleteOutfitHandler(w http.ResponseWriter, r *http.Request) {
    var outfitId string
    authId := r.Header.Get("authId")
    err := json.NewDecoder(r.Body).Decode(&outfitId)
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request body")
        log.Println("Invalid request body for deleteOutfitHandler")
        return
    }
    err = gr.outfitService.DeleteOutfit(authId, outfitId)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Upload error")
        log.Println("Unable to delete outfit from db")
        return
    }
    respond(w, http.StatusOK, nil)
}
