package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	handler "github.com/paznews/paznews-backend/handlers"
	"github.com/paznews/paznews-backend/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "paz"
	password = "password"
	dbname   = "paznews"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	models.InitDB(psqlInfo)

	handler.HandleRequests()
	log.Fatal(http.ListenAndServe(":8880", nil))
}
