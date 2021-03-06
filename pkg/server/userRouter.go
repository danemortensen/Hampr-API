package server

import (
    "log"
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
    subrouter.Put("/delete", ur.deleteUserHandler)
    return subrouter
}

func (ur *userRouter) loginHandler(w http.ResponseWriter, r *http.Request) {
    respond(w, http.StatusOK, nil)
}

func (ur *userRouter) findUserHandler(w http.ResponseWriter,
        r *http.Request) {
    var user bson.M
    authId := r.Header.Get("authId")
    err := ur.userService.FindUser(authId, &user)

    if err != nil {
        log.Printf("User %s not in DB\n", authId)
        err = ur.userService.InsertUser(authId)
        log.Println("check1")
        if err != nil {
            respondWithError(w, http.StatusInternalServerError, "Loading error")
            log.Printf("Unable to insert user into db")
            return
        }
        log.Println("check2")
        err = ur.userService.FindUser(authId, &user)
        if err != nil {
            respondWithError(w, http.StatusInternalServerError, "Loading error")
            log.Printf("Unable to find user after inserting")
            return
        }
    }

    respond(w, http.StatusOK, user)
}

func (ur *userRouter) deleteUserHandler(w http.ResponseWriter,
        r *http.Request) {
    userId := r.Header.Get("authId")
    log.Println(userId)
    err := ur.userService.DeleteUser(userId)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError,
            "Error deleting user")
        log.Printf("Unable to delete user %s from database\n", userId)
        return
    }
    respond(w, http.StatusOK, nil)
}
