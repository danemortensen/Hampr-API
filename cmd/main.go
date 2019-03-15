package main

import (
    "github.com/danemortensen/Hampr-API/pkg/config"
    "github.com/danemortensen/Hampr-API/pkg/db"
    "github.com/danemortensen/Hampr-API/pkg/server"
)

func main() {
    config := config.NewConfig()
    session := db.NewSession(config.Mongo)
    defer session.Close()
    server := server.NewServer(config.Server)
    server.Start()
}
