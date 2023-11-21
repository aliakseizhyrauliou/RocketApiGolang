//go:generate mockgen -source=rocket.go -destination=rocket_mocks_test.go -package=rocket github.com/aliakseizhyrauliou/gRPCApiGo/internal/rocket

package rocket

import (
	"context"
)

type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights string
}

type Store interface {
	GetRocketByID(ctx context.Context, id string) (Rocket, error)
	InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error)
	DeleteRocket(ctx context.Context, id string) error
	GetRocketList(ctx context.Context, take int32, skip int32) ([]Rocket, error)
}

// Responsible for rocket
type Service struct {
	Store Store
}

func (service Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rocket, err := service.Store.GetRocketByID(ctx, id)
	if err != nil {
		return Rocket{}, err
	}

	return rocket, nil
}

func (service Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rocket, err := service.Store.InsertRocket(ctx, rkt)
	if err != nil {
		return Rocket{}, err
	}
	return rocket, nil
}

func (service Service) DeleteRocket(ctx context.Context, id string) error {
	err := service.Store.DeleteRocket(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (service Service) GetRocketList(ctx context.Context, take int32, skip int32) ([]Rocket, error) {
	rockets, err := service.Store.GetRocketList(ctx, take, skip)

	if err != nil {
		return []Rocket{}, nil
	}

	return rockets, nil
}

// Return new service
func New(store Store) Service {
	return Service{
		Store: store,
	}
}
