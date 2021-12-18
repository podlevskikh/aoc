// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type SnailFishTree struct {
	LeftNode   *SnailFishTreeNode
	RightNode  *SnailFishTreeNode
	LeftValue  *SnailFishTreeValue
	RightValue *SnailFishTreeValue
}
type SnailFishTreeNode struct {
	Tree  *SnailFishTree
	Value *SnailFishTreeValue
}
type SnailFishTreeValue struct {
	Value      int
	LeftValue  *SnailFishTreeValue
	RightValue *SnailFishTreeValue
}

func NewSnailFishTree(in []string) (*SnailFishTree, *SnailFishTreeValue, []string) {
	if in[0] == "[" {
		leftTree, leftVal, rest := NewSnailFishTree(in[1:])
		ll := leftVal
		lr := leftVal
		if leftVal == nil {
			ll = leftTree.LeftValue
			lr = leftTree.RightValue
		}
		rightTree, rightVal, rest := NewSnailFishTree(rest[1:])
		rl := rightVal
		rr := rightVal
		if rightVal == nil {
			rl = rightTree.LeftValue
			rr = rightTree.RightValue
		}
		lr.RightValue = rl
		rl.LeftValue = lr
		return &SnailFishTree{
			LeftNode: &SnailFishTreeNode{
				Tree:  leftTree,
				Value: leftVal,
			},
			RightNode: &SnailFishTreeNode{
				Tree:  rightTree,
				Value: rightVal,
			},
			LeftValue:  ll,
			RightValue: rr,
		}, nil, rest[1:]
	}
	val, _ := strconv.Atoi(in[0])
	value := &SnailFishTreeValue{Value: val}
	return nil, value, in[1:]
}

func (t *SnailFishTree) print() string {
	return "[" + t.LeftNode.print() + "," + t.RightNode.print() + "]"
}

func (n *SnailFishTreeNode) print() string {
	if n.Tree == nil {
		return strconv.Itoa(n.Value.Value)
	}
	return n.Tree.print()
}

func (t *SnailFishTree) add(a *SnailFishTree) *SnailFishTree {
	t.printDirections()
	t.RightValue.RightValue = a.LeftValue
	a.LeftValue.LeftValue = t.RightValue
	res := &SnailFishTree{
		LeftNode: &SnailFishTreeNode{
			Tree: t,
		},
		RightNode: &SnailFishTreeNode{
			Tree: a,
		},
		LeftValue:  t.LeftValue,
		RightValue: a.RightValue,
	}
	fmt.Println(res.print())
	res.normalize()
	return res
}

func (t *SnailFishTree) normalize() {
	for {
		processed := false
		for {
			if t.explodeAll(3) {
				fmt.Println(t.print())
				t.printDirections()
				processed = true
				continue
			}
			break
		}
		//for {
			if t.splitAll() {
				processed = true
				fmt.Println(t.print())
				t.printDirections()
				continue
			}
		//	break
		//}
		if !processed {
			break
		}
	}
}

func (t *SnailFishTree) explodeAll(depth int) bool {
	if depth == 0 && t.LeftNode.explode() {
		t.LeftValue = t.LeftNode.Value
		return true
	}
	if t.LeftNode.Tree != nil && t.LeftNode.Tree.explodeAll(depth - 1) {
		t.LeftValue = t.LeftNode.Tree.LeftValue
		return true
	}
	if depth == 0 && t.RightNode.explode() {
		t.RightValue = t.RightNode.Value
		return true
	}
	if t.RightNode.Tree != nil && t.RightNode.Tree.explodeAll(depth - 1) {
		t.RightValue = t.RightNode.Tree.RightValue
		return true
	}
	return false
}

func (t *SnailFishTree) splitAll() bool {
	if t.LeftNode.split() {
		t.LeftValue = t.LeftNode.Tree.LeftValue
		return true
	}
	if t.LeftNode.Tree != nil && t.LeftNode.Tree.splitAll() {
		t.LeftValue = t.LeftNode.Tree.LeftValue
		return true
	}
	if t.RightNode.split() {
		t.RightValue = t.RightNode.Tree.RightValue
		return true
	}
	if t.RightNode.Tree != nil && t.RightNode.Tree.splitAll() {
		t.RightValue = t.RightNode.Tree.RightValue
		return true
	}
	return false
}

func (t *SnailFishTree) printDirections() {
	l := t.LeftValue
	lres := ""
	for {
		if l == nil {
			break
		}
		lres += strconv.Itoa(l.Value)
		l = l.RightValue
	}
	fmt.Println(lres)

	r := t.RightValue
	rres := ""
	for {
		if r == nil {
			break
		}
		rres = strconv.Itoa(r.Value) + rres
		r = r.LeftValue
	}
	fmt.Println(rres)
}

func (n *SnailFishTreeNode) explode() bool {
	if n.Tree != nil {
		n.Value = createValueFromTree(n.Tree)
		n.Tree = nil
		return true
	}
	return false
}

func createValueFromTree(t *SnailFishTree) *SnailFishTreeValue {
	newValue := &SnailFishTreeValue{
		Value:      0,
		LeftValue:  t.LeftValue.LeftValue,
		RightValue: t.RightValue.RightValue,
	}
	if t.LeftValue.LeftValue != nil {
		t.LeftValue.LeftValue.Value += t.LeftValue.Value
		t.LeftValue.LeftValue.RightValue = newValue
	}

	if t.RightValue.RightValue != nil {
		t.RightValue.RightValue.Value += t.RightValue.Value
		t.RightValue.RightValue.LeftValue = newValue
	}
	return newValue
}

func (t *SnailFishTreeNode) split() bool {
	if t.Tree == nil && t.Value.Value >= 10 {
		t.Tree = createTreeFromValue(t.Value)
		t.Value = nil
		return true
	}
	return false
}

func createTreeFromValue(v *SnailFishTreeValue) *SnailFishTree {
	newLeftValue := &SnailFishTreeValue{
		Value:     v.Value / 2,
		LeftValue: v.LeftValue,
	}
	if v.LeftValue != nil {
		v.LeftValue.RightValue = newLeftValue
	}
	newRightValue := &SnailFishTreeValue{
		Value:      v.Value - v.Value/2,
		RightValue: v.RightValue,
	}
	if v.RightValue != nil {
		v.RightValue.LeftValue = newRightValue
	}
	newLeftValue.RightValue = newRightValue
	newRightValue.LeftValue = newLeftValue
	return &SnailFishTree{
		LeftNode: &SnailFishTreeNode{
			Tree:  nil,
			Value: newLeftValue,
		},
		RightNode: &SnailFishTreeNode{
			Tree:  nil,
			Value: newRightValue,
		},
		LeftValue:  newLeftValue,
		RightValue: newRightValue,
	}
}

func (n *SnailFishTreeNode) magnitude() int {
	if n.Tree != nil {
		return n.Tree.magnitude()
	} else {
		return n.Value.Value
	}
}

func (t *SnailFishTree) magnitude() int {
	return 3 * t.LeftNode.magnitude() + 2 * t.RightNode.magnitude()
}

func main() {
	//-------------PART 1-------------
	var res *SnailFishTree
	for _, inputLine := range strings.Split(input, `
`) {
		t, _, _ := NewSnailFishTree(strings.Split(inputLine, ""))
		t.normalize()
		if res == nil {
			res = t
		} else {
		fmt.Println(inputLine + "!!!")
			res = res.add(t)
		}
	}
	fmt.Println(res.magnitude())

	//-------------PART 2-------------
	maxMagnitude := 0
	for i, in1 := range strings.Split(input, `
`) {
		for j, in2 := range strings.Split(input, `
`) {
			if i == j {
				continue
			}
			t1, _, _ := NewSnailFishTree(strings.Split(in1, ""))
			t2, _, _ := NewSnailFishTree(strings.Split(in2, ""))

			sum := t1.add(t2)
			if sum.magnitude() > maxMagnitude {
				maxMagnitude = sum.magnitude()
			}
		}
	}
	fmt.Println(maxMagnitude)

}

//var snailFishInput = "[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]"
//var snailFishInput = "[[[[[9,8],1],2],3],4]"
//var snailFishInput = "[7,[6,[5,[4,[3,2]]]]]"
//var snailFishInput = "[[6,[5,[4,[3,2]]]],1]"
//var snailFishInput = "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"
//var snailFishInput1 = "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"



/*var input = `[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]`*/

/*var input = `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`*/

var input = `[[3,[[6,3],[9,6]]],[6,[[0,9],[9,7]]]]
[[[3,9],[[0,8],[7,6]]],[[[7,9],1],[1,3]]]
[8,[[[9,6],[8,4]],4]]
[5,[[1,2],[3,7]]]
[[[[7,7],5],[[3,5],8]],4]
[[[5,[0,7]],3],[[5,[5,3]],[1,[9,4]]]]
[[[[3,5],[7,1]],6],[[[3,6],[5,6]],[[3,2],5]]]
[[[[2,0],[3,0]],[5,7]],[[4,4],[[9,9],[9,3]]]]
[[[[8,0],7],[[7,1],9]],[[3,[8,6]],8]]
[[6,[7,5]],[[6,8],9]]
[[[9,[1,8]],2],[[[4,0],[9,3]],1]]
[[7,[1,[3,8]]],[[4,7],[8,1]]]
[[[5,5],[[4,5],[2,9]]],[[[7,7],0],8]]
[[[[4,7],3],5],[[[4,3],[3,8]],[[6,5],5]]]
[[[[3,8],2],[1,7]],[[[3,1],4],9]]
[[[[2,1],4],[[9,5],[1,4]]],[[3,5],[[9,1],9]]]
[[[6,[1,8]],[0,0]],[9,[0,3]]]
[[[[2,2],[3,3]],[[4,8],4]],[[[6,8],4],5]]
[4,[[[7,8],[3,4]],[[3,2],9]]]
[[[9,0],3],[[[7,1],4],7]]
[[[1,4],8],[[7,5],[[8,0],[0,7]]]]
[9,[[4,6],[[2,9],1]]]
[[[[1,8],8],6],[[[2,0],6],[0,5]]]
[[[5,5],[6,4]],[[3,8],[9,[7,6]]]]
[[0,[8,[1,4]]],2]
[[[[9,5],0],5],[9,[7,5]]]
[[9,[4,8]],[[8,1],[[8,6],[7,1]]]]
[4,[[[9,6],5],9]]
[[[[3,7],6],0],[[7,7],[[2,7],[9,3]]]]
[[[6,[3,7]],[[8,3],2]],[8,[6,[8,5]]]]
[[[5,[2,7]],[[6,7],3]],[5,[[4,4],1]]]
[[1,0],[[2,8],[[0,4],9]]]
[[[1,4],6],[[[9,8],[1,0]],1]]
[[3,4],[[1,[8,4]],8]]
[[[[9,4],[0,7]],[[5,4],[8,2]]],2]
[5,[[[8,7],[3,4]],[2,4]]]
[[[[1,3],[8,6]],[[3,4],6]],[[8,5],[[9,3],[5,7]]]]
[[0,[[0,9],[7,8]]],[3,9]]
[0,[[8,[2,3]],[[3,5],[4,9]]]]
[[[4,3],[[1,9],[1,5]]],[4,[[9,1],1]]]
[[[[3,6],[2,5]],3],[[8,[8,0]],[[6,9],[5,8]]]]
[7,[[3,[3,6]],[[6,9],[2,7]]]]
[[[[8,3],[6,5]],[[3,9],2]],[6,1]]
[[[2,0],[2,3]],8]
[[1,[[8,7],2]],[[[9,4],8],[4,[9,0]]]]
[[[6,7],[[5,2],3]],[[0,5],[[9,4],[2,6]]]]
[[[9,[5,8]],[[9,3],[6,9]]],5]
[[[5,[4,6]],[5,[3,2]]],[2,[9,[5,4]]]]
[8,6]
[[[4,8],[3,1]],[1,[[7,8],[7,5]]]]
[[4,[[8,8],4]],[5,[8,[3,9]]]]
[[[4,[9,0]],[[0,3],5]],[[5,[3,0]],[6,[2,3]]]]
[[[4,0],8],[[[4,0],7],[[9,6],3]]]
[[8,[[7,8],5]],[[[6,2],8],[1,[0,4]]]]
[[1,[[3,4],[0,8]]],[[6,5],3]]
[[5,2],[[8,6],[1,[9,7]]]]
[5,[6,[[1,3],[1,0]]]]
[[0,[[1,9],[5,6]]],[[[6,2],[5,1]],[[1,2],[1,0]]]]
[[[7,1],4],[[[0,3],3],[[4,8],1]]]
[[3,[9,[3,4]]],[1,[[0,0],[1,4]]]]
[1,[7,[1,[3,7]]]]
[[[0,[5,6]],[[7,4],[5,7]]],[[[6,8],[4,6]],9]]
[[[9,8],[7,[1,3]]],3]
[[[4,[0,3]],[[3,0],6]],[[2,[9,2]],1]]
[[[[1,9],[3,3]],[8,1]],5]
[[7,[5,2]],[[4,[0,1]],[3,3]]]
[[[6,6],[0,6]],[[3,[5,9]],[[4,2],[4,3]]]]
[[[7,[5,4]],[7,1]],9]
[[6,[5,2]],[[7,[0,5]],4]]
[[[8,1],[[7,6],[4,1]]],2]
[[[[4,3],[1,4]],[9,6]],[3,[[2,5],3]]]
[[[[9,3],[5,0]],1],[1,[[9,7],9]]]
[[[8,5],[5,9]],[2,[4,[0,0]]]]
[[[[7,9],2],[[8,8],[6,3]]],[7,[0,9]]]
[[[[6,6],[0,2]],[2,[9,0]]],[[0,9],[9,9]]]
[[[9,[1,3]],[6,5]],[[[1,1],8],[9,[7,2]]]]
[[8,[[8,4],6]],[[4,[5,9]],0]]
[[8,[5,[6,7]]],[[[1,9],9],[0,[0,9]]]]
[[9,[9,[7,3]]],[4,[4,7]]]
[[[[9,3],7],5],[[5,[8,5]],[0,[8,0]]]]
[[[5,[9,0]],[[7,4],[5,3]]],[3,[[1,1],[1,8]]]]
[[1,[[1,4],[5,9]]],[[[9,1],[6,5]],[9,[0,7]]]]
[[[[9,4],9],[5,3]],[[[4,2],[2,2]],[[1,0],0]]]
[[[6,[8,6]],9],[8,[[0,1],[9,7]]]]
[[2,0],[5,[[8,3],4]]]
[[[[0,2],0],8],[8,[[2,5],[8,2]]]]
[[[[7,4],8],[9,[7,5]]],[8,[7,[5,3]]]]
[[2,4],[3,[3,8]]]
[[5,4],[[0,[5,8]],[4,3]]]
[6,[[5,[4,7]],9]]
[[[2,[6,8]],[5,5]],[[[3,0],4],[[6,6],[0,1]]]]
[[[1,[4,2]],[[8,0],8]],[8,[[6,1],[0,0]]]]
[[9,[2,[3,3]]],[[2,6],[[5,2],[5,8]]]]
[[9,[4,4]],[[[8,6],1],2]]
[2,[[[0,7],7],[[7,8],5]]]
[[[4,0],[[1,1],[7,6]]],[[6,7],[[7,2],1]]]
[[[[2,5],0],[[9,5],9]],[6,[7,[6,1]]]]
[[[7,8],1],[[[6,2],0],[[9,7],[3,5]]]]
[[[9,1],0],[3,[[6,1],[6,9]]]]
[[[[9,0],0],[4,[7,0]]],[[6,[4,0]],[8,[4,2]]]]`

