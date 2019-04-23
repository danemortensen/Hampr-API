package server

import (
    "encoding/json"
    "net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
    payload := map[string]string{"error": msg}
    respond(w, code, payload)
}

func respond(w http.ResponseWriter, code int, payload interface{}) {
    // TODO: handle error herer
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}
