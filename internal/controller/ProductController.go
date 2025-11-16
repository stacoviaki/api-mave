package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stacoviaki/api-mave/internal/models"
	"github.com/stacoviaki/api-mave/internal/usecases"
)

type ProductController struct {
	productUseCase usecases.ProductUsecase
}

func NewProductController(usecase usecases.ProductUsecase) ProductController {
	return ProductController{
		productUseCase: usecase,
	}
}

func (pr *ProductController) GetProducts(ctx *gin.Context) {
	product, err := pr.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (pr *ProductController) GetProductById(ctx *gin.Context) {
	idParam := ctx.Param("productId") // Pega o ID passado na URL.

	// Converte o ID de string para UUID.
	idUUID, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: "ID do contato precisa ser um UUID válido",
		})
		return
	}

	// Chama o usecase pra buscar o contato.
	product, err := pr.productUseCase.GetProductById(idUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	// Caso o contato não exista.
	if product == nil {
		ctx.JSON(http.StatusNotFound, models.Response{
			Message: "Contato não encontrado na base de dados",
		})
		return
	}

	// Retorna o contato encontrado.
	ctx.JSON(http.StatusOK, product)
}

func (pr *ProductController) CreateProduct(ctx *gin.Context) {
	var product models.Product

	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := pr.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (pr *ProductController) UpdateProduct(ctx *gin.Context) {
	idParam := ctx.Param("productId") // ID na URL.

	// Converte o ID para UUID.
	uuid, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Lê o corpo da requisição e armazena em updatedProduct.
	var updatedProduct models.Product
	if err := ctx.ShouldBindJSON(&updatedProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Chama o usecase pra atualizar o produto no banco.
	product, err := pr.productUseCase.UpdateProduct(uuid, updatedProduct)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retorna o produto atualizado.
	ctx.JSON(http.StatusOK, product)
}

func (pr *ProductController) DeleteProduct(ctx *gin.Context) {
	idParam := ctx.Param("productId") // Pega o ID na URL.

	// Converte o ID pra UUID.
	idUUID, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Message: "ID do Produto precisa ser um UUID válido",
		})
		return
	}

	// Chama o usecase pra produto o contato.
	deletedID, err := pr.productUseCase.DeleteProduct(idUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao deletar produto",
			"error":   err.Error(),
		})
		return
	}

	// Retorna confirmação de exclusão.
	ctx.JSON(http.StatusOK, gin.H{
		"message": "produto deletado com sucesso",
		"id":      deletedID,
	})
}
