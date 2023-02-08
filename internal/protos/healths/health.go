package healths

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"Server/internal/protos"
	"Server/internal/services"
)

type Serv struct {
	protos.UnimplementedSmallHealthServiceServer
	s services.IService
}

func NewServ(s services.IService) *Serv {
	return &Serv{s: s}
}

func (serv *Serv) Check(ctx context.Context, in *emptypb.Empty) (*protos.Health, error) {
	var response protos.Health
	response.Mongo = serv.s.TouchMongo()
	response.Postgres = serv.s.TouchPostgres()
	return &response, nil
}
