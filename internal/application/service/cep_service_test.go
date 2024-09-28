package service

import (
	"errors"
	"testing"

	"goexpert-temperature-system-by-cep/internal/domain"
	mock_repository "goexpert-temperature-system-by-cep/internal/infra/repository/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CepServiceSuite struct {
	suite.Suite
	mockRepo *mock_repository.CepRepository
	service  CepService
}

func (suite *CepServiceSuite) SetupSuite() {
	suite.mockRepo = mock_repository.NewZipCodeRepository(suite.T())
	suite.service = NewZipCodeService(suite.mockRepo)
}

func (suite *CepServiceSuite) TestGetLocationByZipCode() {
	suite.T().Run("Success", func(t *testing.T) {
		expectedLocation := &domain.Localizacao{
			Cep:        "12345678",
			Logradouro: "Rua Exemplo",
			Bairro:     "Bairro Exemplo",
			Localidade: "Cidade Exemplo",
			Uf:         "UF",
		}

		expectedZipCode := "10001"
		suite.mockRepo.On("BuscaLocalizacaoPorCep", expectedZipCode).Return(expectedLocation, nil)

		result, err := suite.service.BuscaLocalizacaoPorCep(expectedZipCode)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedLocation, result)

		suite.mockRepo.AssertExpectations(t)
	})

	suite.T().Run("Error", func(t *testing.T) {
		expectedError := errors.New("repository error")
		expectedZipCode := "90210"
		suite.mockRepo.On("BuscaLocalizacaoPorCep", expectedZipCode).Return(nil, expectedError)

		result, err := suite.service.BuscaLocalizacaoPorCep(expectedZipCode)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, expectedError.Error())

		suite.mockRepo.AssertExpectations(t)
	})
}

func (suite *CepServiceSuite) TearDownSuite() {
	suite.mockRepo = nil
	suite.service = nil
}

func TestZipCodeServiceSuite(t *testing.T) {
	suite.Run(t, new(CepServiceSuite))
}
