package main

import (
	"log"
	"net"

	handler "github.com/AchillesLe/kitchen/services/orders/handler/orders"
	"github.com/AchillesLe/kitchen/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	listen, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatal("Failed to listen : ", err)
	}

	grpcSV := grpc.NewServer()
	// ------------------- register grpc services -------------------
	OrderService := service.NewOrderService()
	handler.NewGrpcOrderService(grpcSV, OrderService)

	log.Println("Starting gRPC server on ", s.addr)

	return grpcSV.Serve(listen)
}
