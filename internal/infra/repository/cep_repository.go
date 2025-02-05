package repository

import (
	"encoding/json"
	"fmt"
	zipcode "goexpert-temperature-system-by-cep/internal/domain"
	"net/http"
)

//go:generate mockery --name CepRepository --outpkg mock --output mock --filename zipcode.go --with-expecter=true

type CepRepository interface {
	BuscaLocalizacaoPorCep(zipCode string) (*zipcode.Localizacao, error)
}

type zipCodeRepository struct {
	client *http.Client
	url    string
}

func NewZipCodeRepository(client *http.Client, url string) CepRepository {
	return &zipCodeRepository{
		client: client,
		url:    url,
	}
}

func (r *zipCodeRepository) BuscaLocalizacaoPorCep(zipCode string) (*zipcode.Localizacao, error) {
	url := fmt.Sprintf("%s/%s/json/", r.url, zipCode)

	resp, err := r.client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var location zipcode.Localizacao
		if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
			return nil, err
		}
		return &location, nil
	}

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("can not find zipcode")
	}

	return nil, fmt.Errorf("invalid zipcode")
}
