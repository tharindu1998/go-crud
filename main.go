package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/tharindu1998/go-crud/internal/database"
)

type apiConfig struct{
	DB * database.Queries
}

func main() {
	fmt.Println("Hello, World!")

	godotenv.Load(".env")

	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("PORT environment variable is not set")
	}
	fmt.Printf("Server will run on port: %s\n", portStr)

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}
	fmt.Printf("Server will run on port: %s\n", dbURL)

	conn, err1 := sql.Open("postgres",dbURL)
	if err1 != nil {
		log.Fatal("Cannot connect to database",err1)
	}

	queries := database.New(conn)

	apiConfig := apiConfig{
		DB: queries,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of the browsers
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerErr)
	
	v1Router.Post("/users", apiConfig.handlerCreateUser)

	router.Mount("/v1", v1Router)

	srv:= &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", portStr),
	}
	
	err:= srv.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		log.Fatal(err)
	} else {
		fmt.Println("Server started successfully")
	}

	
	
}