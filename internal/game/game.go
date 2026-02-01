package game

import "github.com/pabloduke/naviterm"

type Game struct {
}

func New() *Game {
	return &Game{}
}

func (g *Game) Init() error {
	return naviterm.Init()
}

func (g *Game) Close() {
	naviterm.Close()
}

func (g *Game) Run() {
	naviterm.PrintTextWithSpinner(2, 2, "Paws in the Machine")
}
