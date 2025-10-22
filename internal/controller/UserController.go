package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stacoviaki/api-mave/internal/models"
	"github.com/stacoviaki/api-mave/internal/usecases"
)

type userController struct {
	userUseCase usecases.UserUseCase
}

func NewUserController(usecase usecases.UserUseCase) userController {
	return userController{
		userUseCase: usecase,
	}
}

func (u *userController) GetUsers(ctx *gin.Context) {
	users, err := u.userUseCase.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, users)
}

func (u *userController) GetUserById(ctx *gin.Context) {
	idParam := ctx.Param("userId")

	// Tenta converter o parâmetro para UUID
	idUUID, err := uuid.Parse(idParam)
	if err != nil {
		response := models.Response{
			Message: "ID do Usuário precisa ser um UUID válido",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Busca o usuário pelo UUID
	user, err := u.userUseCase.GetUserById(idUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if user == nil {
		response := models.Response{
			Message: "Usuário não encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (u *userController) CreateUser(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	insertedUser, err := u.userUseCase.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedUser)
}

func (u *userController) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("userId")

	// Tenta converter o parâmetro para UUID
	idUUID, err := uuid.Parse(idParam)
	if err != nil {
		response := models.Response{
			Message: "ID do Usuário precisa ser um UUID válido",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	//deleta o usuário
	deletedID, err := u.userUseCase.DeleteUser(idUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao deletar usuário",
			"error":   err.Error(),
		})
		return
	}

	// Retorna resposta de sucesso
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Usuário deletado com sucesso",
		"id":      deletedID,
	})
}
