package service

import (
	"errors"
	"testing"

	"goexpert-temperature-system-by-cep/internal/domain"
	mock_repositoy "goexpert-temperature-system-by-cep/internal/infra/repository/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ClicaServiceSuite struct {
	suite.Suite
	mockRepo *mock_repositoy.ClimaRepository
	service  ClimaService
}

func (suite *ClicaServiceSuite) SetupSuite() {
	suite.mockRepo = mock_repositoy.NewWeatherRepository(suite.T())
	suite.service = NewClimaService(suite.mockRepo)
}

func (suite *ClicaServiceSuite) TestGetWeatherByLocation() {
	suite.T().Run("Success", func(t *testing.T) {
		expectedWeather := &domain.Clima{TemperaturaCelsius: 25.0}
		expectedLocation := "New York"
		suite.mockRepo.On("BuscaClimaPorLocalizacao", expectedLocation).Return(expectedWeather, nil)

		result, err := suite.service.BuscaClimaPorLocalizacao(expectedLocation)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedWeather, result)

		suite.mockRepo.AssertExpectations(t)
	})

	suite.T().Run("Error", func(t *testing.T) {
		expectedError := errors.New("repository error")
		expectedLocation := "London"
		suite.mockRepo.On("BuscaClimaPorLocalizacao", expectedLocation).Return(nil, expectedError)

		result, err := suite.service.BuscaClimaPorLocalizacao(expectedLocation)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, expectedError.Error())

		suite.mockRepo.AssertExpectations(t)
	})
}

func (suite *ClicaServiceSuite) TearDownSuite() {
	suite.mockRepo = nil
	suite.service = nil
}

func TestWeatherServiceSuite(t *testing.T) {
	suite.Run(t, new(ClicaServiceSuite))
}
