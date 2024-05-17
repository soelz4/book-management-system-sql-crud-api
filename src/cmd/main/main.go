package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"

	"go-bookstore/src/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting Server at PORT 9010")
	err := http.ListenAndServe(":9010", r)
	if err != nil {
		log.Fatal(err)
	}
}
