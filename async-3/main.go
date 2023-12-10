package main

import (
	"async-3/server"
	"async-3/server/config"
	"net/http"
)

func main() {
	run()
}

func run() {
	router := http.NewServeMux()

	port := ":9999"

	db := config.CreateConnection()

	server.StartServer(router, port, db)
}
