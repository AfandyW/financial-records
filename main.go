package main

import (
	"github.com/AfandyW/financial-records/controllers"
	_ "github.com/lib/pq"
)

//note
//buat penerapan singleton pada database
//refactoring code ke layar layar , repository, service, handler(controller)
//jika ada error, taruhkan return, agar code tidak lanjut kebawa
// buat pagination, pada penarikan data get all . misalnya /page=2, pagination, limit
// rangkuman buku clean code,  harus bisa menjawab interviwer (apa itu clean code)

var server = controllers.Server{}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "financial"
)

func main() {
	server.Initialize(port, user, password, host, dbname)
	server.Run(":3000")
}
