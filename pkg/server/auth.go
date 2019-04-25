package server

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "github.com/go-chi/chi"
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "log"
)

func (s *Server) newAuthRouter() *chi.Mux {

    authRouter := chi.NewRouter()
    authRouter.Get("/code", s.authCodeHandler)
    return authRouter
}

func (s *Server) authMiddleware(next http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        authId := r.Header.Get("authId")
        token := r.Header.Get("token")
        secret := s.config.Auth.AppSecret

        if authId == "" || token == "" {
            respondWithError(w, http.StatusInternalServerError, "Authorization error")
            return
        }

        h := hmac.New(sha256.New, []byte(secret))
        h.Write([]byte(token))
        proof := hex.EncodeToString(h.Sum(nil))

        fmt.Println("authId:", authId, "proof:", proof, "token:", token)

        exchangeUrl := fmt.Sprintf("https://graph.accountkit.com/v1.3/me?access_token=%s&appsecret_proof=%s", token, proof)

        resp, err := http.Get(exchangeUrl)
        if err != nil {
            respondWithError(w, http.StatusInternalServerError, "Authorization error")
            log.Println(err)
            return
        }
        defer resp.Body.Close()

        b, err := ioutil.ReadAll(resp.Body)
        var f interface{}
        err = json.Unmarshal(b, &f)
        m := f.(map[string]interface{})
        expected := m["id"]

        if authId == expected {
            next.ServeHTTP(w, r)
            log.Println("login succeeded")
        } else {
            respondWithError(w, http.StatusInternalServerError, "Authorization error")
            log.Println("login failed")
        }
    }
    return http.HandlerFunc(fn)
}

func (s *Server) authCodeHandler(w http.ResponseWriter, r *http.Request) {
    code := r.Header.Get("Auth-Code")
    //fmt.Printf("code: %s\n", code)

    exchangeUrl := fmt.Sprintf("https://graph.accountkit.com/v1.3/access_token?grant_type=authorization_code&code=%s&access_token=AA|%s|%s", code, s.config.Auth.AppId, s.config.Auth.AppSecret)
    resp, err := http.Get(exchangeUrl)
    if err != nil {
        fmt.Println("There was an error")
    }
    defer resp.Body.Close()

    b, err := ioutil.ReadAll(resp.Body)
    var f interface{}
    err = json.Unmarshal(b, &f)
    m := f.(map[string]interface{})
    //fmt.Println(m)

    if m["error"] != nil {
        respondWithError(w, http.StatusInternalServerError, "Authorization error")
        return
    }

    authId := m["id"]
    token := m["access_token"]
    if authId == nil || token == nil {
        respondWithError(w, http.StatusInternalServerError, "Authorization error")
        return
    }

    respond(w, http.StatusOK, map[string]interface{}{"id": authId, "access_token": token})
}
