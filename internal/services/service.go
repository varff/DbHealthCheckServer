package services

import "Server/internal/repo"

type IService interface {
	TouchMongo() string
	TouchPostgres() string
}

type Service struct {
	repo repo.IRepository
}

func NewService(repo repo.IRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) TouchMongo() string {
	err := s.repo.PingMongo()
	if err != nil {
		return err.Error()
	}
	return "Ok"
}

func (s *Service) TouchPostgres() string {
	err := s.repo.PingPostgres()
	if err != nil {
		return err.Error()
	}
	return "Ok"
}
