package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)

	var products []model.Product
	if err != nil {
		fmt.Println("Error querying products:", err)
		return []model.Product{}, err
	}

	for rows.Next() {
		var p model.Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)

		if err != nil {
			fmt.Println("Error scanning product:", err)
			return []model.Product{}, err
		}

		products = append(products, p)
	}

	rows.Close() // Close rows after we're done

	return products, nil
}
