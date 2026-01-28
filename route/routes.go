package route

import (
	"category-crud/handler"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// SetupRoutes configures all API routes
func Configure(handlerGroup *handler.HandlerGroup) *mux.Router {
	r := mux.NewRouter()

	// Root route - redirect to Swagger
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusMovedPermanently)
	}).Methods("GET")

	// Category endpoints
	r.HandleFunc("/categories", handlerGroup.Category.Create).Methods("POST")
	r.HandleFunc("/categories", handlerGroup.Category.GetAll).Methods("GET")
	r.HandleFunc("/categories/{id}", handlerGroup.Category.GetByID).Methods("GET")
	r.HandleFunc("/categories/{id}", handlerGroup.Category.Update).Methods("PUT")
	r.HandleFunc("/categories/{id}", handlerGroup.Category.Delete).Methods("DELETE")

	// Product endpoints
	r.HandleFunc("/products", handlerGroup.Product.Create).Methods("POST")
	r.HandleFunc("/products", handlerGroup.Product.GetAll).Methods("GET")
	r.HandleFunc("/products/{id}", handlerGroup.Product.GetByID).Methods("GET")
	r.HandleFunc("/products/{id}", handlerGroup.Product.Update).Methods("PUT")
	r.HandleFunc("/products/{id}", handlerGroup.Product.Delete).Methods("DELETE")

	// Swagger documentation
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return r
}
