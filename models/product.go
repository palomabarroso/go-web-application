package models

import (
	"github.com/palomabarroso/go-web-application/db"
)

type Product struct {
	ID                int
	Name, Description string
	Price             float64
	Quantity          int
}

func ListAllProducts() []Product {
	db := db.ConnectDB()
	allProducts, err := db.Query("SELECT * FROM products")
	if err != nil {
		panic(err.Error())
	}
	p := Product{}
	products := []Product{}

	for allProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = allProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.ID = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	defer db.Close()
	return products
}

func NewProduct(name, description string, price float64, quantity int) {
	db := db.ConnectDB()
	newProduct, err := db.Prepare("INSERT INTO products (name, description, price, quantity) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	newProduct.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectDB()
	deleteProduct, err := db.Prepare("DELETE FROM products WHERE id = $1")

	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.ConnectDB()

	produtoDoBanco, err := db.Query("select * from products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	productToUpdate := Product{}

	for produtoDoBanco.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = produtoDoBanco.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		productToUpdate.ID = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Price = price
		productToUpdate.Quantity = quantity
	}
	defer db.Close()
	return productToUpdate
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := db.ConnectDB()

	updateProduct, err := db.Prepare("UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	updateProduct.Exec(name, description, price, quantity, id)
	defer db.Close()
}
