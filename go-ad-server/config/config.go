package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	MongoURI string
	DBName   string
	Port     string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using defaults")
	}

	return &Config{
		MongoURI: getEnv("MONGO_URI", "mongodb://root:examplepassword@localhost:27017"),
		DBName:   getEnv("DB_NAME", "ad_server"),
		Port:     getEnv("PORT", "8080"),
	}
}

func ConnectDB(cfg *Config) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client.Database(cfg.DBName)
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
