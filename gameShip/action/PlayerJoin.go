package action

import (
	"errors"

	"golang.org/x/net/context"

	"github.com/argonlaser/bingo/gameShip/game"
	"github.com/argonlaser/bingo/gameShip/gameShipRpc"
	"go.uber.org/zap"
)

// PlayerJoin allows player to join a game
func (s *GameShipRPCServer) PlayerJoin(ctx context.Context, in *gameShipRpc.PlayerJoinRequest) (*gameShipRpc.PlayerJoinResponse, error) {
	var funcName = zap.String("FuncName", "actions|PlayerJoin|")
	s.Logger.Info("funcEvent", funcName, zap.String("Event", "Start"))
	s.Logger.Debug("funcVariable", funcName, zap.Any("PlayerJoinRequest", in))

	var newPlayer = game.NewPlayer(in.PlayerId)

	// game presence check
	reqGame, err := s.GameStore.GetByGameID(in.GameId)
	if err != nil {
		s.Logger.Error("funcError", funcName, zap.Error(err))
		return nil, err
	}

	// player absence check
	var alreadyJoined bool
	for _, p := range reqGame.Players {
		if p.ID == in.PlayerId {
			alreadyJoined = true
			break
		}
	}
	if alreadyJoined {
		s.Logger.Error("funcError", funcName, zap.Error(err))
		return nil, errors.New("PlayerAlreadyJoined")
	}

	reqGame.Players = append(reqGame.Players, newPlayer)

	var result = &gameShipRpc.PlayerJoinResponse{
		GameId: reqGame.ID,
	}

	s.Logger.Debug("funcVariable", funcName, zap.Any("PlayerJoinResponse", result))
	s.Logger.Info("funcEvent", funcName, zap.String("Event", "End"))
	return result, nil
}
