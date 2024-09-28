package repository

import (
	"encoding/json"
	"fmt"
	clima "goexpert-temperature-system-by-cep/internal/domain"
	"net/http"
	"net/url"
)

//go:generate mockery --name ClimaRepository --outpkg mock --output mock --filename clima.go --with-expecter=true

type ClimaRepository interface {
	BuscaClimaPorLocalizacao(location string) (*clima.Clima, error)
}

type climaRepository struct {
	client *http.Client
	url    string
	apiKey string
}

func NewWeatherRepository(client *http.Client, url, apiKey string) ClimaRepository {
	return &climaRepository{
		client: client,
		url:    url,
		apiKey: apiKey,
	}
}

func (r *climaRepository) BuscaClimaPorLocalizacao(location string) (*clima.Clima, error) {
	escapedLocation := url.QueryEscape(location)
	url := fmt.Sprintf("%s?key=%s&q=%s", r.url, r.apiKey, escapedLocation)
	resp, err := r.client.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var weatherResponse clima.ClimaResponse
		if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
			return nil, err
		}
		return &weatherResponse.Current, nil
	}

	return nil, fmt.Errorf("could not fetch clima data")
}
