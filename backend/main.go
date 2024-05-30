package main

import (
	"fmt"

	"github.com/pedroafleite/game-golang/backend/deck"
)

func main() {
    fmt.Println("Hello, World!")
    x := deck.New()
    fmt.Println(x)
}
