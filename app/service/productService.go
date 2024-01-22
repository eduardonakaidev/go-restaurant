package service

import (
	"database/sql"
	"time"

	"github.com/eduardonakaidev/go-restaurant/connection"
	"github.com/eduardonakaidev/go-restaurant/models"
	"github.com/google/uuid"
)


func CreateProduct(product models.Product) (*models.Product, error) {
	// inicia a comunicação com o banco de dados
	conn ,_ := connection.ConnectPostgresDB()
	// fecha a conecção ao terminar o codigo
	defer conn.Close()
	// defini o id do produto
	product.ID = uuid.New()
	// defini o createdat para a hora atual
	product.CreatedAt = time.Now()
	// query sql
	sqlqueryrow := `
	INSERT INTO "product" ("id", "name", "description", "priceInCents", "stock", "createdAt", "updatedAt", "isActive") 
	VALUES($1,$2,$3,$4,$5,$6,$7,$8);`
	// executa a query sql
	_, err := conn.Query(sqlqueryrow, product.ID, product.Name, product.Description, product.PriceInCents, product.Stock, product.CreatedAt, product.UpdatedAt, product.IsActive)

	return &product, err
}
func GetByIDProduct(id uuid.UUID)(product models.Product, err error){
// inicia a comunicação com o banco de dados
	conn ,_ := connection.ConnectPostgresDB()
	// fecha a conecção ao terminar o codigo
	defer conn.Close()
	// defini o updatedat como nulltime
	var updatedAt sql.NullTime
	// executa a query sql
    err = conn.QueryRow(`SELECT id, name, description, "priceInCents", stock, "createdAt", "updatedAt", "isActive" FROM product WHERE id=$1`, id).Scan(
        &product.ID,
        &product.Name,
        &product.Description,
        &product.PriceInCents,
        &product.Stock,
        &product.CreatedAt,
        &updatedAt,  
        &product.IsActive,
    )
	// verifica se deu erro na query
    if err != nil {
        return product, err
    }
	// verifica se o updatedat esta null memso e adiciona ele ao objeto product

    if updatedAt.Valid {
        product.UpdatedAt = &updatedAt.Time
    } else {
        product.UpdatedAt = nil
    }

	return




}
