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

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	query, err := pr.connection.Prepare("INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return 0, err
	}
	defer query.Close() // Close the prepared statement

	var id int
	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return 0, err
	}

	return id, nil
}
