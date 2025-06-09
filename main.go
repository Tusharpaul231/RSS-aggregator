package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
	"D:\My-Projects\RSS-aggregator\internal\database\database.go"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load()

	// Load environment variables
	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("Port not found")
	}

	// Initialize the database connection
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("Database URL not found")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to the database:", err)
	}


	apiCfg := &apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link", "X-Total-Count"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)
	v1Router.Post("/user", apiCfg.handlerUserCreate)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portStr,
	}

	fmt.Println("Server started at Port :", portStr)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
