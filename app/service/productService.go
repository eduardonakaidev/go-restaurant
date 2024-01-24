package service

import (
	"errors"
	"time"

	"github.com/eduardonakaidev/go-restaurant/connection"
	"github.com/eduardonakaidev/go-restaurant/models"
	"github.com/google/uuid"
)

func CreateProduct(product models.Product) (*models.Product, error) {
	// inicia a comunicação com o banco de dados
	conn, _ := connection.ConnectPostgresDB()
	// fecha a conecção ao terminar o codigo
	defer conn.Close()
	// executa a query sql
	var productAlreadyExists models.Product
	var err error
	_ = conn.QueryRow(`SELECT id FROM product WHERE name=$1`, product.Name).Scan(
		&productAlreadyExists.ID,

	)
	if productAlreadyExists.ID != uuid.Nil {
		return nil, errors.New("Produto com o mesmo nome já existe")
	}

	// verifica se deu erro na query
	if err != nil {
		return nil, err
	}
	// defini o id do produto
	product.ID = uuid.New()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	// query sql
	sqlqueryrow := `
	INSERT INTO "product" ("id", "name", "description", "priceInCents", "stock", "createdAt", "updatedAt", "isActive") 
	VALUES($1,$2,$3,$4,$5,$6,$7,$8);`
	// executa a query sql
	_, err = conn.Query(sqlqueryrow, product.ID, product.Name, product.Description, product.PriceInCents, product.Stock, product.CreatedAt, product.UpdatedAt, product.IsActive)

	return &product, err
}
func GetByIDProduct(id uuid.UUID) (product models.Product, err error) {
	// inicia a comunicação com o banco de dados
	conn, _ := connection.ConnectPostgresDB()
	// fecha a conecção ao terminar o codigo
	defer conn.Close()
	// defini o updatedat como nulltime
	
	// executa a query sql
	err = conn.QueryRow(`SELECT id, name, description, "priceInCents", stock, "createdAt", "updatedAt", "isActive" FROM product WHERE id=$1`, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.PriceInCents,
		&product.Stock,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.IsActive,
	)
	// verifica se deu erro na query
	if err != nil {
		return product, errors.New("Produto não encontrado")
	}
	// verifica se o updatedat esta null mesmo e adiciona ele ao objeto product

	
	return
}

func ListProduct() (products []models.Product, err error) {
	// inicia a comunicação com o banco de dados
	conn, _ := connection.ConnectPostgresDB()
	// fecha a conecção ao terminar o codigo
	defer conn.Close()
	// executa a query sql
	rows, err := conn.Query(`SELECT id, name, description, "priceInCents", stock, "createdAt", "updatedAt", "isActive" FROM product`)
	// verifica se deu erro na query
	if err != nil {
		return products,nil
	}
	for rows.Next() {
		var product models.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.PriceInCents, &product.Stock, &product.CreatedAt, &product.UpdatedAt, &product.IsActive)
		if err != nil {
			continue
		}
		products = append(products, product)
	}

	return products, err
}
func UpdateProduct(id uuid.UUID, productDto models.UpdateProductDTO) (*models.Product, error) {
	// Inicia a comunicação com o banco de dados
	conn, err := connection.ConnectPostgresDB()
	if err != nil {
		return nil, err
	}
	// Fecha a conexão ao término do código
	defer conn.Close()

	// Executa a query SQL para obter os dados existentes do produto
	var productdbquery models.Product
	err = conn.QueryRow(`
        SELECT id, name, description, "priceInCents", stock, "createdAt", "updatedAt", "isActive"
        FROM product WHERE id=$1
    `, id).Scan(
		&productdbquery.ID,
		&productdbquery.Name,
		&productdbquery.Description,
		&productdbquery.PriceInCents,
		&productdbquery.Stock,
		&productdbquery.CreatedAt,
		&productdbquery.UpdatedAt,
		&productdbquery.IsActive,
	)

	// Verifica se houve erro na query
	if err != nil {
		return nil, errors.New("Produto não encontrado")
	}

	// Atualiza os campos do produto com os valores do DTO
	productdbquery.Name = productDto.Name

	productdbquery.Description = productDto.Description

	productdbquery.PriceInCents = productDto.PriceInCents

	if productDto.Stock > 0 {
		productdbquery.Stock = productDto.Stock
	}
	productdbquery.UpdatedAt = time.Now()
	productdbquery.IsActive = productDto.IsActive

	// Executa a query SQL de atualização
	_, err = conn.Exec(`
        UPDATE product
        SET name=$1, description=$2, "priceInCents"=$3, stock=$4, "updatedAt"=$5
        WHERE id=$6
    `, productdbquery.Name, productdbquery.Description, productdbquery.PriceInCents,
		productdbquery.Stock, productdbquery.UpdatedAt, productdbquery.ID)

	// Retorna o produto atualizado e o erro (nil se não houver erro)
	if err != nil {
		return nil, errors.New("Erro ao atualizar o produto")
	}
	return &productdbquery,nil
}
func DeleteProduct(id uuid.UUID) error {
	// Inicia a comunicação com o banco de dados
	conn, err := connection.ConnectPostgresDB()
	if err != nil {
		return err
	}
	// Fecha a conexão ao término do código
	defer conn.Close()

	// Executa a query SQL de exclusão
	result, err := conn.Exec(`
        DELETE FROM product WHERE id=$1
    `, id)

	if err != nil {
		return err
	}

	// Verifica se pelo menos uma linha foi excluída
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Produto não encontrado")
	}

	// Retorna o erro (nil se não houver erro)
	return nil
}

