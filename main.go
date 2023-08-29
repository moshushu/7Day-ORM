package main

import (
	"fmt"
	"gee-orm/handlers"

	_ "github.com/lib/pq"
)

func main() {
	conStr := "host=localhost port=5432 dbname=gee-orm user=postgres password=postgres sslmode=disable"
	engine, _ := handlers.NewEngin("postgres", conStr)
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("drop table if exists person;").Exec()
	_, _ = s.Raw("create table person(name text);").Exec()
	_, _ = s.Raw("create table person(name text);").Exec()
	res, _ := s.Raw("INSERT INTO person(name) values ($1), ($2);", "Tom", "Sam").Exec()
	count, _ := res.RowsAffected()
	fmt.Println("Exec success,affected", count)
}
