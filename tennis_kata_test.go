package main

import (
	"fmt"
	"testing"
)

var player1 = &Player{"Sepry", 0}
var player2 = &Player{"Haryandi", 0}
var g = TennisGame{player1, player2}

// reset players point for every test
func Teardown() {
	player1.Point = 0
	player2.Point = 0
}

func TestAllPlayerLove(t *testing.T) {
	t.Log(g.printScore())
	t.Log(g.GetScore())

	expected := "Love-All"
	actual := g.GetScore()

	if actual != expected {
		t.Error(fmt.Sprintf("Expected result is %d but instead got %d!", expected, actual))
	}

	Teardown()
}

func TestFirstPlayerWinFirstGame(t *testing.T) {
	g.AddPoint(player1)

	t.Log(g.printScore())
	t.Log(g.GetScore())

	expected := "Fifteen-Love"
	actual := g.GetScore()

	if actual != expected {
		t.Error(fmt.Sprintf("Expected result is %d but instead got %d!", expected, actual))
	}

	Teardown()
}
