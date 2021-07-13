package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	DB     *sql.DB
	Router *chi.Mux
}

func (server *Server) Initialize(DBPort int, DBUser, DBPassword, DBHost, DBName string) {
	var err error

	// Open the connection
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DBHost, DBPort, DBUser, DBPassword, DBName)

	server.DB, err = sql.Open("postgres", psqlconn)

	if err != nil {
		fmt.Printf("Cannot connect to %s database", DBName)
		log.Fatal(" Error : ", err)
	} else {
		fmt.Printf("Connected to %s database", DBName)
	}

	server.Router = chi.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(port string) {
	fmt.Printf("Listening to port %s", port)
	http.ListenAndServe(port, server.Router)

	server.DB.Close()
}
