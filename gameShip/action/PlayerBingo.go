package action

import (
	"golang.org/x/net/context"

	"github.com/argonlaser/bingo/gameShip/gameShipRpc"
	"go.uber.org/zap"
)

// PlayerBingo creates a Game object and saves in a global store
func (s *GameShipRPCServer) PlayerBingo(ctx context.Context, in *gameShipRpc.PlayerBingoRequest) (*gameShipRpc.PlayerBingoResponse, error) {
	var funcName = zap.String("FuncName", "actions|PlayerBingo|")
	s.Logger.Info("funcEvent", funcName, zap.String("Event", "Start"))
	s.Logger.Debug("funcVariable", funcName, zap.Any("PlayerBingoRequest", in))

	reqGame, err := s.GameStore.GetByGameID(in.GameId)
	if err != nil {
		s.Logger.Error("funcError", funcName, zap.Error(err))
		return nil, err
	}

	reqPlayer, err := reqGame.GetPlayerByID(in.PlayerId)
	if err != nil {
		s.Logger.Error("funcError", funcName, zap.Error(err))
		return nil, err
	}

	bingoResult, err := reqGame.PlayerBingo(reqPlayer)
	if err != nil {
		s.Logger.Error("funcError", funcName, zap.Error(err))
		return nil, err
	}

	var result = &gameShipRpc.PlayerBingoResponse{
		GameId:   reqGame.ID,
		PlayerId: reqPlayer.ID,
		Position: bingoResult.Position,
	}

	s.Logger.Debug("funcVariable", funcName, zap.Any("PlayerBingoResponse", result))
	s.Logger.Info("funcEvent", funcName, zap.String("Event", "End"))
	return result, nil
}
