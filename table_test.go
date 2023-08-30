package geeorm

import (
	"testing"

	_ "github.com/lib/pq"
)

type Users struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {
	conStr := "host=localhost port=5432 dbname=gee-orm user=postgres password=postgres sslmode=disable"
	engine, _ := NewEngine("postgres", conStr)
	defer engine.Close()
	s := engine.NewSession().Model(&Users{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}
