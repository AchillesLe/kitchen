package main

import (
	"log"
	"net/http"

	handler "github.com/AchillesLe/kitchen/services/orders/handler/orders"
	"github.com/AchillesLe/kitchen/services/orders/service"
)

type HttpServer struct {
	addr string
}

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{addr: addr}
}

func (h *HttpServer) Run() error {
	router := http.NewServeMux()

	ordersHttpHandler := handler.NewOrdersHttpHandler(service.NewOrderService())
	ordersHttpHandler.RegisterRoute(router)

	log.Println("Starting http server on ", h.addr)
	return http.ListenAndServe(h.addr, router)
}
