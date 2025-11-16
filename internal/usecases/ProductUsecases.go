package usecases

import (
	"github.com/google/uuid"
	"github.com/stacoviaki/api-mave/internal/models"
	"github.com/stacoviaki/api-mave/internal/repositories"
)

type ProductUsecase struct {
	repositories repositories.ProductRepositories
}

func NewProductUseCase(repo repositories.ProductRepositories) ProductUsecase {
	return ProductUsecase{
		repositories: repo,
	}
}

func (pr *ProductUsecase) GetProducts() ([]models.Product, error) {
	return pr.repositories.GetProducts()
}

func (pr *ProductUsecase) CreateProduct(product models.Product) (models.Product, error) {
	productId, err := pr.repositories.CreateProduct(product)
	if err != nil {
		return models.Product{}, err
	}

	product.ID = productId
	return product, nil
}

func (pr *ProductUsecase) GetProductById(uuid_product uuid.UUID) (*models.Product, error) {
	product, err := pr.repositories.GetProductById(uuid_product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pr *ProductUsecase) UpdateProduct(id_product uuid.UUID, UpdateProduct models.Product) (*models.Product, error) {
	product, err := pr.repositories.UpdateProduct(id_product, UpdateProduct)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pr *ProductUsecase) DeleteProduct(product uuid.UUID) (uuid.UUID, error) {
	deletedID, err := pr.repositories.DeleteProduct(product)
	if err != nil {
		return uuid.Nil, err
	}
	return deletedID, nil
}
