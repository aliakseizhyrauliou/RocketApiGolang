package main

import (
	"github.com/aliakseizhyrauliou/gRPCApiGo/internal/db"
	"github.com/aliakseizhyrauliou/gRPCApiGo/transport/grpc"
	"log"
)

func Run() error {
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
