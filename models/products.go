package models

import (
	"web_store/db"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func SelectAllProducts() []Product {
	db := db.Dbconnector()

	selectAllProducts, err := db.Query("Select * from products order by product_id asc")
	if err != nil {
		panic(err.Error())
	}

	dbProduct := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		dbProduct.ID = id
		dbProduct.Name = name
		dbProduct.Description = description
		dbProduct.Price = price
		dbProduct.Quantity = quantity

		products = append(products, dbProduct)
	}
	defer db.Close()
	return products
}

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := db.Dbconnector()
	insertData, err := db.Prepare("insert into products(product_name, product_description, product_price, product_quantity) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insertData.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.Dbconnector()
	deleteData, err := db.Prepare("delete from products where product_id =$1")
	if err != nil {
		panic(err.Error())
	}
	deleteData.Exec(id)
	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.Dbconnector()
	dbProduct, err := db.Query("select * from products where product_id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	editedProduct := Product{}

	for dbProduct.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = dbProduct.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		editedProduct.ID = id
		editedProduct.Name = name
		editedProduct.Description = description
		editedProduct.Price = price
		editedProduct.Quantity = quantity
	}
	defer db.Close()
	return editedProduct
}

func UpdateProduct(id int, name string, description string, price float64, quantity int) {
	db := db.Dbconnector()
	updateProd, err := db.Prepare("update products set product_name=$1, product_description=$2, product_price=$3, product_quantity=$4 where product_id=$5")
	if err != nil {
		panic(err.Error())
	}
	updateProd.Exec(name, description, price, quantity, id)
	defer db.Close()
}
