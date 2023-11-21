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

	databaseRocket, err := serv.Service.GetRocketByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	protoRocket := &rkt.Rocket{
		Id:   databaseRocket.ID,
		Name: databaseRocket.Name,
		Type: databaseRocket.Type,
	}

	response := &rkt.GetRocketResponse{
		Rocket: protoRocket,
	}

	return response, nil

}

func (srv *GRPCRocketService) AddRocket(ctx context.Context, req *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {
	protoRocket := req.GetRocket()

	databaseRocket, err := srv.Service.InsertRocket(ctx, rocket.Rocket{
		Name: protoRocket.Name,
		Type: protoRocket.Type,
	})

	if err != nil {
		return &rkt.AddRocketResponse{}, err
	}

	return &rkt.AddRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   databaseRocket.ID,
			Name: databaseRocket.Name,
			Type: databaseRocket.Type,
		},
	}, nil
}

func (srv *GRPCRocketService) DeleteRocket(ctx context.Context, req *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	protoRocket := req.GetRocket()

	err := srv.Service.DeleteRocket(ctx, protoRocket.Id)

	if err != nil {
		return &rkt.DeleteRocketResponse{
			Status: "ERROR",
		}, err
	}

	return &rkt.DeleteRocketResponse{
		Status: "OK",
	}, nil
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
