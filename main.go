package main

import (
    "fmt"
    "log"
    "net/http"
	"witt_backend/login"
	"witt_backend/ws"
)


func setupRoutes() {
    http.HandleFunc("/api/login", login.Login)
    http.HandleFunc("/ws", ws.WsEndpoint)
}

func main() {
    fmt.Println("Inicio Witt Backend v1")
    setupRoutes()
    log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
