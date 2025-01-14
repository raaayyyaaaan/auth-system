package main

import (
	"auth-system/internal/database"
	"auth-system/internal/handlers"
	"auth-system/internal/logger"
	"auth-system/internal/middleware"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Initialize logger
	logger := logger.NewLogger()

	// Load .env file
	if err := godotenv.Load(); err != nil {
		logger.Error("Warning: .env file not found")
	}

	// Build connection string from environment variables
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	// Initialize database
	db, err := database.InitDB(connectionString)
	if err != nil {
		logger.Error("Database initialization failed: %v", err)
		return
	}
	defer db.Close()

	// Initialize router
	router := mux.NewRouter()
	authHandler := handlers.NewAuthHandler(db)

	// Public routes
	router.HandleFunc("/register", authHandler.Register).Methods("POST")
	router.HandleFunc("/login", authHandler.Login).Methods("POST")
	router.HandleFunc("/logout", authHandler.Logout).Methods("POST")

	// Protected routes
	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/profile", authHandler.ProtectedResource).Methods("GET")

	// Add basic logging middleware
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("Request: %s %s", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	})

	// Start server
	logger.Info("Server starting on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.Error("Server failed to start: %v", err)
	}
}
