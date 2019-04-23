package server

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

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
    for k, v := range m {
    switch vv := v.(type) {
        case string:
            fmt.Println(k, "is string", vv)
        case float64:
            fmt.Println(k, "is float64", vv)
        case []interface{}:
            fmt.Println(k, "is an array:")
            for i, u := range vv {
                fmt.Println(i, u)
            }
        default:
            fmt.Println(k, "is of a type I don't know how to handle")
        }
    }

    println(resp)
}
