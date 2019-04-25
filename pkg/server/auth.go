package server

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "github.com/go-chi/chi"
)

func (s *Server) newAuthRouter() *chi.Mux {

    authRouter := chi.NewRouter()
    authRouter.Get("/code", s.authCodeHandler)
    //secret := s.config.Auth.AppSecret
    return authRouter
}

func authMiddleware(next http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        //accessToken :=
        // do auth
        // if valid user:
            next.ServeHTTP(w, r)
        // else
            // http.Error(w, "Unauthenticated", 500);
    }
    return http.HandlerFunc(fn)
}

func (s *Server) authCodeHandler(w http.ResponseWriter, r *http.Request) {
    code := r.Header.Get("Auth-Code")
    fmt.Printf("code: %s\n", code)

    exchangeUrl := fmt.Sprintf("https://graph.accountkit.com/v1.3/access_token?grant_type=authorization_code&code=%s&access_token=AA|%s|%s", code, s.config.Auth.AppId, s.config.Auth.AppSecret)
    resp, err := http.Get(exchangeUrl)
    if err != nil {
        fmt.Println("There was an error")
    }
    defer resp.Body.Close()

    // var body map[string]interface{}
    // err = json.NewDecoder(resp.Body).Decode(&body)

    // error check

    b, err := ioutil.ReadAll(resp.Body)
    var f interface{}
    err = json.Unmarshal(b, &f)
    m := f.(map[string]interface{})
    // for k, v := range m {
    // switch vv := v.(type) {
    //     case string:
    //         fmt.Println(k, "is string", vv)
    //     case float64:
    //         fmt.Println(k, "is float64", vv)
    //     case []interface{}:
    //         fmt.Println(k, "is an array:")
    //         for i, u := range vv {
    //             fmt.Println(i, u)
    //         }
    //     default:
    //         fmt.Println(k, "is of a type I don't know how to handle")
    //     }
    // }
    fmt.Println(m)
    // fmt.Println(reflect.TypeOf(m["error"]))
    if m["error"] != nil {
        respondWithError(w, http.StatusInternalServerError, "Authorization error")
        return
    }

    userId := m["id"]
    accessToken := m["access_token"]
    if userId == nil || accessToken == nil {
        respondWithError(w, http.StatusInternalServerError, "Authorization error")
        return
    }

    respond(w, http.StatusOK, map[string]interface{}{"id": userId, "access_token": accessToken})
    // response := map[string]interface{}{""}
}
