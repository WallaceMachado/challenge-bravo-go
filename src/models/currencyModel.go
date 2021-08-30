package models

import "time"

type Currency struct {
	ID         string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name       string    `json:"name"`
	Code       string    `json:"code"`
	ValueInUSD float64   `json:"valueInUSD"`
	Created_at time.Time `json:"Created_at,omitempty"`
	Updated_at time.Time `json:"Updated_at,omitempty"`
}

type Teste2 struct {
	ValueInUSD float64 `json:"valueInUSD"`
}
