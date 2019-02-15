package main

import (
	"capBook"
	"capBook/database_config"
	"github.com/99designs/gqlgen/handler"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

var db *gorm.DB

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db = database_config.DbConn()
	defer db.Close()
	db.AutoMigrate(&capBook.User{})
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(capBook.NewExecutableSchema(capBook.Config{Resolvers: &capBook.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
