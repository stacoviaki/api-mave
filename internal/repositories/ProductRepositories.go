package repositories

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/stacoviaki/api-mave/internal/models"
)

type ProductRepositories struct {
	connection *sql.DB
}

func NewProductRepositories(connection *sql.DB) ProductRepositories {
	return ProductRepositories{
		connection: connection,
	}
}

func (pr *ProductRepositories) GetProducts() ([]models.Product, error) {
	// Query no banco
	query := "SELECT * FROM public.products"

	// Executar a query
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []models.Product{}, err
	}

	var productList []models.Product // lista com todos os produtos
	var productObj models.Product    // armazena temporariamente cada linha

	// Percorre cada linha retornada e preenche os dados no struct
	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.ProductName,
			&productObj.SalePrice,
			&productObj.ProductPrice,
			&productObj.ProductReference,
			&productObj.FiscalType,
			&productObj.ICMSOrigin,
			&productObj.NCM,
			&productObj.ProductFiscalType,
			&productObj.CEST,
			&productObj.NBM,
			&productObj.CreatedAt,
			&productObj.UpdatedAt,
		)
		if err != nil {
			fmt.Println(err)
			return []models.Product{}, err
		}

		// Adiciona o produto lido à lista
		productList = append(productList, productObj)
	}

	rows.Close()
	return productList, nil
}

func (pr *ProductRepositories) GetProductById(uuid_product uuid.UUID) (*models.Product, error) {
	// Prepara a query SQL com parâmetro (evita SQL Injection)
	query, err := pr.connection.Prepare("SELECT * FROM products WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var product models.Product

	// Executa a query e armazena o resultado no struct
	err = query.QueryRow(uuid_product).Scan(
		&product.ID,
		&product.ProductName,
		&product.SalePrice,
		&product.ProductPrice,
		&product.ProductReference,
		&product.FiscalType,
		&product.ICMSOrigin,
		&product.NCM,
		&product.ProductFiscalType,
		&product.CEST,
		&product.NBM,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		// Se não encontrar nenhuma linha, retorna nil sem erro
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	query.Close()
	return &product, nil
}

func (pr *ProductRepositories) CreateProduct(product models.Product) (uuid.UUID, error) {
	var id uuid.UUID

	query, err := pr.connection.Prepare(`
		INSERT INTO products
		(product_name, sale_price, product_price, product_reference, fiscal_type, icms_origin, ncm, product_fiscal_type, cest, nbm) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id
	`)
	if err != nil {
		fmt.Println(err)
		return uuid.Nil, err
	}

	err = query.QueryRow(
		&product.ProductName,
		&product.SalePrice,
		&product.ProductPrice,
		&product.ProductReference,
		&product.FiscalType,
		&product.ICMSOrigin,
		&product.NCM,
		&product.ProductFiscalType,
		&product.CEST,
		&product.NBM,
	).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return uuid.Nil, err
	}

	query.Close()
	return id, nil
}

func (repo *ProductRepositories) UpdateProduct(id_product uuid.UUID, product models.Product) (*models.Product, error) {
	query := `
		UPDATE products
		SET 
			product_name = $1,
			sale_price = $2,
			product_price = $3,
			product_reference = $4,
			fiscal_type = $5,
			icms_origin = $6,
			ncm = $7,
			product_fiscal_type = $8,
			cest = $9,
			nbm = $10,
			updated_at = now()
		WHERE id = $11
		RETURNING 
			id, product_name, sale_price, product_price, product_reference,
			fiscal_type, icms_origin, ncm, product_fiscal_type, cest, nbm,
			updated_at;
		`

	row := repo.connection.QueryRow(
		query,
		product.ProductName, product.SalePrice, product.ProductPrice, product.ProductReference, product.FiscalType, product.ICMSOrigin, product.NCM, product.ProductFiscalType, product.CEST, product.NBM, id_product,
	)

	var updatedProduct models.Product

	err := row.Scan(
		&updatedProduct.ID,
		&updatedProduct.ProductName,
		&updatedProduct.SalePrice,
		&updatedProduct.ProductPrice,
		&updatedProduct.ProductReference,
		&updatedProduct.FiscalType,
		&updatedProduct.ICMSOrigin,
		&updatedProduct.NCM,
		&updatedProduct.ProductFiscalType,
		&updatedProduct.CEST,
		&updatedProduct.NBM,
		&updatedProduct.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &updatedProduct, nil
}

func (pr *ProductRepositories) DeleteProduct(id uuid.UUID) (uuid.UUID, error) {
	var deletedID uuid.UUID

	// Executa o DELETE e retorna o ID apagado
	query := `DELETE FROM products WHERE id = $1 RETURNING id;`
	err := pr.connection.QueryRow(query, id).Scan(&deletedID)
	if err != nil {
		fmt.Println("Erro ao deletar contato:", err)
		return uuid.Nil, err
	}

	return deletedID, nil
}
