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
    garmentService := db.NewGarmentService(session.Copy(), config.Mongo)
    outfitService := db.NewOutfitService(session.Copy(), config.Mongo)
    userService := db.NewUserService(session.Copy(), config.Mongo)

    server := server.NewServer(garmentService, outfitService, userService, config.Server)

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
    a := NewApp()
    a.Run()
}
