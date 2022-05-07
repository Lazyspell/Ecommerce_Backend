package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/lazyspell/Ecommerce_Backend/config"
	"github.com/lazyspell/Ecommerce_Backend/handlers"
)

func Routes(app *config.AppConfig) http.Handler {

	config.LoadConfig()
	mux := chi.NewRouter()
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	mux.Get("/google_login", handlers.Repo.GoogleLogin)
	mux.Get("/google_callback", handlers.Repo.GoogleCallback)
	mux.Get("/categories", handlers.Repo.GetAllCategories)
	mux.Get("/categories/id", handlers.Repo.GetCategoryById)
	mux.Get("/users/all", handlers.Repo.GetAllUsers)
	mux.Get("/users/id", handlers.Repo.GetUserById)

	mux.Post("/users/new", handlers.Repo.NewUser)

	mux.Delete("/users/delete", handlers.Repo.DeleteUser)

	return mux
}
