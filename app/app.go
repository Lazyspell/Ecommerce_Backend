package app

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lazyspell/Ecommerce_Backend/config"
	"github.com/lazyspell/Ecommerce_Backend/driver"
	"github.com/lazyspell/Ecommerce_Backend/handlers"
	"github.com/lazyspell/Ecommerce_Backend/helpers"
	r "github.com/lazyspell/Ecommerce_Backend/routes"
)

var app config.AppConfig

const portNumber = ":8080"

func Start() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()

	fmt.Println(fmt.Sprintf("Starting application on http://localhost:8080"))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: r.Routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {

	//read flags
	// inProduction := flag.Bool("production", true, "Application is in production")
	flag.StringVar(&app.DB.DbHost, "host", os.Getenv("DIGITAL_OCEAN_DB1_HOST"), "Database host")
	flag.StringVar(&app.DB.DbName, "name", os.Getenv("DIGITAL_OCEAN_DB1_DATABASE"), "Database name")
	flag.StringVar(&app.DB.DbUser, "user", os.Getenv("DIGITAL_OCEAN_DB1_USER"), "Database user")
	flag.StringVar(&app.DB.DbPass, "password", os.Getenv("DIGITAL_OCEAN_DB1_PASS"), "Database password")
	flag.StringVar(&app.DB.DbPort, "port", "25060", "Databae port")
	flag.StringVar(&app.DB.DbSSL, "ssl", "require", "Database ssl settings (disable, prefer, require)")

	flag.Parse()

	if app.DB.DbName == "" || app.DB.DbUser == "" {
		fmt.Println("Missing require flags")
		os.Exit(1)
	}

	log.Println("Connecting to database...")
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", app.DB.DbHost, app.DB.DbPort, app.DB.DbName, app.DB.DbUser, app.DB.DbPass, app.DB.DbSSL)

	db, err := driver.ConnectSQL(connectionString)
	if err != nil {
		log.Fatal("Can not connect to database! Dying...")
	}
	log.Println("Connected to database!")

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandler(repo)
	helpers.NewHelpers(&app)

	return db, nil

}
