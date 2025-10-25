package usecases

import (
	"github.com/google/uuid"
	"github.com/stacoviaki/api-mave/internal/models"
	"github.com/stacoviaki/api-mave/internal/repositories"
)

// 🔹 ContactUseCase
// Representa a camada de regras de negócio (intermediária entre Controller e Repository).
// Aqui ficam as lógicas que fazem sentido para o domínio da aplicação.
type ContactUseCase struct {
	repositories repositories.ContactRepositories // acesso ao banco via repositório
}

// 🔹 NewContactUseCase
// Construtor da camada de UseCase.
// Recebe um repositório já conectado e devolve o UseCase pronto pra uso.
func NewContactUseCase(repo repositories.ContactRepositories) ContactUseCase {
	return ContactUseCase{
		repositories: repo,
	}
}

// 🔹 GetContact
// Retorna todos os c contatosadastrados.
// Apenas repassa a chamada para o repositório, sem regras extras.
func (co *ContactUseCase) GetContacts() ([]models.Contact, error) {
	return co.repositories.GetContacts()
}

// 🔹 GetContactById
// Busca um contatos específico pelo seu UUID.
// Se o contatos não existir, retorna nil (sem erro).
func (co *ContactUseCase) GetContactById(id_contact uuid.UUID) (*models.Contact, error) {
	contact, err := co.repositories.GetContactById(id_contact)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

// 🔹 CreateContact
// Cria um novo contatos no banco de dados.
// Primeiro insere no repositório, depois adiciona o ID retornado ao struct.
func (co *ContactUseCase) CreateContact(contact models.Contact) (models.Contact, error) {
	// cria o contacto no banco e recebe o UUID gerado
	contactId, err := co.repositories.CreateContact(contact)
	if err != nil {
		return models.Contact{}, err
	}

	contact.ID = contactId
	return contact, nil
}

// 🔹 UpdateContact
// Atualiza os dados de um contato existente.
// Recebe o UUID e o novo objeto de contato, e retorna o contato atualizado.
func (co *ContactUseCase) UpdateContact(id_contact uuid.UUID, updatedContact models.Contact) (*models.Contact, error) {
	contact, err := co.repositories.UpdateContact(id_contact, updatedContact)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

// 🔹 DeleteContact
// Exclui um usuário do banco com base no UUID.
// Retorna o ID deletado em caso de sucesso.
func (co *ContactUseCase) DeleteUser(id_contact uuid.UUID) (uuid.UUID, error) {
	deletedID, err := co.repositories.DeleteContact(id_contact)
	if err != nil {
		return uuid.Nil, err
	}
	return deletedID, nil
}
