package main

import (
   "github.com/danemortensen/hampr/pkg/server"
)

func main() {
   s := server.NewServer()
   s.Start()
}
