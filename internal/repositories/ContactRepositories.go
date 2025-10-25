package repositories

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/stacoviaki/api-mave/internal/models"
)

// ContactRepositories é a camada responsável por acessar o banco de dados
// e executar operações na tabela `contacts` (SELECT, INSERT, UPDATE, DELETE).
type ContactRepositories struct {
	connection *sql.DB // conexão ativa com o PostgreSQL
}

func NewContactRepositories(connection *sql.DB) ContactRepositories {
	return ContactRepositories{
		connection: connection,
	}
}

// GetContacts busca todos os contatos cadastrados no banco de dados.
func (co *ContactRepositories) GetContacts() ([]models.Contact, error) {
	query := "SELECT * FROM public.contacts"

	// Executa a query no banco.
	rows, err := co.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []models.Contact{}, err
	}

	var ContactList []models.Contact // lista com todos os contatos
	var ContactObj models.Contact    // armazena temporariamente cada linha

	// Percorre cada linha retornada e preenche os dados no struct
	for rows.Next() {
		err = rows.Scan(
			&ContactObj.ID,
			&ContactObj.ContactType,
			&ContactObj.FullName,
			&ContactObj.TradeName,
			&ContactObj.CompanyName,
			&ContactObj.Cpf,
			&ContactObj.Rg,
			&ContactObj.Cnpj,
			&ContactObj.ForeingID,
			&ContactObj.StateRegistration,
			&ContactObj.MunicipalRegistration,
			&ContactObj.PostalCode,
			&ContactObj.Street,
			&ContactObj.Number,
			&ContactObj.Complement,
			&ContactObj.District,
			&ContactObj.City,
			&ContactObj.State,
			&ContactObj.CountryID,
			&ContactObj.Phone,
			&ContactObj.Mobile,
			&ContactObj.Email,
			&ContactObj.Website,
			&ContactObj.JobTitle,
			&ContactObj.Tags,
			&ContactObj.CreatedAt,
			&ContactObj.UpdatedAt,
		)
		if err != nil {
			fmt.Println(err)
			return []models.Contact{}, err
		}

		ContactList = append(ContactList, ContactObj)
	}

	rows.Close() // fecha o cursor do banco
	return ContactList, nil
}

// GetContactById busca um unico contato pelo UUID
func (co *ContactRepositories) GetContactById(uuid_contact uuid.UUID) (*models.Contact, error) {
	// Prepara a query SQL com parâmetro (evita SQL Injection)
	query, err := co.connection.Prepare("SELECT * FROM contacts WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var contact models.Contact

	// Executa a query e armazena o resultado no struct
	err = query.QueryRow(uuid_contact).Scan(
		&contact.ID,
		&contact.ContactType,
		&contact.FullName,
		&contact.TradeName,
		&contact.CompanyName,
		&contact.Cpf,
		&contact.Rg,
		&contact.Cnpj,
		&contact.ForeingID,
		&contact.StateRegistration,
		&contact.MunicipalRegistration,
		&contact.PostalCode,
		&contact.Street,
		&contact.Number,
		&contact.Complement,
		&contact.District,
		&contact.City,
		&contact.State,
		&contact.CountryID,
		&contact.Phone,
		&contact.Mobile,
		&contact.Email,
		&contact.Website,
		&contact.JobTitle,
		&contact.Tags,
		&contact.CreatedAt,
		&contact.UpdatedAt,
	)
	if err != nil {
		// Se não encontrar nenhuma linha, retorna nil sem erro
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	query.Close()
	return &contact, nil
}

// CreateContact insere um novo contatos na tabela `Contacts` e retorna o UUID gerado.
func (co *ContactRepositories) CreateContact(contact models.Contact) (uuid.UUID, error) {
	var id uuid.UUID

	// Cria o comando SQL com RETURNING id para capturar o UUID criado
	query, err := co.connection.Prepare(`
		INSERT INTO contacts 
		(contact_type, full_name, trade_name, company_name, cpf, rg, cnpj, foreing_id, state_registration, municipal_registration, postal_code, street, number, complement, district, city, state, country_id, phone, mobile, email, website, job_title, tags) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24) RETURNING id
	`)
	if err != nil {
		fmt.Println(err)
		return uuid.Nil, err
	}

	// Executa o INSERT e pega o ID retornado pelo banco
	err = query.QueryRow(
		&contact.ContactType,
		&contact.FullName,
		&contact.TradeName,
		&contact.CompanyName,
		&contact.Cpf,
		&contact.Rg,
		&contact.Cnpj,
		&contact.ForeingID,
		&contact.StateRegistration,
		&contact.MunicipalRegistration,
		&contact.PostalCode,
		&contact.Street,
		&contact.Number,
		&contact.Complement,
		&contact.District,
		&contact.City,
		&contact.State,
		&contact.CountryID,
		&contact.Phone,
		&contact.Mobile,
		&contact.Email,
		&contact.Website,
		&contact.JobTitle,
		&contact.Tags,
	).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return uuid.Nil, err
	}

	query.Close()
	return id, nil
}

func (repo *ContactRepositories) UpdateContact(id_contact uuid.UUID, contact models.Contact) (*models.Contact, error) {
	query := `
		UPDATE contacts
		SET 
			contact_type = $1,
			full_name = $2,
			trade_name = $3,
			company_name = $4,
			cpf = $5,
			rg = $6,
			cnpj = $7,
			foreing_id = $8,
			state_registration = $9,
			municipal_registration = $10,
			postal_code = $11,
			street = $12,
			number = $13,
			complement = $14,
			district = $15,
			city = $16,
			state = $17,
			country_id = $18,
			phone = $19,
			mobile = $20,
			email = $21,
			website = $22,
			job_title = $23,
			tags = $24,
			updated_at = now()
		WHERE id = $25
		RETURNING 
			id, contact_type, full_name, trade_name, company_name,
			cpf, rg, cnpj, foreing_id, state_registration, municipal_registration,
			postal_code, street, number, complement, district, city, state,
			country_id, phone, mobile, email, website, job_title, tags,
			created_at, updated_at;
	`

	row := repo.connection.QueryRow(
		query,
		contact.ContactType, contact.FullName, contact.TradeName, contact.CompanyName,
		contact.Cpf, contact.Rg, contact.Cnpj, contact.ForeingID, contact.StateRegistration,
		contact.MunicipalRegistration, contact.PostalCode, contact.Street, contact.Number,
		contact.Complement, contact.District, contact.City, contact.State, contact.CountryID,
		contact.Phone, contact.Mobile, contact.Email, contact.Website, contact.JobTitle,
		contact.Tags, id_contact,
	)

	var updatedContact models.Contact

	err := row.Scan(
		&updatedContact.ID, &updatedContact.ContactType, &updatedContact.FullName,
		&updatedContact.TradeName, &updatedContact.CompanyName, &updatedContact.Cpf,
		&updatedContact.Rg, &updatedContact.Cnpj, &updatedContact.ForeingID,
		&updatedContact.StateRegistration, &updatedContact.MunicipalRegistration,
		&updatedContact.PostalCode, &updatedContact.Street, &updatedContact.Number,
		&updatedContact.Complement, &updatedContact.District, &updatedContact.City,
		&updatedContact.State, &updatedContact.CountryID, &updatedContact.Phone,
		&updatedContact.Mobile, &updatedContact.Email, &updatedContact.Website,
		&updatedContact.JobTitle, &updatedContact.Tags, &updatedContact.CreatedAt,
		&updatedContact.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &updatedContact, nil
}

// DeleteContact remove um contato do banco de dados pelo seu UUID.
func (co *ContactRepositories) DeleteContact(id uuid.UUID) (uuid.UUID, error) {
	var deletedID uuid.UUID

	// Executa o DELETE e retorna o ID apagado
	query := `DELETE FROM contacts WHERE id = $1 RETURNING id;`
	err := co.connection.QueryRow(query, id).Scan(&deletedID)
	if err != nil {
		fmt.Println("Erro ao deletar contato:", err)
		return uuid.Nil, err
	}

	return deletedID, nil
}
