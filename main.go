package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type CITY struct {
	name       string
	population int
}

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=yoichiro password=password dbname=go_practice")

	defer db.Close()
	if err != nil {
		fmt.Println(err)

	}

	rows, err := db.Query("select * from cities")
	if err != nil {
		fmt.Println(err)
	}

	var cities []CITY
	for rows.Next() {
		var c CITY
		rows.Scan(&c.name, &c.population)
		cities = append(cities, c)
	}
	// name | population
	// tokyo| 900
	// osaka | 800
	// nagoya | 600

	fmt.Printf("%v", cities)

}
