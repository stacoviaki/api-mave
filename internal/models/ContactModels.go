package models

import (
	"time"

	"github.com/google/uuid"
)

type Contact struct {
	ID                    uuid.UUID `json:"id"`                     // Identificador único (UUID)
	ContactType           string    `json:"contact_type"`           // Tipo de contato ((individual)CPF ou (company)CNPJ)
	FullName              string    `json:"full_name"`              // Nome completo (pessoa física)
	TradeName             string    `json:"trade_name"`             // Nome fantasia (empresa)
	CompanyName           string    `json:"company_name"`           // Razão social (empresa)
	Cpf                   string    `json:"cpf"`                    // CPF
	Rg                    string    `json:"rg"`                     // RG
	Cnpj                  string    `json:"cnpj"`                   // CNPJ
	ForeingID             string    `json:"foreing_id"`             // ID estrangeiro (se aplicável)
	StateRegistration     string    `json:"state_registration"`     // Inscrição estadual
	MunicipalRegistration string    `json:"municipal_registration"` // Inscrição municipal
	PostalCode            string    `json:"postal_code"`            // CEP
	Street                string    `json:"street"`                 // Rua
	Number                string    `json:"number"`                 // Número
	Complement            string    `json:"complement"`             // Complemento
	District              string    `json:"district"`               // Bairro
	City                  string    `json:"city"`                   // Cidade
	State                 string    `json:"state"`                  // Estado
	CountryID             int       `json:"country_id"`             // Código do país (DDD ou ID do país)
	Phone                 string    `json:"phone"`                  // Telefone fixo
	Mobile                string    `json:"mobile"`                 // Celular
	Email                 string    `json:"email"`                  // E-mail
	Website               string    `json:"website"`                // Site
	JobTitle              string    `json:"job_title"`              // Cargo
	Tags                  string    `json:"tags"`                   // Tags de classificação
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}
