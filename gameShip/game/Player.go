package game

import "errors"

// Player is an instance of the Player
type Player struct {
	ID      string
	Board   [25]BoardBox
	Bingoed bool
}

// NewPlayer gives new player instance
func NewPlayer(ID string, board []BoardBox) *Player {
	var player = &Player{
		ID: ID,
	}

	var i = 0
	for rowVal := 1; rowVal <= 5; rowVal++ {
		for colVal := 1; colVal <= 5; colVal++ {
			player.Board[i].IsStriked = board[i].IsStriked
			player.Board[i].Row = board[i].Row
			player.Board[i].Col = board[i].Col
			player.Board[i].Val = board[i].Val
			i++
		}
	}

	return player
}

// Strike a box in a player's board
func (p *Player) Strike(row, col int32) error {
	boxToBeStriked, err := p.getGameBoardBox(row, col)

	if err != nil {
		return err
	}

	if boxToBeStriked.IsStriked {
		return errors.New("Could not strike, already striked value")
	}

	boxToBeStriked.IsStriked = true
	return nil
}

func (p *Player) canBingo() bool {
	var strikedCount = 0

	for rowVal := 1; rowVal <= 5; rowVal++ {
		var isRowFullyStriked = true

		for colVal := 1; colVal <= 5; colVal++ {
			var box, err = p.getGameBoardBox(int32(rowVal), int32(colVal))
			if err == nil {
				if !box.IsStriked {
					isRowFullyStriked = false
					break
				}
			}
		}

		if isRowFullyStriked {
			strikedCount++
		}
	}

	if strikedCount >= 5 {
		return true
	}

	for colVal := 1; colVal <= 5; colVal++ {
		var isColFullyStriked = true

		for rowVal := 1; rowVal <= 5; rowVal++ {
			var box, err = p.getGameBoardBox(int32(rowVal), int32(colVal))
			if err == nil {
				if !box.IsStriked {
					isColFullyStriked = false
					break
				}
			}
		}

		if isColFullyStriked {
			strikedCount++
		}
	}

	if strikedCount >= 5 {
		return true
	}

	return false
}

func (p *Player) getGameBoardBox(row, col int32) (*BoardBox, error) {
	for i := 0; i < len(p.Board); i++ {
		if (p.Board[i].Row == row) && (p.Board[i].Col == col) {
			return &p.Board[i], nil
		}
	}

	return nil, errors.New("GameBoardBox is not found")
}

// Bingo allows user to bingo
func (p *Player) Bingo() error {
	if p.Bingoed {
		return errors.New("Already bingoed player. Could not bingo more than once.")
	}

	if p.canBingo() {
		p.Bingoed = true
		return nil
	}

	return errors.New("Not enough striked count")
}
