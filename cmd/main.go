package main

import (
    "github.com/danemortensen/Hampr-API/pkg/config"
    "github.com/danemortensen/Hampr-API/pkg/db"
    "github.com/danemortensen/Hampr-API/pkg/server"
)

type App struct {
    server *server.Server
    session *db.Session
    config *config.Config
}

func NewApp() *App {
    config := config.NewConfig()
    session := db.NewSession(config.Mongo)
    server := server.NewServer(config.Server, session)
    app := App {
        server: server,
        session: session,
        config: config,
    }
    return &app
}

func (a *App) Run() {
    defer a.session.Close()
    a.server.Start()
}

func main() {
    // config := config.NewConfig()
    // session := db.NewSession(config.Mongo)
    // defer session.Close()
    // server := server.NewServer(config.Server)
    // server.Start()

    a := NewApp()
    a.Run()
}
