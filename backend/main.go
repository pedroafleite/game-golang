package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/pedroafleite/game-golang/backend/deck"
)

func main() {
    fmt.Println("Starting")
    x := deck.New()
    fmt.Println(x[:3])
    
    http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request) {
        log.Println("Hello, World!")
        d, _ := io.ReadAll(r.Body)
        log.Printf("Data %s\n", d)

        fmt.Fprintf(rw, "Hello %s\n", d)
        fmt.Fprintf(rw, "Your cards are %s\n", x[:3])
    })

    http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
        log.Println("Goodbye, World!")
    })

    http.ListenAndServe(":9000", nil)
}
