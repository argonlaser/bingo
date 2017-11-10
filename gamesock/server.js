const app = require('http').createServer(handler)
const io = require('socket.io')(app)
const fs = require('fs')
const path = require('path')
const logger = require('./Logger/winston.js')

app.listen(8080)
logger.info('Server is running in port 8080')

function handler (req, res) {
    fs.readFile(path.join(__dirname, '/public/index.html'),
		function (err, data) {
		    if (err) {
			res.writeHead(500)
			logger.error('Error loading index.html')
			return res.end('Error loading index.html')
		    }

		    logger.info('Served index.html')
		    res.writeHead(200)
		    res.end(data)
		})
}

io.on('connection', require('./socket.js'))
