package main

import (
	"log"
	"net/http"

	handler "github.com/jacobschwantes/order-managemnt-grpc/services/orders/handler/orders"
	"github.com/jacobschwantes/order-managemnt-grpc/services/orders/service"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	orderService := service.NewOrderService()
	orderHandler := handler.NewOrdersHttpHandler(orderService)
	orderHandler.RegisterRouter(router)

	log.Printf("orders service listening on %s", s.addr)

	return http.ListenAndServe(s.addr, router)
}
