package models

import "github.com/eduardonakaidev/configs"

func Insert(product Product) (err error){
	conn, err := configs.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()
	sql := `INSERT INTO products (name, description, price) VALUES($1,$2,$3) RETURNING id`
	err = conn.QueryRow(sql, product.ID,product.Name,product.Description,product.PriceInCents,product.Stock,product.CreatedAt,product.UpdatedAt).Scan()
	return 
}