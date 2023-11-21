package grpc

import (
	"context"
	"github.com/aliakseizhyrauliou/gRPCApiGo/internal/db"
	"github.com/aliakseizhyrauliou/gRPCApiGo/internal/rocket"
	rkt "github.com/aliakseizhyrauliou/gRPCApiGo/protos/rocket/v1"
	"google.golang.org/grpc"
	"net"
)

type GRPCRocketService struct {
	rkt.UnimplementedRocketServiceServer
	rocket.Service
}

func (serv *GRPCRocketService) GetRocket(ctx context.Context, req *rkt.GetRocketRequest) (*rkt.GetRocketResponse, error) {

	rocket, err := serv.GetRocketByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	// Преобразование вашего внутреннего типа Rocket в протобуф-тип
	protoRocket := &rkt.Rocket{
		Id:   rocket.ID,
		Name: rocket.Name,
		Type: rocket.Type,
	}

	// Создание ответа
	response := &rkt.GetRocketResponse{
		Rocket: protoRocket,
	}

	return response, nil

}

func (srv *GRPCRocketService) AddRocket(context.Context, *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {
	return &rkt.AddRocketResponse{}, nil
}

func (srv *GRPCRocketService) DeleteRocket(context.Context, *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	return &rkt.DeleteRocketResponse{}, nil
}

// StartGRPCServer starts Ports GRPC server
func StartGRPCServer(store db.Store) error {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	rktService := NewGRPCRocketService(store)
	rkt.RegisterRocketServiceServer(grpcServer, rktService)

	if err := grpcServer.Serve(listener); err != nil {
		return err
	}

	return nil
}

func NewGRPCRocketService(store db.Store) *GRPCRocketService {

	service := rocket.New(store)
	return &GRPCRocketService{Service: service}
}
