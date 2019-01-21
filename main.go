package main

import (
   "flag"
   "log"
   "net/http"

   "github.com/danemortensen/Hampr-API/dbao"

   "github.com/go-chi/chi"
)

func addUser(w http.ResponseWriter, r *http.Request) {
   log.Print("Adding user")
}

func main() {
   router := chi.NewRouter()

   router.Put("/Users", addUser)

   log.Print("Server running on port 3000")
   log.Fatal(http.ListenAndServe(":3000", router))
}
