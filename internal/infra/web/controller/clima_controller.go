package controller

import (
	"goexpert-temperature-system-by-cep/internal/application/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClimaController struct {
	climaService service.ClimaService
	cepService   service.CepService
}

func NewClimaController(climaService service.ClimaService, cepService service.CepService) *ClimaController {
	return &ClimaController{
		climaService: climaService,
		cepService:   cepService,
	}
}

func (h *ClimaController) BuscaClima(c *gin.Context) {
	zipCode := c.Param("cep")

	if len(zipCode) != 8 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid zipcode"})
		return
	}

	location, _ := h.cepService.BuscaLocalizacaoPorCep(zipCode)
	if location.Localidade == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "can not find zipcode"})
		return
	}

	clima, err := h.climaService.BuscaClimaPorLocalizacao(location.Localidade)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch clima data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"temp_C": clima.TemperaturaCelsius,
		"temp_F": clima.TemperaturaCelsius*1.8 + 32,
		"temp_K": clima.TemperaturaCelsius + 273.15,
	})
}
