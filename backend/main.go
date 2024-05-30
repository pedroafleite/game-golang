package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pedroafleite/game-golang/backend/deck"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,

  // We'll need to check the origin of our connection
  // this will allow us to make requests from our React
  // development server to here.
  // For now, we'll do no checking and just allow any connection
  CheckOrigin: func(r *http.Request) bool { return true },
}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
    for {
    // read in a message
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
    // print out that message for clarity
        fmt.Println(string(p))

        if err := conn.WriteMessage(messageType, p); err != nil {
            log.Println(err)
            return
        }

    }
}

// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Host)

  // upgrade this connection to a WebSocket
  // connection
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
  }
  // listen indefinitely for new messages coming
  // through on our WebSocket connection
    reader(ws)
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
