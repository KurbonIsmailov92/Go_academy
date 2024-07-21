package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Product struct {
	id      int
	model   string
	company string
	price   int
}

func ConnectToDB() (*sql.DB, error) {
	connStr := "user=kurbon password=ismoil dbname=online_store_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to database")

	return db, nil
}

func CloseDBConn(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT id, model, company, price FROM products;")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		err = rows.Scan(&p.id, &p.model, &p.company, &p.price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}

	return products, nil
}

func GetProductByID(db *sql.DB, id int) (Product, error) {
	var p Product
	row := db.QueryRow("SELECT id, model, company, price FROM products WHERE id = $1", id)
	err := row.Scan(&p.id, &p.model, &p.company, &p.price)
	if err != nil {
		return Product{}, err
	}

	return p, nil
}

func main() {

	// LIFO - Last In First Out
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//defer fmt.Println(3)
	//defer fmt.Println(7)
	// 5 6 1 2 3 7

	db, err := ConnectToDB()
	if err != nil {
		panic(err)
	}

	defer CloseDBConn(db)

	products, err := GetAllProducts(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, product := range products {
		fmt.Println(product.id, product.model, product.company, product.price)
	}

	product, err := GetProductByID(db, 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(product.id, product.model, product.company, product.price)

}
