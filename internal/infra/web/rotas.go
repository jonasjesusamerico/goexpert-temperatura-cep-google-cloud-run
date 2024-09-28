package entrypoint

import (
	"goexpert-temperature-system-by-cep/internal/infra/web/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(weatherController *controller.ClimaController) *gin.Engine {
	r := gin.Default()
	r.GET("/clima/:cep", weatherController.BuscaClima)
	return r
}
