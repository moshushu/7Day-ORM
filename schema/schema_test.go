package schema

import (
	"geeorm/dialect"
	"testing"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("postgres")

func TestParse(t *testing.T) {
	schema := Parse(&User{}, TestDial)
	if schema.Name != "user" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}
	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
}

func TestToSnakeCase(t *testing.T) {
	in := []string{"Name", "name", "personData", "PersonData"}
	out := []string{"name", "name", "person_data", "person_data"}
	for i := 0; i < len(in); i++ {
		if out[i] != toSnakeCase(in[i]) {
			t.Fatalf("in %s;out %s", out[i], in[i])
		}
	}
}
