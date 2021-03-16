const jsonServer = require('json-server')
const server = jsonServer.create()
const router = jsonServer.router('negativacoes.json')
const middlewares = jsonServer.defaults()

server.use(middlewares)
// CORS
server.set('access-control-allow-origin', '*');
server.set('access-control-allow-methods', '*');
server.set('access-control-allow-headers', '*');
// -----
server.use(router)

server.listen(3333, () => {
    console.log('Server listening...')
})