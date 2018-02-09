package store

import (
	"errors"

	"github.com/argonlaser/bingo/gameShip/game"
)

// LocalGameStore stores the game data in local hard drive
// if the go process exits, then all the game data will be gone
type LocalGameStore struct {
	gameList []*game.Game
}

// Add adds the game to the store
func (s *LocalGameStore) Add(g *game.Game) error {
	s.gameList = append(s.gameList, g)
	return nil
}

// GetByGameID gets game by gameID
func (s *LocalGameStore) GetByGameID(gameID string) (*game.Game, error) {
	var reqGame *game.Game

	for _, g := range s.gameList {
		if g.ID == gameID {
			reqGame = g
			break
		}
	}

	if reqGame == nil {
		return nil, errors.New("GameNotFound")
	}

	return reqGame, nil
}

// Remove removes the game from the store
func (s *LocalGameStore) Remove(g *game.Game) error {
	s.gameList = append(s.gameList, g)
	return nil
}
