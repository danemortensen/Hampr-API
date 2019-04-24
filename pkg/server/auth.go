package server

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    // "reflect"
)

func printJson(m map[string]interface{}) {
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
    //     case map[string]interface{}:
    //         printJson(v)
    //     default:
    //         fmt.Println(k, "is of a type I don't know how to handle")
    //     }
    // }
    fmt.Println(m)
}

func (s *Server) handleAuthCode(w http.ResponseWriter, r *http.Request) {
    code := r.Header.Get("Auth-Code")
    fmt.Printf("code: %s\n", code)

    exchangeUrl := fmt.Sprintf("https://graph.accountkit.com/v1.3/access_token?grant_type=authorization_code&code=%s&access_token=AA|%s|%s", code, s.config.Auth.Id, s.config.Auth.Secret)
    resp, err := http.Get(exchangeUrl)
    if err != nil {
        fmt.Println("There was an error")
    }
    defer resp.Body.Close()

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
    printJson(m)
    // fmt.Println(reflect.TypeOf(m["error"]))
    if m["error"] != nil {
        respondWithError(w, http.StatusInternalServerError, "Authorization error")
        return
    }

    id := m["id"]
    access_token := m["access_token"]
    if id == nil || access_token == nil {
        respondWithError(w, http.StatusInternalServerError, "Authorization error")
        return
    }

    respond(w, http.StatusOK, map[string]interface{}{"id": id, "access_token": access_token})
    // response := map[string]interface{}{""}
}
