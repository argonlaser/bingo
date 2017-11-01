package store

import "github.com/argonlaser/bingo/gameShip/game"

// GameStore will contain the reference to the games
// or methods to reach those game references
type GameStore interface {
	Add(g *game.Game) error
	GetByGameID(gameID string) (*game.Game, error)
}
