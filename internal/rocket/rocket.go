//go:generate mockgen -source=rocket.go -destination=rocket_mocks_test.go -package=rocket github.com/aliakseizhyrauliou/gRPCApiGo/internal/rocket

package rocket

import (
	"context"
	"log"
)

type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights string
}

type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// Responsible for rocket
type Service struct {
	Store Store
}

func (service Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rocket, err := service.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}

	return rocket, nil
}

func (service Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	log.Println("HEEY3")

	rocket, err := service.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}
	return rocket, nil
}

func (service Service) DeleteRocket(ctx context.Context, id string) error {
	err := service.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}

// Return new service
func New(store Store) Service {
	return Service{
		Store: store,
	}
}
