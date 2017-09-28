package main

import (
	"fmt"
	"strconv"
)

type Player struct {
	Name  string
	Point int
}

type TennisGame struct {
	Player1 *Player
	Player2 *Player
}

func (g TennisGame) AddPoint(player *Player) {
	player.Point += 1
}

func (g TennisGame) GetScore() string {

	baseScore := [4]string{"Love", "Fifteen", "Thirty", "Forty"}
	var score string

	if g.Player1.Point < 4 && g.Player2.Point < 4 && !(g.Player1.Point+g.Player2.Point == 6) {

		s := baseScore[g.Player1.Point]

		if g.Player1.Point == g.Player2.Point {
			score = s + "-All"
		} else {
			score = s + "-" + baseScore[g.Player2.Point]
		}
	} else if g.Player1.Point == g.Player2.Point {
		score = "Deuce"
	}

	return score
}

func (g TennisGame) printScore() string {
	a := g.Player1.Name + " " + strconv.Itoa(g.Player1.Point) + " - " + strconv.Itoa(g.Player2.Point) + " " + g.Player2.Name
	return a
}

func main() {
	var player1 = &Player{"Sepry", 0}
	var player2 = &Player{"Haryandi", 0}
	var g = TennisGame{player1, player2}

	g.AddPoint(player1)
	fmt.Println(g.printScore())
	fmt.Println(g.GetScore())
}
