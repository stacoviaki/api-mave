package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID               uuid.UUID `json:"id"`                          // UUID
	ProductName      string    `json:"product_name"`                // Nome do produto (obrigatório)
	SalePrice        float64   `json:"sale_price"`                  // Preço de venda (obrigatório)
	ProductPrice     float64   `json:"product_price"`               // Preço de custo (obrigatório)
	ProductReference *string   `json:"product_reference,omitempty"` // Referência opcional

	FiscalType        string  `json:"fiscal_type"`         // Obrigatório
	ICMSOrigin        string  `json:"icms_origin"`         // Obrigatório
	NCM               string  `json:"ncm"`                 // Obrigatório
	ProductFiscalType string  `json:"product_fiscal_type"` // Obrigatório
	CEST              *string `json:"cest,omitempty"`      // Opcional
	NBM               *string `json:"nbm,omitempty"`       // Opcional

	CreatedAt time.Time `json:"created_at"` // Data de criação
	UpdatedAt time.Time `json:"updated_at"` // Data de atualização
}
