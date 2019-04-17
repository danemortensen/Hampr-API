package server

import (
    "encoding/json"
    "net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
    response, _ := json.Marshal(map[string]string{"error": msg})

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}
