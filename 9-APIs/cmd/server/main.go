package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/henrymoreirasilva/go-expert/9-APIs/configs"
	"github.com/henrymoreirasilva/go-expert/9-APIs/internal/entity"
	"github.com/henrymoreirasilva/go-expert/9-APIs/internal/infra/database"
	"github.com/henrymoreirasilva/go-expert/9-APIs/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/producs", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)

	http.ListenAndServe(":8080", r)

}

// 172.25.240.1

// netsh interface portproxy add v4tov4 listenport=8080 listenaddress=0.0.0.0 connectport=8080 connectaddress=172.25.240.1
