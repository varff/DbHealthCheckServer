package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"Server/internal/protos"
	"Server/internal/protos/healths"
	"Server/internal/repo"
	"Server/internal/services"
	"Server/internal/settings"
)

func main() {
	sets, err := settings.NewDBSetting()
	if err != nil {
		return
	}
	repository, err := repo.NewRepository(sets)
	if err != nil {
		return
	}
	service := services.NewService(repository)
	serv := healths.NewServ(service)
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	protos.RegisterSmallHealthServiceServer(s, serv)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
