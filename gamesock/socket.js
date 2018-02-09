const path = require('path')
const grpc = require('grpc')
const logger = require('./Logger/winston.js')

const PROTO_PATH = path.join(__dirname, '/../gameShip/gameShipRpc/gameShipRpc.proto')
const proto = grpc.load(PROTO_PATH)
const gameShipRpc = proto.gameShipRpc

const client = new gameShipRpc.gameShipRpc('localhost:50051',
                                       grpc.credentials.createInsecure())

logger.silly(gameShipRpc)

let socket = function (socket) {
  let user = new Player(socket)

  // handle all connected related events.
  logger.info('Socket connected = ', socket.id)

  socket.on('user.strike', function (payload) {
    logger.debug(payload)
	  // emit to all people in room and give chance to next person
    payload.currPlayerIndex = game.increaseCurrentPlayerIndex()
    game.strikeAll(payload.striked)
    game.sendEvent('striked', payload)
  })

  socket.on('game.join', function (payload) {
    logger.warn(payload.board)
    logger.debug('game.join event :', socket.id)
    user.gameBoard.fillCheckbox(payload.board)
    game = global.gamePool.AllocateGameForUser(user)
  })

  socket.on('user.bingo', function (payload) {
    user.bingo(payload)
  })

    // handle all disconnected related events
  socket.on('disconnect', function () {
    logger.info('Disconnected: ' + socket.id)

	  // remove player from the game
    game && game.removePlayer(user)

    if (game && game.players.length === 0) {
	    global.gamePool.RemoveGame(game)
    }
  })
}

module.exports = socket

var gameId

client.newGame({ creatorId: 'player-1'}, function (err, res) {
  console.log('*********************')
  console.log('New game created')
  console.log(err)
  console.log(res)
  gameId = res.gameId
  client.playerJoin({ playerId: 'player-2', gameId: gameId}, function (err, res) {
    console.log('*********************')
    console.log('Player 2 joined')
    console.log(err)
    console.log(res)
    client.gameStart({ gameId: gameId}, function (err, res) {
      console.log('*********************')
      console.log('gameStarted')
      console.log(err)
      console.log(res)
      client.playerStrikeBox({ gameId: gameId, playerId: 'player-1', row: 1, column: 1}, function (err, res) {
        console.log('*********************')
        console.log('player-1 striked r1 c1')
        console.log(err)
        console.log(res)
        client.playerStrikeBox({ gameId: gameId, playerId: 'player-1', row: 1, column: 1}, function (err, res) {
          console.log('*********************')
          console.log('player-1 trying to strike again r1 c1 => should fail')
          console.log(err)
          console.log(res)
          client.playerStrikeBox({ gameId: gameId, playerId: 'player-2', row: 2, column: 2}, function (err, res) {
            console.log('*********************')
            console.log('player-2 strikes r2 c2 ')
            console.log(err)
            console.log(res)
            client.playerBingo({ gameId: gameId, playerId: 'player-2'}, function (err, res) {
              console.log('*********************')
              console.log('player-2 bingoes')
              console.log(err)
              console.log(res)
            })
          })
        })
      })
    })
  })
})
