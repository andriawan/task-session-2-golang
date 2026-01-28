package app

import (
	"category-crud/config"
	"category-crud/db"
	"category-crud/route"
	"category-crud/store"
	"log"
	"net/http"
)

// @title Category CRUD API
// @version 1.0
// @description API for managing categories with CRUD operations
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

func Start() {
	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	db, err := db.Configure(*config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	store := store.NewCategoryStore()
	r := route.Configure(store)

	port := config.Server.Port
	log.Println("Server starting on :" + port)
	log.Println("Swagger documentation available at http://localhost:8080/swagger/index.html")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
