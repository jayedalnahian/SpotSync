package main

import (
	"SpotSync/internal/config"
	"SpotSync/internal/server"
	
)



func main() {
	cfg := config.LoadEnv()
	db := config.ConnectDatabase(cfg)

	

	server.Start( db, cfg)
}
