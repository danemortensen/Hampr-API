package server

import (
    "encoding/json"
    "errors"
    "net/http"

    "github.com/go-chi/chi"
)

type Garment struct {
    Name string
}

func newGarmentRouter() http.Handler {
    r := chi.NewRouter()
    r.Post("/new", addGarmentHandler)
    return r
}

func addGarmentHandler(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()

    // g, err := decodeGarment(r)
    // if err != nil {
    //     respondWithError(w, http.StatusBadRequest, "Invalid request payload")
    //     return
    // }


}

func decodeGarment(r *http.Request) (Garment, error) {
    var g Garment

    if r.Body == nil {
        return g, errors.New("Missing request body")
    }

    err := json.NewDecoder(r.Body).Decode(&g)

    return g, err
}
