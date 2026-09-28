package main

import "fmt"

type Player struct {
	Name  string
	Score int
}

// Passes by value (creates a copy)
func updateScoreCopy(p Player) {
	p.Score += 10
}

// Passes by reference (modifies original)
func updateScorePointer(p *Player) {
	p.Score += 10
}

func main() {
	player1 := Player{Name: "Alex", Score: 50}

	updateScoreCopy(player1)
	fmt.Println("After Copy Update:", player1.Score) // Still 50

	updateScorePointer(&player1)
	fmt.Println("After Pointer Update:", player1.Score) // Now 60
}