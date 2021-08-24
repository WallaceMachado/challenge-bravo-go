package models

import "time"

//Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID         string    `json:"id,omitempty"`
	Name       string    `json:"name"`
	Code       string    `json:"code"`
	ValueInUSD float64   `json:"valueInUSD"`
	Created_at time.Time `json:"Created_at,omitempty"`
	Updated_at time.Time `json:"Updated_at,omitempty"`
}
