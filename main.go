package main

import (
	"github.com/AfandyW/financial-records/controllers"
	_ "github.com/lib/pq"
)

var server = controllers.Server{}

const (
	host     = "localhost"
	port     = 5432
	user     = "mapple"
	password = "afandy"
	dbname   = "financial"
)

func main() {
	server.Initialize(port, user, password, host, dbname)
	server.Run(":3000")
}
