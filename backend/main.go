package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/pedroafleite/game-golang/backend/deck"
	"github.com/pedroafleite/game-golang/backend/websocket"
)

// define our WebSocket endpoint
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
    fmt.Println("WebSocket Endpoint Hit")
    conn, err := websocket.Upgrade(w, r)
    if err != nil {
        fmt.Fprintf(w, "%+v\n", err)
    }

    client := &websocket.Client{
        Conn: conn,
        Pool: pool,
    }

    fmt.Println("here") 
    var x = deck.New()
    log.Printf("Your cards are %s\n", x[:3])
    if len(pool.Clients) >= 1 {            
        log.Printf("Your cards are %s\n", x[:3])
    }

    pool.Register <- client
    client.Read()
}

func setupRoutes() {
    http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
        log.Println("Hello, World!")
        d, err := io.ReadAll(r.Body)
        if err != nil {
            http.Error(rw, "Oops", http.StatusBadRequest)
            return
        }

        log.Printf("Data %s\n", d)

        fmt.Fprintf(rw, "Hello %s\n", d)
    })

    http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
        log.Println("Goodbye, World!")
    })

    pool := websocket.NewPool()
    go pool.Start()

    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serveWs(pool, w, r)
    })
}

func main() {

    fmt.Println("Distributed Chat App v0.01")
    setupRoutes()
    http.ListenAndServe(":8080", nil)
}
