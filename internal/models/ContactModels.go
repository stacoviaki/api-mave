package models

import (
	"time"

	"github.com/google/uuid"
)

type Contact struct {
	ID                    uuid.UUID `json:"id"`                               // Identificador único (UUID)
	ContactType           string    `json:"contact_type"`                     // Tipo de contato ((individual)CPF) ou (company)CNPJ)
	FullName              *string   `json:"full_name,omitempty"`              // Nome completo (pessoa física)
	TradeName             *string   `json:"trade_name,omitempty"`             // Nome fantasia (empresa)
	CompanyName           *string   `json:"company_name,omitempty"`           // Razão social (empresa)
	Cpf                   *string   `json:"cpf,omitempty"`                    // CPF
	Rg                    *string   `json:"rg,omitempty"`                     // RG
	Cnpj                  *string   `json:"cnpj,omitempty"`                   // CNPJ
	ForeingID             *string   `json:"foreing_id,omitempty"`             // ID estrangeiro (se aplicável)
	StateRegistration     *string   `json:"state_registration,omitempty"`     // Inscrição estadual
	MunicipalRegistration *string   `json:"municipal_registration,omitempty"` // Inscrição municipal
	PostalCode            *string   `json:"postal_code,omitempty"`            // CEP
	Street                *string   `json:"street,omitempty"`                 // Rua
	Number                *string   `json:"number,omitempty"`                 // Número
	Complement            *string   `json:"complement,omitempty"`             // Complemento
	District              *string   `json:"district,omitempty"`               // Bairro
	City                  *string   `json:"city,omitempty"`                   // Cidade
	State                 *string   `json:"state,omitempty"`                  // Estado
	CountryID             *int      `json:"country_id,omitempty"`             // Código do país (DDD ou ID do país)
	Phone                 *string   `json:"phone,omitempty"`                  // Telefone fixo
	Mobile                *string   `json:"mobile,omitempty"`                 // Celular
	Email                 *string   `json:"email,omitempty"`                  // E-mail
	Website               *string   `json:"website,omitempty"`                // Site
	JobTitle              *string   `json:"job_title,omitempty"`              // Cargo
	Tags                  *string   `json:"tags,omitempty"`                   // Tags de classificação
	CreatedAt             time.Time `json:"created_at,omitempty"`
	UpdatedAt             time.Time `json:"updated_at,omitempty"`
}
