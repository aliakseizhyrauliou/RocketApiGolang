package main

import (
	"fmt"
	"github.com/aliakseizhyrauliou/gRPCApiGo/internal/db"
	"github.com/aliakseizhyrauliou/gRPCApiGo/transport/grpc"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Run() error {
	loadEnvVariables()

	rocketStore, err := db.New()

	if err != nil {
		return err
	}

	if err = rocketStore.Migrate(); err != nil {
		log.Println("error while running migration")
		return err
	}

	err = grpc.StartGRPCServer(rocketStore)
	if err != nil {
		log.Println("error while running migration")
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func loadEnvVariables() {
	if os.Getenv("SERVER_ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Println("Error loading .env")
		}
	}
}
