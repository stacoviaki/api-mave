package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stacoviaki/api-mave/internal/models"
	"github.com/stacoviaki/api-mave/internal/usecases"
)

// userController é o responsável por receber as requisições HTTP
// e repassar os dados para a camada de usecase.
// Ele não contém regras de negócio — apenas entrada e saída de dados.
type userController struct {
	userUseCase usecases.UserUseCase
}

// NewUserController cria uma nova instância do controller de usuários.
// Ele recebe a camada de usecase e retorna um controller configurado.
func NewUserController(usecase usecases.UserUseCase) userController {
	return userController{
		userUseCase: usecase,
	}
}

// GetUsers responde à rota GET /users
// Busca todos os usuários cadastrados no banco.
func (u *userController) GetUsers(ctx *gin.Context) {
	users, err := u.userUseCase.GetUsers()
	if err != nil {
		// Se ocorrer erro, retorna status 500 com o erro.
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	// Se der certo, retorna lista de usuários.
	ctx.JSON(http.StatusOK, users)
}

// GetUserById responde à rota GET /user/:userId
// Busca um usuário específico pelo ID.
func (u *userController) GetUserById(ctx *gin.Context) {
	idParam := ctx.Param("userId") // Pega o ID passado na URL.

	// Converte o ID de string para UUID.
	idUUID, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: "ID do Usuário precisa ser um UUID válido",
		})
		return
	}

	// Chama o usecase pra buscar o usuário.
	user, err := u.userUseCase.GetUserById(idUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	// Caso o usuário não exista.
	if user == nil {
		ctx.JSON(http.StatusNotFound, models.Response{
			Message: "Usuário não encontrado na base de dados",
		})
		return
	}

	// Retorna o usuário encontrado.
	ctx.JSON(http.StatusOK, user)
}

// CreateUser responde à rota POST /user
// Cria um novo usuário com base nos dados enviados no corpo da requisição.
func (u *userController) CreateUser(ctx *gin.Context) {
	var user models.User

	// Converte o JSON recebido em um struct User.
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	// Envia o usuário para a camada de usecase para salvar no banco.
	insertedUser, err := u.userUseCase.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	// Retorna o usuário criado e status 201 (Created).
	ctx.JSON(http.StatusCreated, insertedUser)
}

// UpdateUser responde à rota PUT /user/:userId
// Atualiza os dados de um usuário existente.
func (u *userController) UpdateUser(ctx *gin.Context) {
	idParam := ctx.Param("userId") // ID na URL.

	// Converte o ID para UUID.
	uuid, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Lê o corpo da requisição e armazena em updatedUser.
	var updatedUser models.User
	if err := ctx.ShouldBindJSON(&updatedUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Chama o usecase pra atualizar o usuário no banco.
	user, err := u.userUseCase.UpdateUser(uuid, updatedUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retorna o usuário atualizado.
	ctx.JSON(http.StatusOK, user)
}

// DeleteUser responde à rota DELETE /user/:userId
// Remove um usuário do banco de dados pelo UUID.
func (u *userController) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("userId") // Pega o ID na URL.

	// Converte o ID pra UUID.
	idUUID, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: "ID do Usuário precisa ser um UUID válido",
		})
		return
	}

	// Chama o usecase pra deletar o usuário.
	deletedID, err := u.userUseCase.DeleteUser(idUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao deletar usuário",
			"error":   err.Error(),
		})
		return
	}

	// Retorna confirmação de exclusão.
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Usuário deletado com sucesso",
		"id":      deletedID,
	})
}
