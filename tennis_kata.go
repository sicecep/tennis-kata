package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	score := ""

	tempScore := 0

	if g.Player1.Point == g.Player2.Point {
		switch g.Player1.Point {
		case 0:
			score = "Love-All"
		case 1:
			score = "Fifteen-All"
		case 2:
			score = "Thirty-All"
		default:
			score = "Deuce"
		}
	} else if g.Player1.Point >= 4 || g.Player2.Point >= 4 {
		minusResult := g.Player1.Point - g.Player2.Point
		if minusResult == 1 {
			score = "Advantage " + g.Player1.Name
		} else if minusResult == -1 {
			score = "Advantage " + g.Player2.Name
		} else if minusResult >= 2 {
			score = "Win for " + g.Player1.Name
		} else {
			score = "Win for " + g.Player2.Name
		}
	} else {
		for i := 1; i < 3; i++ {
			if i == 1 {
				tempScore = g.Player1.Point
			} else {
				score += "-"
				tempScore = g.Player2.Point
			}
			switch tempScore {
			case 0:
				score += "Love"
			case 1:
				score += "Fifteen"
			case 2:
				score += "Thirty"
			case 3:
				score += "Forty"
			}
		}
	}
	return score
}

func (g TennisGame) printScore() string {
	a := g.Player1.Name + " " + strconv.Itoa(g.Player1.Point) + " - " + strconv.Itoa(g.Player2.Point) + " " + g.Player2.Name
	return a
}

func main() {
	fmt.Println("== Set player name ==")
	players := make([]string, 2)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 1; i < 3; i++ {
		fmt.Printf("Enter player %d name: ", i)
		scanner.Scan()
		players[i-1] = scanner.Text()
	}

	var player1 = &Player{players[0], 0}
	var player2 = &Player{players[1], 0}
	var g = TennisGame{player1, player2}

	fmt.Println("== Start game ==")
	score := g.GetScore()
	for strings.HasPrefix(score, "Win for") == false {
		fmt.Printf("Add point to player (1/2/q): ")
		scanner.Scan()
		switch scanner.Text() {
		case "1":
			g.AddPoint(player1)
		case "2":
			g.AddPoint(player2)
		case "q":
			fmt.Println("== Exit game ==")
			os.Exit(1)
		default:
			fmt.Println("wrong input")
		}
		score = g.GetScore()
		fmt.Println(g.printScore())
		fmt.Println(score)
	}
	fmt.Println("== End game ==")
}
