package models

import "valdson/store/db"

type Product struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func SearchAllProducts() []Product {
	db := db.ConnectDB()

	selectAllProducts, err := db.Query("select * from product order by id asc")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}

	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectAllProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		products = append(products, p)

	}

	defer db.Close()

	return products
}

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := db.ConnectDB()

	insertNewData, err := db.Prepare("insert into product(name, description,price,quantity) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertNewData.Exec(name, description, price, quantity)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectDB()

	deleteProduct, err := db.Prepare("delete from product where id = $1")
	if err != nil {
		panic(err.Error())

	}

	deleteProduct.Exec(id)

	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.ConnectDB()

	productActive, err := db.Query("select * from product where id = $1", id)

	if err != nil {
		panic(err.Error())
	}

	productToRefresh := Product{}

	for productActive.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = productActive.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())

		}

		productToRefresh.Id = id
		productToRefresh.Nome = nome
		productToRefresh.Descricao = descricao
		productToRefresh.Preco = preco
		productToRefresh.Quantidade = quantidade

	}

	defer db.Close()
	return productToRefresh

}

func UpdateProduct(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectDB()

	updateProduct, err := db.Prepare("update product set name = $1, description = $2, price = $3, quantity = $4 where id = $5")
	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
