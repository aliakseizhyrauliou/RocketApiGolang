package grpc

import (
	"context"
	"fmt"
	"github.com/aliakseizhyrauliou/gRPCApiGo/internal/db"
	"github.com/aliakseizhyrauliou/gRPCApiGo/internal/rocket"
	rkt "github.com/aliakseizhyrauliou/gRPCApiGo/protos/rocket/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCRocketService struct {
	rkt.UnimplementedRocketServiceServer
	Service rocket.Service
}

func (serv *GRPCRocketService) GetRocket(ctx context.Context, req *rkt.GetRocketRequest) (*rkt.GetRocketResponse, error) {
	r, err := serv.Service.GetRocketByID(ctx, req.GetId())
	if err != nil {
		return &rkt.GetRocketResponse{}, err
	}

	return &rkt.GetRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   r.ID,
			Name: r.Name,
			Type: r.Type,
		},
	}, nil
}

func (serv *GRPCRocketService) AddRocket(ctx context.Context, req *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {
	r := req.GetRocket()

	if (rocket.Service{}) == serv.Service {
		log.Println("Failed to initialize rocketService")
	}

	savingRocket, err := serv.Service.InsertRocket(ctx, rocket.Rocket{
		ID:   r.Id,
		Name: r.Name,
		Type: r.Type,
	})

	if err != nil {
		return &rkt.AddRocketResponse{}, err
	}

	// Assuming savingRocket is a pointer to Rocket, dereference it
	return &rkt.AddRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   savingRocket.ID,
			Name: savingRocket.Name,
			Type: savingRocket.Type,
		},
	}, nil
}

func (serv *GRPCRocketService) DeleteRocket(ctx context.Context, req *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	r := req.GetRocket()

	err := serv.Service.DeleteRocket(ctx, r.Id)

	if err != nil {
		return &rkt.DeleteRocketResponse{Status: "error"}, err
	}

	return &rkt.DeleteRocketResponse{Status: "ok"}, nil
}

// StartGRPCServer starts Ports GRPC server
func StartGRPCServer(db db.Store) error {
	grpcServer := grpc.NewServer()
	rocketService := rocket.New(db)

	if (rocket.Service{}) == rocketService {
		return fmt.Errorf("Failed to initialize rocketService")
	}

	grpcService := GRPCRocketService{
		Service: rocketService,
	}

	rkt.RegisterRocketServiceServer(grpcServer, &grpcService)

	con, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	err = grpcServer.Serve(con)

	return nil
}
