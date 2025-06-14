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
	"github.com/Tusharpaul231/RSS-aggregator/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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



	// Initialize the database connection using individual environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatal("One or more required DB environment variables are missing")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	if err = conn.Ping(); err != nil {
		log.Fatal("Failed to ping the database:", err)
	}
	fmt.Println("✅ Connected to the database!")



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
	v1Router.Get("/user", apiCfg.handlerGetUser)

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
