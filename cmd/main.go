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
		log.Fatalf("settings creation error: %s", err)
	}
	repository, err := repo.NewRepository(sets)
	if err != nil {
		log.Fatalf("repository creation error: %s", err)
	}
	service := services.NewService(repository)
	serv := healths.NewServ(service)
	listener, err := net.Listen("tcp", ":9000")
		log.Fatalf("did not connect: %s", err)
	if err != nil {
	}

	s := grpc.NewServer()
	protos.RegisterSmallHealthServiceServer(s, serv)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
