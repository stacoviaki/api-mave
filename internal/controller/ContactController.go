package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stacoviaki/api-mave/internal/models"
	"github.com/stacoviaki/api-mave/internal/usecases"
)

type contactController struct {
	contactUseCase usecases.ContactUseCase
}

// NewContactController cria uma nova instância do controller de contatos.
// Ele recebe a camada de usecase e retorna um controller configurado.
func NewContactController(usecase usecases.ContactUseCase) contactController {
	return contactController{
		contactUseCase: usecase,
	}
}

// GetContact responde à rota GET /contacs
// Busca todos os contatos cadastrados no banco.
func (co *contactController) GetContacts(ctx *gin.Context) {
	contact, err := co.contactUseCase.GetContacts()
	if err != nil {
		// Se ocorrer erro, retorna status 500 com o erro.
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	// Se der certo, retorna lista de contatos.
	ctx.JSON(http.StatusOK, contact)
}

// GetContactById responde à rota GET /contact/:contactId
// Busca um contato específico pelo ID.
func (co *contactController) GetContactById(ctx *gin.Context) {
	idParam := ctx.Param("contactId") // Pega o ID passado na URL.

	// Converte o ID de string para UUID.
	idUUID, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: "ID do contato precisa ser um UUID válido",
		})
		return
	}

	// Chama o usecase pra buscar o contato.
	contact, err := co.contactUseCase.GetContactById(idUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	// Caso o contato não exista.
	if contact == nil {
		ctx.JSON(http.StatusNotFound, models.Response{
			Message: "Contato não encontrado na base de dados",
		})
		return
	}

	// Retorna o contato encontrado.
	ctx.JSON(http.StatusOK, contact)
}

// CreateContact responde à rota POST /contact
// Cria um novo contato com base nos dados enviados no corpo da requisição.
func (co *contactController) CreateContact(ctx *gin.Context) {
	var contact models.Contact

	// Converte o JSON recebido em um struct Contact.
	err := ctx.BindJSON(&contact)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	// Envia o contato para a camada de usecase para salvar no banco.
	insertedContact, err := co.contactUseCase.CreateContact(contact)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	// Retorna o contato criado e status 201 (Created).
	ctx.JSON(http.StatusCreated, insertedContact)
}

// UpdateContact responde à rota PUT /contact/:contactId
// Atualiza os dados de um usuário existente.
func (co *contactController) UpdateContact(ctx *gin.Context) {
	idParam := ctx.Param("contactId") // ID na URL.

	// Converte o ID para UUID.
	uuid, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Lê o corpo da requisição e armazena em updatedContact.
	var updatedContact models.Contact
	if err := ctx.ShouldBindJSON(&updatedContact); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Chama o usecase pra atualizar o contato no banco.
	contact, err := co.contactUseCase.UpdateContact(uuid, updatedContact)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retorna o contato atualizado.
	ctx.JSON(http.StatusOK, contact)
}

// DeleteContact responde à rota DELETE /contact/:contactId
// Remove um contato do banco de dados pelo UUID.
func (co *contactController) DeleteContact(ctx *gin.Context) {
	idParam := ctx.Param("contactId") // Pega o ID na URL.

	// Converte o ID pra UUID.
	idUUID, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: "ID do Contato precisa ser um UUID válido",
		})
		return
	}

	// Chama o usecase pra deletar o contato.
	deletedID, err := co.contactUseCase.DeleteContact(idUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao deletar contato",
			"error":   err.Error(),
		})
		return
	}

	// Retorna confirmação de exclusão.
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Contato deletado com sucesso",
		"id":      deletedID,
	})
}
