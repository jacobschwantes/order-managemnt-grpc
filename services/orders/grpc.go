package main

import (
	"log"
	"net"

	handler "github.com/jacobschwantes/order-managemnt-grpc/services/orders/handler/orders"
	"github.com/jacobschwantes/order-managemnt-grpc/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewgRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// register our grpc services

	orderService := service.NewOrderService()
	handler.NewOrdersGrpcHandler(grpcServer, orderService)

	log.Printf("orders service listening on %s", s.addr)

	return grpcServer.Serve(lis)
}
