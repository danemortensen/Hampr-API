package server

import (
    "log"
    "fmt"
    "net/http"
    "github.com/danemortensen/Hampr-API/pkg"
    "github.com/go-chi/chi"
    "gopkg.in/mgo.v2/bson"
)

/**

Use the following names for all service operations:
Insert
Find
Update
Delete

**/

type userRouter struct {
    userService root.UserService
}

func (s *Server) newUserRouter() *chi.Mux {
    subrouter := chi.NewRouter()
    ur := &userRouter {
        userService: s.userService,
    }
    subrouter.Get("/login", ur.loginHandler)
    subrouter.Get("/find", ur.findUserHandler)
    return subrouter
}

func (ur *userRouter) loginHandler(w http.ResponseWriter, r *http.Request) {
    respond(w, http.StatusOK, nil)
}

func (ur *userRouter) findUserHandler(w http.ResponseWriter,
        r *http.Request) {
    ids, ok := r.URL.Query()["id"]
    if !ok || len(ids[0]) < 1 {
        respondWithError(w, http.StatusBadRequest, "Loading error")
        log.Println("Url Param 'id' is missing")
        return
    }

    var user bson.M
    err := ur.userService.FindUser(ids[0], &user)

    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Loading error")
        log.Println("Unable to find user in database")
        return
    }

    respond(w, http.StatusOK, user)
}
