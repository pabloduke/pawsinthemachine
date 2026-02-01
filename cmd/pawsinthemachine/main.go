package main

import (
	"github.com/pabloduke/pawsinthemachine/internal/game"
)

func main() {
	g := game.New()

	err := g.Init()
	if err != nil {
		panic(err)
	}
	defer g.Close()

	g.Run()
}
