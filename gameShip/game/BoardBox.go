package game

import "math/rand"

// BoardBox instance
type BoardBox struct {
	IsStriked bool
	Row       int32
	Col       int32
	Val       int32
}

// fillBoardBox gets player in a game by ID
func (b *BoardBox) fillBoardBox() (*BoardBox, error) {
	list := rand.Perm(25)
	for i, _ := range list {
		list[i]++
	}

	var i = 0
	for rowVal := 1; rowVal <= 5; rowVal++ {
		for colVal := 1; colVal <= 5; colVal++ {
			b.IsStriked = false
			b.Row = int32(rowVal)
			b.Col = int32(colVal)
			b.Val = int32(list[i-1])
			i++
		}
	}
}
