package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/henrymoreirasilva/go-expert/9-APIs/configs"
	"github.com/henrymoreirasilva/go-expert/9-APIs/internal/entity"
	"github.com/henrymoreirasilva/go-expert/9-APIs/internal/infra/database"
	"github.com/henrymoreirasilva/go-expert/9-APIs/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config, err := configs.LoadConfig("./")
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

	// Evita que o servidor seja interrompido
	r.Use(middleware.Recoverer)

	// Insere chave:valor no contexto da requisição
	r.Use(middleware.WithValue("tokenAuth", config.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", config.JWTExperesIn))

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)
	r.Get("/users/{email}", userHandler.GetUser)

	http.ListenAndServe(":8080", r)

}

// 172.25.240.1

// netsh interface portproxy add v4tov4 listenport=8080 listenaddress=0.0.0.0 connectport=8080 connectaddress=172.25.240.1
