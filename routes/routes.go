package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/lazyspell/Ecommerce_Backend/config"
	"github.com/lazyspell/Ecommerce_Backend/handlers"
)

var password = "Elaine I will love you forever and always"

func Routes(app *config.AppConfig) http.Handler {

	config.LoadConfig()
	mux := chi.NewRouter()
	mux.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"http://localhost:3000"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	mux.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwtauth.New("HS256", []byte(password), nil)))
		r.Use(jwtauth.Authenticator)
		r.Get("/users/all", handlers.Repo.GetAllUsers)
		r.Delete("/users/delete", handlers.Repo.DeleteUser)
	})

	mux.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwtauth.New("HS256", []byte(password), nil)))
		r.Use(jwtauth.Authenticator)
	})

	mux.Post("/login", handlers.Repo.LoginUser)

	mux.Post("/google_login", handlers.Repo.GoogleUserLogin)
	mux.Get("/google_callback", handlers.Repo.GoogleCallback)
	mux.Get("/categories", handlers.Repo.GetAllCategories)
	mux.Get("/categories/id", handlers.Repo.GetCategoryById)

	mux.Get("/users/id", handlers.Repo.GetUserById)
	mux.Post("/users/new", handlers.Repo.NewUser)
	mux.Get("/users/all", handlers.Repo.GetAllUsers)

	return mux
}
