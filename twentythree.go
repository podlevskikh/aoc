// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

const maxLen = 4

type Flat struct {
	RoomA   *Room
	RoomB   *Room
	RoomC   *Room
	RoomD   *Room
	Hallway []*HallwayPart
	Weight  int
}

func (f *Flat) GetAmphipods() Amphipods {
	as := Amphipods{}
	for _, a := range f.RoomA.AmphipodStack {
		as = append(as, a)
	}
	for _, a := range f.RoomB.AmphipodStack {
		as = append(as, a)
	}
	for _, a := range f.RoomC.AmphipodStack {
		as = append(as, a)
	}
	for _, a := range f.RoomD.AmphipodStack {
		as = append(as, a)
	}
	for _, h := range f.Hallway {
		if h.Amphipod != nil {
			as = append(as, h.Amphipod)
		}
	}
	return as
}

func (f *Flat) InHallway(a *Amphipod) bool {
	for _, h := range f.Hallway {
		if h.Amphipod != nil && h.Amphipod == a {
			return true
		}
	}

	return false
}

type Room struct {
	Name          string
	AmphipodStack map[int]*Amphipod
	MaxStack      int
}

func (r *Room) SameName(a *Amphipod) bool {
	return r.Name == strings.Split(a.Name, "")[0]
}

func (a *Amphipod) SameName(a1 *Amphipod) bool {
	return strings.Split(a.Name, "")[0] == strings.Split(a1.Name, "")[0]
}

func (r *Room) CanEnter(a *Amphipod) bool {
	if !r.SameName(a) {
		return false
	}

	if len(r.AmphipodStack) >= r.MaxStack {
		return false
	}

	if len(r.AmphipodStack) == 0 {
		return true
	}

	for i := 0 ; i < len(r.AmphipodStack); i++ {
		if !r.AmphipodStack[i].SameName(a) {
			return false
		}
	}

	return true
}

func (r *Room) CanLeave(a *Amphipod) bool {
	if r.AmphipodStack[len(r.AmphipodStack)-1] == a {
		return true
	}

	return false
}

type HallwayPart struct {
	Index        int
	UnderTheRoom *Room
	Amphipod     *Amphipod
}

type Amphipods []*Amphipod

type Amphipod struct {
	Name string
}

func (a Amphipods) Len() int {
	return len(a)
}

func (a Amphipods) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

func (a Amphipods) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a *Amphipod) InFinishPosition(f *Flat) bool {
	for _, r := range []*Room{f.RoomA, f.RoomB, f.RoomC, f.RoomD} {
		if !r.SameName(a) {
			continue
		}
		for i := 0; i < len(r.AmphipodStack); i++ {
			if r.AmphipodStack[i] == a {
				return true
			}
			if !r.AmphipodStack[i].SameName(a) {
				break
			}
		}
	}
	return false
}

func (a *Amphipod) RoomIn(f *Flat) *Room {
	for _, r := range []*Room{f.RoomA, f.RoomB, f.RoomC, f.RoomD} {
		for _, ra := range r.AmphipodStack {
			if ra == a {
				return r
			}
		}
	}
	return nil
}

func (a *Amphipod) HallwayIn(f *Flat) *HallwayPart {
	for _, h := range f.Hallway {
		if h.Amphipod == a {
			return h
		}
	}
	return nil
}

func (a *Amphipod) CanPassToTargetRoom(f *Flat) *Room {
	var targetRoom *Room
	for _, r := range []*Room{f.RoomA, f.RoomB, f.RoomC, f.RoomD} {
		if r.SameName(a) {
			targetRoom = r
		}
	}
	var hallwayUnderTargetRoom *HallwayPart
	for _, h := range f.Hallway {
		if h.UnderTheRoom == targetRoom {
			hallwayUnderTargetRoom = h
		}
	}

	var hallwayForAmphipod *HallwayPart
	roomIn := a.RoomIn(f)
	for _, h := range f.Hallway {
		if roomIn != nil && h.UnderTheRoom == roomIn || h.Amphipod == a {
			hallwayForAmphipod = h
		}
	}

	if isFreeWayFromHallwayToHallway(hallwayForAmphipod, hallwayUnderTargetRoom, f) {
		return targetRoom
	}
	return nil
}

func (a *Amphipod) getPossibleHallways(f *Flat) []*HallwayPart {
	for _, h := range f.Hallway {
		if h.Amphipod == a {
			return []*HallwayPart{}
		}
	}

	var hallwayForAmphipod *HallwayPart
	roomIn := a.RoomIn(f)
	for _, h := range f.Hallway {
		if roomIn != nil && h.UnderTheRoom == roomIn {
			hallwayForAmphipod = h
		}
	}

	hs := []*HallwayPart{}
	for _, h := range f.Hallway {
		if h.Amphipod != nil {
			continue
		}
		if h.UnderTheRoom != nil {
			continue
		}
		if isFreeWayFromHallwayToHallway(hallwayForAmphipod, h, f) {
			hs = append(hs, h)
		}
	}

	return hs
}

func isFreeWayFromHallwayToHallway(from *HallwayPart, to *HallwayPart, f *Flat) bool {
	for _, h := range f.Hallway {
		if h.Index > int(math.Min(float64(from.Index), float64(to.Index))) &&
			h.Index < int(math.Max(float64(from.Index), float64(to.Index))) {
			if h.Amphipod != nil {
				return false
			}
		}
	}
	return true
}

func main() {
	flat := getStartPosition()
	flat.print()

	getFinishFlats(flat, 1)
}

var optimalFinish *Flat

func getFinishFlats(f *Flat, depth int) {
	fl := getPossibleFlats(f)
	for _, pf := range fl {
		if optimalFinish != nil && optimalFinish.Weight <= pf.Weight {
			continue
		}
		if pf.Weight > 60000 {
			continue
		}
		if pf.IsFinish() {
			if optimalFinish == nil || optimalFinish.Weight > pf.Weight {
				optimalFinish = pf
				optimalFinish.print()
			}
		} else {
			getFinishFlats(pf, depth + 1)
		}
	}
}

func (f *Flat) IsFinish() bool {
	for _, a := range f.GetAmphipods() {
		if !a.InFinishPosition(f) {
			return false
		}
	}
	return true
}

func getPossibleFlats(flat *Flat) []*Flat {
	flats := []*Flat{}
	as := flat.GetAmphipods()
	sort.Sort(as)
	for _, a := range as {
		if a.InFinishPosition(flat) {
			continue
		}
		r := a.RoomIn(flat)
		if r != nil && !r.CanLeave(a) {
			continue
		}

		targetR := a.CanPassToTargetRoom(flat)
		if targetR != nil && targetR.CanEnter(a) {
			fm := flat.moveToRoom(a, targetR)
			flats = append(flats, fm)
			continue
		}
		if r != nil {
			for _, h := range a.getPossibleHallways(flat) {
				fm := flat.moveToHallway(a, h)
				flats = append(flats, fm)
			}
		}
	}
	return flats
}

func getStartPosition() *Flat {
	roomA := &Room{
		Name: "A",
		AmphipodStack: map[int]*Amphipod{
			3: &Amphipod{"D4"},
			2: &Amphipod{"D3"},
			1: &Amphipod{"D2"},
			0: &Amphipod{"D1"},
		},
		MaxStack: maxLen,
	}
	roomB := &Room{
		Name: "B",
		AmphipodStack: map[int]*Amphipod{
			3: &Amphipod{"A1"},
			2: &Amphipod{"C2"},
			1: &Amphipod{"B1"},
			0: &Amphipod{"C1"},
		},
		MaxStack: maxLen,
	}
	roomC := &Room{
		Name: "C",
		AmphipodStack: map[int]*Amphipod{
			3: &Amphipod{"C3"},
			2: &Amphipod{"B2"},
			1: &Amphipod{"A2"},
			0: &Amphipod{"B3"},
		},
		MaxStack: maxLen,
	}
	roomD := &Room{
		Name: "D",
		AmphipodStack: map[int]*Amphipod{
			3: &Amphipod{"A4"},
			2: &Amphipod{"A3"},
			1: &Amphipod{"C4"},
			0: &Amphipod{"B4"},
		},
		MaxStack: maxLen,
	}

/*
	roomA := &Room{
		Name: "A",
		AmphipodStack: map[int]*Amphipod{
			3: &Amphipod{"B1"},
			2: &Amphipod{"D1"},
			1: &Amphipod{"D2"},
			0: &Amphipod{"A1"},
		},
		MaxStack: maxLen,
	}
	roomB := &Room{
		Name: "B",
		AmphipodStack: map[int]*Amphipod{
			3: &Amphipod{"C1"},
			2: &Amphipod{"C2"},
			1: &Amphipod{"B2"},
			0: &Amphipod{"D3"},
		},
		MaxStack: maxLen,
	}
	roomC := &Room{
		Name: "C",
		AmphipodStack: map[int]*Amphipod{
			3: &Amphipod{"B3"},
			2: &Amphipod{"B4"},
			1: &Amphipod{"A2"},
			0: &Amphipod{"C3"},
		},
		MaxStack: maxLen,
	}
	roomD := &Room{
		Name: "D",
		AmphipodStack: map[int]*Amphipod{
			3: &Amphipod{"D4"},
			2: &Amphipod{"A3"},
			1: &Amphipod{"C4"},
			0: &Amphipod{"A4"},
		},
		MaxStack: maxLen,
	}*/

	return &Flat{
		RoomA: roomA,
		RoomB: roomB,
		RoomC: roomC,
		RoomD: roomD,
		Hallway: []*HallwayPart{
			&HallwayPart{
				Index: 0,
			},
			&HallwayPart{
				Index: 1,
			},
			&HallwayPart{
				Index:        2,
				UnderTheRoom: roomA,
			},
			&HallwayPart{
				Index: 3,
			},
			&HallwayPart{
				Index:        4,
				UnderTheRoom: roomB,
			},
			&HallwayPart{
				Index: 5,
			},
			&HallwayPart{
				Index:        6,
				UnderTheRoom: roomC,
			},
			&HallwayPart{
				Index: 7,
			},
			&HallwayPart{
				Index:        8,
				UnderTheRoom: roomD,
			},
			&HallwayPart{
				Index: 9,
			},
			&HallwayPart{
				Index: 10,
			},
		},
	}
}

func getEmpty() *Flat {
	roomA := &Room{
		Name:          "A",
		AmphipodStack: map[int]*Amphipod{},
		MaxStack:      maxLen,
	}
	roomB := &Room{
		Name:          "B",
		AmphipodStack: map[int]*Amphipod{},
		MaxStack:      maxLen,
	}
	roomC := &Room{
		Name:          "C",
		AmphipodStack: map[int]*Amphipod{},
		MaxStack:      maxLen,
	}
	roomD := &Room{
		Name:          "D",
		AmphipodStack: map[int]*Amphipod{},
		MaxStack:      maxLen,
	}

	return &Flat{
		RoomA: roomA,
		RoomB: roomB,
		RoomC: roomC,
		RoomD: roomD,
		Hallway: []*HallwayPart{
			&HallwayPart{
				Index: 0,
			},
			&HallwayPart{
				Index: 1,
			},
			&HallwayPart{
				Index:        2,
				UnderTheRoom: roomA,
			},
			&HallwayPart{
				Index: 3,
			},
			&HallwayPart{
				Index:        4,
				UnderTheRoom: roomB,
			},
			&HallwayPart{
				Index: 5,
			},
			&HallwayPart{
				Index:        6,
				UnderTheRoom: roomC,
			},
			&HallwayPart{
				Index: 7,
			},
			&HallwayPart{
				Index:        8,
				UnderTheRoom: roomD,
			},
			&HallwayPart{
				Index: 9,
			},
			&HallwayPart{
				Index: 10,
			},
		},
	}
}

func copyRoomAmphipods(from *Room, to *Room, exclude *Amphipod) {
	for i := 0; i < len(from.AmphipodStack); i++ {
		if from.AmphipodStack[i] != exclude {
			to.AmphipodStack[len(to.AmphipodStack)] = &Amphipod{Name: from.AmphipodStack[i].Name}
		}
	}
}

func (f *Flat) moveToHallway(a *Amphipod, part *HallwayPart) *Flat {
	pNew := 0
	newF := getEmpty()
	copyRoomAmphipods(f.RoomA, newF.RoomA, a)
	copyRoomAmphipods(f.RoomB, newF.RoomB, a)
	copyRoomAmphipods(f.RoomC, newF.RoomC, a)
	copyRoomAmphipods(f.RoomD, newF.RoomD, a)

	for _, h := range newF.Hallway {
		if h.Index == part.Index {
			pNew = h.Index
			h.Amphipod = &Amphipod{Name: a.Name}
			continue
		}
		for _, ho := range f.Hallway {
			if h.Index == ho.Index {
				if ho.Amphipod != nil {
					h.Amphipod = &Amphipod{Name: ho.Amphipod.Name}
					break
				}
			}
		}
	}

	p := f.getAmphipodWeight(a.Name)

	move := 0
	for {
		if p <= 10 {
			break
		}
		p -= 10
		move += 1
	}
	move += int(math.Abs(float64(p - pNew)))

	newF.Weight = f.Weight + move*getCoef(a.Name)

	return newF
}

func (f *Flat) moveToRoom(a *Amphipod, r *Room) *Flat {
	newF := getEmpty()
	copyRoomAmphipods(f.RoomA, newF.RoomA, a)
	copyRoomAmphipods(f.RoomB, newF.RoomB, a)
	copyRoomAmphipods(f.RoomC, newF.RoomC, a)
	copyRoomAmphipods(f.RoomD, newF.RoomD, a)

	if newF.RoomA.Name == r.Name {
		newF.RoomA.AmphipodStack[len(newF.RoomA.AmphipodStack)] = &Amphipod{Name: a.Name}
	}
	if newF.RoomB.Name == r.Name {
		newF.RoomB.AmphipodStack[len(newF.RoomB.AmphipodStack)] = &Amphipod{Name: a.Name}
	}
	if newF.RoomC.Name == r.Name {
		newF.RoomC.AmphipodStack[len(newF.RoomC.AmphipodStack)] = &Amphipod{Name: a.Name}
	}
	if newF.RoomD.Name == r.Name {
		newF.RoomD.AmphipodStack[len(newF.RoomD.AmphipodStack)] = &Amphipod{Name: a.Name}
	}

	for _, h := range newF.Hallway {
		for _, ho := range f.Hallway {
			if h.Index == ho.Index {
				if ho.Amphipod != nil && ho.Amphipod != a {
					h.Amphipod = &Amphipod{Name: ho.Amphipod.Name}
					break
				}
			}
		}
	}

	p := f.getAmphipodWeight(a.Name)
	pNew := newF.getAmphipodWeight(a.Name)

	move := 0
	for {
		if p <= 10 {
			break
		}
		p -= 10
		move += 1
	}
	for {
		if pNew <= 10 {
			break
		}
		pNew -= 10
		move += 1
	}
	move += int(math.Abs(float64(p - pNew)))

	newF.Weight = f.Weight + move*getCoef(a.Name)

	return newF
}

func getCoef(name string) int {
	switch strings.Split(name, "")[0] {
	case "A":
		return 1
	case "B":
		return 10
	case "C":
		return 100
	case "D":
		return 1000
	}
	return 0
}

func (f *Flat) getAmphipodWeight(a string) int {
	for i, ar := range f.RoomA.AmphipodStack {
		if ar.Name == a {
			return 2 + 10*(f.RoomA.MaxStack-i)
		}
	}
	for i, ar := range f.RoomB.AmphipodStack {
		if ar.Name == a {
			return 4 + 10*(f.RoomB.MaxStack-i)
		}
	}
	for i, ar := range f.RoomC.AmphipodStack {
		if ar.Name == a {
			return 6 + 10*(f.RoomC.MaxStack-i)
		}
	}
	for i, ar := range f.RoomD.AmphipodStack {
		if ar.Name == a {
			return 8 + 10*(f.RoomD.MaxStack-i)
		}
	}
	for _, ho := range f.Hallway {
		if ho.Amphipod != nil && ho.Amphipod.Name == a {
			return ho.Index
		}
	}
	return 0
}

func (f *Flat) print() {
	fmt.Println("Weight", f.Weight)
	fmt.Println("#############")
	l := "#"
	for _, h := range f.Hallway {
		if h.Amphipod == nil {
			l += "."
		} else {
			l += strings.Split(h.Amphipod.Name, "")[0]
		}
	}
	fmt.Println(l + "#")
	for i := maxLen - 1; i >= 0; i-- {
		l = "#"
		for _, h := range f.Hallway {
			if h.UnderTheRoom == nil {
				l += "#"
			} else if _, ok := h.UnderTheRoom.AmphipodStack[i]; !ok {
				l += "."
			} else {
				l += strings.Split(h.UnderTheRoom.AmphipodStack[i].Name, "")[0]
			}
		}
		fmt.Println(l + "#")
	}
	fmt.Println("#############")
}
