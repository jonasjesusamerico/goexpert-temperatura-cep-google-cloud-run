package service

import (
	"goexpert-temperature-system-by-cep/internal/domain"
	"goexpert-temperature-system-by-cep/internal/infra/repository"
)

//go:generate mockery --name CepService --outpkg mock --output mock --filename zipcode.go --with-expecter=true

type CepService interface {
	BuscaLocalizacaoPorCep(zipCode string) (*domain.Localizacao, error)
}

type cepService struct {
	repository repository.CepRepository
}

func NewZipCodeService(repo repository.CepRepository) CepService {
	return &cepService{
		repository: repo,
	}
}

func (s *cepService) BuscaLocalizacaoPorCep(zipCode string) (*domain.Localizacao, error) {
	return s.repository.BuscaLocalizacaoPorCep(zipCode)
}
