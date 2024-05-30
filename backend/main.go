package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pedroafleite/game-golang/backend/deck"
	"github.com/pedroafleite/game-golang/backend/websocket"
)

// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Host)
    
    ws, err := websocket.Upgrade(w, r)
    if err != nil {
        fmt.Fprintf(w, "%+V\n", err)
    }
    go websocket.Writer(ws)
    websocket.Reader(ws)
}

func setupRoutes(x *deck.Deck) {
    http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
        log.Println("Hello, World!")
        d, err := io.ReadAll(r.Body)
        if err != nil {
            http.Error(rw, "Oops", http.StatusBadRequest)
            return
        }

        log.Printf("Data %s\n", d)

        fmt.Fprintf(rw, "Hello %s\n", d)
        fmt.Fprintf(rw, "Your cards are %s\n", x[:3])
    })

    http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
        log.Println("Goodbye, World!")
    })

    http.HandleFunc("/ws", serveWs)
}

func main() {
    fmt.Println("Starting")
    var x0 = deck.New()
    x := &x0
    fmt.Println(x0)

    setupRoutes(x)
    http.ListenAndServe(":8080", nil)
}
