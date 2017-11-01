package action

import (
	"golang.org/x/net/context"

	"github.com/argonlaser/bingo/gameShip/game"
	"github.com/argonlaser/bingo/gameShip/gameShipRpc"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

// NewGame creates a Game object and saves in a global store
func (s *GameShipRPCServer) NewGame(ctx context.Context, in *gameShipRpc.NewGameRequest) (*gameShipRpc.NewGameResponse, error) {
	var funcName = zap.String("FuncName", "actions|NewGame|")
	s.Logger.Info("funcEvent", funcName, zap.String("Event", "Start"))
	s.Logger.Debug("funcVariable", funcName, zap.Any("NewGameRequest", in))

	var creator = game.NewPlayer(in.CreatorId)

	var newGame = &game.Game{
		ID:      uuid.NewV4().String(),
		Players: []*game.Player{creator},
		Creator: creator,
	}

	if err := s.GameStore.Add(newGame); err != nil {
		s.Logger.Error("funcError", funcName, zap.Error(err))
		return nil, err
	}

	var result = &gameShipRpc.NewGameResponse{
		GameId: newGame.ID,
	}

	s.Logger.Debug("funcVariable", funcName, zap.Any("NewGameResponse", result))
	s.Logger.Info("funcEvent", funcName, zap.String("Event", "End"))
	return result, nil
}
