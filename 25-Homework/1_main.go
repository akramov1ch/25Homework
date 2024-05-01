// // 1_may.go

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/lib/pq"
// )

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "vakhaboff"
// 	dbname   = "shaxboz"
// )

// type Customer struct {
// 	ID   int
// 	Name string
// 	City string
// }

// func main() {
// 	// dbURL := "postgres://postgres:postgres@localhost:5432/demo?sslmode=disable"
// 	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)

// 	// fmt.Println(postgresqlDbInfo)

// 	db, err := sql.Open("postgres", postgresqlDbInfo)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal("failed to ping", err)
// 	}

// 	fmt.Println("successfully connected")

// 	id := 10000000

// 	var c Customer
// 	sql := `
// 		SELECT customer_id, customer_name, city
// 		FROM customers
// 		WHERE customer_id = $1
// 	`

// 	err = db.QueryRow(sql, id).Scan(&c.ID, &c.Name, &c.City)
// 	if err != nil {
// 		log.Fatal("failed to query row", err)
// 	}

// 	fmt.Printf("%+v\n", c)
// }







package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "vakhaboff"
	dbname   = "shaxboz"
)

type Customer struct {
	ID   int
	Name string
	City string
}

func main() {
	// dbURL := "postgres://postgres:postgres@localhost:5432/demo?sslmode=disable"
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// fmt.Println(postgresqlDbInfo)

	db, err := sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("failed to ping", err)
	}

	fmt.Println("successfully connected")

	country := os.Args[1]
	// city := os.Args[2]

	var c Customer
	sql := `
		SELECT customer_id, customer_name, city
		FROM customers
		WHERE country = $1
	`
	var customers []Customer

	rows, err := db.Query(sql, country)
	if err != nil {
		log.Fatal("failed to query rows", err)
	}
	defer rows.Close()

	for rows.Next() {
		var c Customer
		err = rows.Scan(&c.ID, &c.City, &c.Name)
		if err != nil {
			log.Fatal("failed to scan a row", err)
		}
		customers = append(customers, c)
	}

	for _, v := range customers {
		fmt.Printf("%+v\n", v)
	}

	// err = db.QueryRow(sql, country).Scan(&c.ID, &c.City, &c.Name)
	// if err != nil {
	// 	log.Fatal("failed to query row", err)
	// }

	fmt.Printf("%+v\n", c)
}
