package main

import (
	"sportsync/internal/config"
	"sportsync/internal/database"
	"sportsync/internal/server"
)

func main() {
	env := config.LoadEnv()
	db := database.ConnectDB(env)
	server.StartServer(db, env)

}
