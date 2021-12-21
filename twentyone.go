// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
)

type DiracDiceGame struct {
	p1Moves bool
	Count   int64
	p1      *Player
	p2      *Player
}

func (g *DiracDiceGame) Key() string {
	key := ""
	if g.p1Moves {
		key += "p1"
	} else {
		key += "p2"
	}
	return key + "_" + strconv.Itoa(g.p1.Position) + "_" + strconv.Itoa(g.p1.Score) +
		"_" + strconv.Itoa(g.p2.Position) + "_" + strconv.Itoa(g.p2.Score)
}

func (g *DiracDiceGame) Move() []*DiracDiceGame {
	if g.p1Moves {
		return g.Move1()
	} else {
		return g.Move2()
	}
}

func (g *DiracDiceGame) Move1() []*DiracDiceGame {
	gs := make([]*DiracDiceGame, 0, 7)
	moves := map[int]int64{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}

	for die, count := range moves {
		newGame := &DiracDiceGame{
			p1Moves: !g.p1Moves,
			Count:   g.Count * count,
			p1:      g.p1.MoveNewPlayer(die),
			p2:      g.p2,
		}
		gs = append(gs, newGame)
	}

	return gs
}

func (g *DiracDiceGame) Move2() []*DiracDiceGame {
	gs := make([]*DiracDiceGame, 0, 7)
	moves := map[int]int64{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}

	for die, count := range moves {
		newGame := &DiracDiceGame{
			p1Moves: !g.p1Moves,
			Count:   g.Count * count,
			p1:      g.p1,
			p2:      g.p2.MoveNewPlayer(die),
		}
		gs = append(gs, newGame)
	}

	return gs
}

type Player struct {
	Position int
	Score    int
}

func (p *Player) MoveDeterministic(die *Die) {
	for i := 0; i < 3; i++ {
		p.Position += die.Value
		p.Position -= (p.Position / 10) * 10
		if p.Position == 0 {
			p.Position = 10
		}
		die.Value += 1
		if die.Value == 101 {
			die.Value = 1
		}
		die.Rolled++
	}

	p.Score += p.Position
}

func (p *Player) MoveNewPlayer(die int) *Player {
	newP := &Player{
		Position: p.Position + die,
		Score:    p.Score,
	}
	if newP.Position > 10 {
		newP.Position -= 10
	}
	newP.Score += newP.Position
	return newP
}

type Die struct {
	Value  int
	Rolled int
}

func main() {
	//-----------------PART 1--------------------

	/*p1 := &Player{Position: 7, Score: 0}
	p2 := &Player{Position: 9, Score: 0}

	die := &Die{
		Value:  1,
		Rolled: 0,
	}
	movingPlayer := p1

	for {
		if p1.Score >= 1000 || p2.Score >= 1000 {
			break
		}
		movingPlayer.MoveDeterministic(die)
		if movingPlayer == p1 {
			movingPlayer = p2
		} else {
			movingPlayer = p1
		}
	}
	fmt.Println(p1)
	fmt.Println(p2)
	losingPlayer := p1
	if losingPlayer.Score >= 1000 {
		losingPlayer = p2
	}
	fmt.Println(losingPlayer.Score * die.Rolled)*/

	//-----------------PART 2--------------------

	p1 := &Player{Position: 7, Score: 0}
	p2 := &Player{Position: 9, Score: 0}

	g := &DiracDiceGame{
		p1Moves: true,
		Count:   1,
		p1:      p1,
		p2:      p2,
	}

	gamesToPlay := []*DiracDiceGame{g}

	finishedGames := []*DiracDiceGame{}
	for i := 0; i<40; i++ {
		newGames := map[string]*DiracDiceGame{}
		for _, ga := range gamesToPlay {
			for _, newGa := range ga.Move() {
				if _, ok := newGames[newGa.Key()]; ok {
					newGames[newGa.Key()].Count += newGa.Count
				} else {
					newGames[newGa.Key()] = newGa
				}
			}
		}
		gamesToPlay = []*DiracDiceGame{}
		for _, newGame := range newGames {
			if newGame.p1.Score >= 21 || newGame.p2.Score >= 21 {
				finishedGames = append(finishedGames, newGame)
			} else {
				gamesToPlay = append(gamesToPlay, newGame)
			}
		}
	}

	player1Win := int64(0)
	player2Win := int64(0)

	for _, ga := range finishedGames {
		if ga.p1.Score >= 21 {
			player1Win += ga.Count
		} else {
			player2Win += ga.Count
		}

	}
	fmt.Println(player1Win)
	fmt.Println(player2Win)

}
