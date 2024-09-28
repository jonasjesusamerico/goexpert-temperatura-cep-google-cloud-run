package service

import (
	clima "goexpert-temperature-system-by-cep/internal/domain"
	"goexpert-temperature-system-by-cep/internal/infra/repository"
)

//go:generate mockery --name ClimaService --outpkg mock --output mock --filename clima.go --with-expecter=true

type ClimaService interface {
	BuscaClimaPorLocalizacao(location string) (*clima.Clima, error)
}

type climaService struct {
	repository repository.ClimaRepository
}

func NewClimaService(repo repository.ClimaRepository) ClimaService {
	return &climaService{
		repository: repo,
	}
}

func (s *climaService) BuscaClimaPorLocalizacao(location string) (*clima.Clima, error) {
	return s.repository.BuscaClimaPorLocalizacao(location)
}
