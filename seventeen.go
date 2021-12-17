// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Map struct {
	TargetLeftX   int
	TargetRightX  int
	TargetTopY    int
	TargetBottomY int
}

type Game struct {
	Map   *Map
	ShotX int
	ShotY int
}

type Step struct {
	X int
	Y int
}

func NewGame(m *Map, shotX, shotY int) *Game {
	return &Game{
		Map:   m,
		ShotX: shotX,
		ShotY: shotY,
	}
}

func (g *Game) Play() *Step {
	s := &Step{0, 0}

	maxTopStep := s

	for {
		if g.isOutSide(s) {
			return nil
		}
		if g.isInTarget(s) {
			return maxTopStep
		}
		nextStep := g.NextStep(s)
		s = nextStep
		if maxTopStep.Y < s.Y {
			maxTopStep = s
		}
	}
}

func (g *Game) isOutSide(s *Step) bool {
	return s.X > g.Map.TargetRightX || s.Y < g.Map.TargetBottomY
}

func (g *Game) isInTarget(s *Step) bool {
	return s.X >= g.Map.TargetLeftX &&
		s.X <= g.Map.TargetRightX &&
		s.Y >= g.Map.TargetBottomY &&
		s.Y <= g.Map.TargetTopY
}

func (g *Game) NextStep(s *Step) *Step {
	if g.ShotX != 0 {
		defer func() { g.ShotX = g.ShotX - 1 }()
	}
	defer func() { g.ShotY = g.ShotY - 1 }()

	return &Step{
		X: s.X + g.ShotX,
		Y: s.Y + g.ShotY,
	}
}

type Res struct {
	Step  Step
	shotX int
	shotY int
	res   int
}

func main() {
	maxStep := Res{Step: Step{0, 0}, shotX: 0, shotY: 0}
	countSuccess := 0
	for y := -106; y < initY; y++ {
		for x := 0; x < initX; x++ {
			g := NewGame(m, x, y)
			gameMaxStep := g.Play()
			if gameMaxStep != nil {
				if gameMaxStep.Y > maxStep.Step.Y {
					maxStep = Res{Step: *gameMaxStep, shotX: x, shotY: y, res: x * y}
				}
				countSuccess++
			}
		}
	}

	fmt.Println(countSuccess)
	fmt.Println(maxStep)
}

/*var m = &Map{
	TargetLeftX:   20,
	TargetRightX:  30,
	TargetTopY:    -5,
	TargetBottomY: -10,
}
var initY = 11
var initX = 31*/

var m = &Map{
	TargetLeftX:   206,
	TargetRightX:  250,
	TargetTopY:    -57,
	TargetBottomY: -105,
}
var initY = 106
var initX = 251
