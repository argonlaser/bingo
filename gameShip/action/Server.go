package action

import (
	"github.com/argonlaser/bingo/gameShip/store"
	"go.uber.org/zap"
)

// GameShipRPCServer implments the gameShipRpc interface
type GameShipRPCServer struct {
	Logger   *zap.Logger
	GameStore store.GameStore
}
