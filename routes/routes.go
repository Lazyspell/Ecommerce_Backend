package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/lazyspell/Ecommerce_Backend/config"
	"github.com/lazyspell/Ecommerce_Backend/handlers"
)

var tokenAuth *jwtauth.JWTAuth
var password = "Elaine I will love you forever and always"

func Routes(app *config.AppConfig) http.Handler {

	config.LoadConfig()
	mux := chi.NewRouter()
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	mux.Group(func(r chi.Router) {
		tokenAuth = jwtauth.New("HS256", []byte(password), nil)
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/users/all", handlers.Repo.GetAllUsers)
	})

	mux.Get("/google_login", handlers.Repo.GoogleLogin)
	mux.Get("/google_callback", handlers.Repo.GoogleCallback)
	mux.Get("/categories", handlers.Repo.GetAllCategories)
	mux.Get("/categories/id", handlers.Repo.GetCategoryById)

	mux.Get("/users/id", handlers.Repo.GetUserById)

	mux.Post("/users/new", handlers.Repo.NewUser)
	mux.Post("/users/login", handlers.Repo.LoginUser)

	mux.Delete("/users/delete", handlers.Repo.DeleteUser)

	return mux
}
