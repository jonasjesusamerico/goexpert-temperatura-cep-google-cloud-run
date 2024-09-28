package main

import (
	"goexpert-temperature-system-by-cep/internal/application/service"
	"goexpert-temperature-system-by-cep/internal/infra/repository"
	entrypoint "goexpert-temperature-system-by-cep/internal/infra/web"
	"goexpert-temperature-system-by-cep/internal/infra/web/controller"
	"net/http"
	"time"
)

type Configuracao struct {
	ViaCepURL     string
	WeatherAPIURL string
	WeatherAPIKey string
}

func NovaConfiguracao() *Configuracao {
	return &Configuracao{
		ViaCepURL:     "https://viacep.com.br/ws",
		WeatherAPIURL: "http://api.weatherapi.com/v1/current.json",
		WeatherAPIKey: "e5bd00e528e346ff8a840254213009",
	}
}

func NewHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
	}
}

func main() {
	conf := NovaConfiguracao()
	httpClient := NewHTTPClient()

	zipCodeRepo := repository.NewZipCodeRepository(httpClient, conf.ViaCepURL)
	weatherRepo := repository.NewWeatherRepository(httpClient, conf.WeatherAPIURL, conf.WeatherAPIKey)

	cepService := service.NewZipCodeService(zipCodeRepo)
	climaService := service.NewClimaService(weatherRepo)

	weatherController := controller.NewClimaController(climaService, cepService)

	r := entrypoint.SetupRouter(weatherController)
	r.Run(":8080")
}
