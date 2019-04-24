package server

import (
    "encoding/json"
    "net/http"
    "github.com/go-chi/chi"
    "github.com/danemortensen/Hampr-API/pkg/db"
    "gopkg.in/mgo.v2/bson"
    //"log"
)

type Outfit struct {
    Id string
    Name string
    Photo string
    Garments []string
    Dates []string
}

type outfitRouter struct {
    session *db.Session
    router *chi.Mux
}

func newOutfitRouter(session *db.Session) *outfitRouter {
    r := chi.NewRouter()
    o := outfitRouter {
        session: session,
        router: r,
    }
    r.Post("/new", addOutfitHandler)
    r.Post("/user", o.getUser)

    return &o
}

func (o *outfitRouter) getUser(w http.ResponseWriter, r *http.Request) {
    var body bson.M
    var result bson.M
    err := json.NewDecoder(r.Body).Decode(&body)
    if err != nil {
        http.Error(w, "Invalid Request Body", 400);
        return
    }
    o.session.Find("users", bson.M{"_id": body["userId"]}, &result)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
    // bin, err := bson.Marshal(result)
    // if err != nil {
    //     log.Print(err)
    // }
    // w.Write(bin)
}

func addOutfitHandler(w http.ResponseWriter, r *http.Request) {
    // defer r.Body.Close()
    // o, err := decodeOutfit(r)

    // g, err := decodeGarment(r)
    // if err != nil {
    //     respondWithError(w, http.StatusBadRequest, "Invalid request payload")
    //     return
    // }


}

func helloWorld2(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello"))
}

// func decodeOutfit(r *http.Request) (Outfit, error) {
//     var o Outfit
//
//     if r.Body == nil {
//         return o, errors.New("Missing request body")
//     }
//
//     err := json.NewDecoder(r.Body).Decode(&o)
//
//     return o, err
// }
