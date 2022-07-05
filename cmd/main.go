package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type house struct {
	address         string
	houseNumber     uint
	apartmentsCount uint
	floorsCount     uint
}

func main() {
	connStr := "user=postgres password=5gq5qe95 dbname=apartmentdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//result, err := db.Exec("insert into houses (address, housenumber, apartmentscount, floorscount) "+
	//	"values ('Карамысло, дом 5', $1, $2, $3)",
	//	99, 53, )
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(result.RowsAffected()) // количество добавленных строк

	//rows, err := db.Query("select * from houses")
	//if err != nil {
	//	panic(err)
	//}
	//defer rows.Close()
	//houses := []house{}
	//
	//for rows.Next() {
	//	p := house{}
	//	err := rows.Scan(&p.address, &p.houseNumber, &p.apartmentsCount, &p.floorsCount)
	//	if err != nil {
	//		fmt.Println(err)
	//		continue
	//	}
	//	houses = append(houses, p)
	//}
	//for _, p := range houses {
	//	fmt.Println(p.address, p.houseNumber, p.apartmentsCount, p.floorsCount)
	//}

	//result, err := db.Exec("update houses set apartmentsCount = $1 where houseNumber = $2", 777, 5)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(result.RowsAffected()) // количество обновленных строк

	result, err := db.Exec("delete from Products where id = $1", 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected()) // количество удаленных строк
}
