package main

import "fmt"

type Player struct {
	Name  string
	Score int
}

func (p *Player) AddScore(score int) {
	p.Score += score
}

func (p Player) DisplayInfo() {
	fmt.Printf("Nama Pemain : %s, Score : %d", p.Name, p.Score)
}

func main() {
	player := Player{
		Name:  "John",
		Score: 0,
	}

	player.AddScore(10)
	player.AddScore(5)
	player.DisplayInfo() // Output yang diharapkan: "Nama pemain: John, Skor: 15"
}
