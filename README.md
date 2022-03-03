# game-stats-api

A simple REST API server in Go using:
* SQLite Database
* gorilla/mux router
* swaggo/swag API documentation

Can serve a SPA UI frontend in addition:
* [game-stats-vue](https://github.com/yannickbilcot/game-stats-vue)


## Installation and Run

#### Download the project

```bash
git clone https://github.com/yannickbilcot/game-stats-api.git
cd game-stats-api
```

#### To serve the UI frontend (optional)
* UI demo:  https://game-stats-demo.herokuapp.com/

```bash
# fetch the ui submodule
git submodule update --init --recursive
# build the ui
cd ui
npm install
quasar build
```

#### To generate the swagger documentation (optional)
* API documentation:  https://game-stats-demo.herokuapp.com/swagger/

```bash
go get -u github.com/swaggo/swag/cmd/swag
swag init
```

#### Build and start the server
```bash
go build
./game-stats-api
```
