package main

import (
	"log"
	"os"
	"strconv"

	"github.com/AfandyW/financial-records/controllers"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var server = controllers.Server{}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	server.Initialize(port, user, password, host, dbname)
	server.Run(":3000")
}
