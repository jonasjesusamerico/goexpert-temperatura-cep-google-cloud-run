package domain

type ClimaResponse struct {
	Current Clima `json:"current"`
}

type Clima struct {
	TemperaturaCelsius float64 `json:"temp_c"`
}
