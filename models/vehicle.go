package models

import "time"

type Vehicle struct {
	ID          int64     `json:"id"`
	Type        string    `json:"type"`        // tipo do veículo (car, moto, caminhão, etc)
	Brand       string    `json:"brand"`       // marca
	Model       string    `json:"model"`       // modelo
	Year        int       `json:"year"`        // ano de fabricação
	Price       float64   `json:"price"`       // preço
	Description string    `json:"description"` // descrição opcional
	CreatedAt   time.Time `json:"created_at"`  // data de criação
	UpdatedAt   time.Time `json:"updated_at"`  // data de atualização
}
