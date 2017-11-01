package action

import (
	"golang.org/x/net/context"

	"github.com/argonlaser/bingo/gameShip/gameShipRpc"
	"go.uber.org/zap"
)

// GameStart starts the game
func (s *GameShipRPCServer) GameStart(ctx context.Context, in *gameShipRpc.GameStartRequest) (*gameShipRpc.GameStartResponse, error) {
	var funcName = zap.String("FuncName", "actions|GameStart|")
	s.Logger.Info("funcEvent", funcName, zap.String("Event", "Start"))
	s.Logger.Debug("funcVariable", funcName, zap.Any("GameStartRequest", in))

	// game presence check
	reqGame, err := s.GameStore.GetByGameID(in.GameId)
	if err != nil {
		s.Logger.Error("funcError", funcName, zap.Error(err))
		return nil, err
	}

	reqGame.Start()

	var result = &gameShipRpc.GameStartResponse{
		GameId:   reqGame.ID,
		PlayerId: reqGame.CurrentPlayer.ID,
	}

	s.Logger.Debug("funcVariable", funcName, zap.Any("GameStartResponse", result))
	s.Logger.Info("funcEvent", funcName, zap.String("Event", "End"))
	return result, nil
}
