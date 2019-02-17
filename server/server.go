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

const defaultPort = "7777"

var db *gorm.DB

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db = database_config.DbConn()
	defer db.Close()
	db.AutoMigrate(&capBook.User{}, &capBook.Author{}, &capBook.Book{}, &capBook.Location{}, &capBook.Publisher{}, &capBook.Rental{})
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", corsAccess(handler.GraphQL(capBook.NewExecutableSchema(capBook.Config{Resolvers: &capBook.Resolver{}}))))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func corsAccess(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if req.Method == "OPTIONS" {
			return
		}
		next(w, req)
	})
}
