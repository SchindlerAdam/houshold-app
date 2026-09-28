package server

import (
	"context"
	"database/sql"
	"fmt"
	"houshold-app/internal/handler"
	"houshold-app/internal/repository"
	"houshold-app/internal/service"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func StartApp() {
	loadEnvValues()
	router := gin.Default()
	db := createDB()
	initializeApplication(router, db)
	router.Run("localhost:8080")
}

func loadEnvValues() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Cannot find .env file: %v", err)
	}
}

func initializeApplication(router *gin.Engine, db *sql.DB) {
	registerUserRoutes(router, db)
}

func registerUserRoutes(router *gin.Engine, db *sql.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router.POST("/user", userHandler.CreateUser)
}

func createDB() *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSL := os.Getenv("DB_SSLMODE")
	schema := os.Getenv("SCHEMA")

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&search_path=%s",
		dbUser, dbPass, dbHost, dbPort, dbName, dbSSL, schema,
	)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Cannot open the DB connection: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("Cannot connect to the DB: %v", err)
	}
	log.Println("Connected to the database successfully!")
	return db
}
