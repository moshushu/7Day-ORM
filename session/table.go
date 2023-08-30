// 用于放置操作数据库表相关的代码
package session

import (
	"fmt"
	"geeorm/log"
	"geeorm/schema"
	"reflect"
	"strings"
)

func (s *Session) Model(value interface{}) *Session {
	// nil or different model,update refTable
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("Model is not set")
	}
	return s.refTable
}

func (s *Session) CreateTable() error {
	table := s.RefTable()
	var cols []string
	for _, f := range table.Fields {
		cols = append(cols, fmt.Sprintf("%s %s %s", f.Name, f.Type, f.Tag))
	}
	desc := strings.Join(cols, ",")
	_, err := s.Raw(fmt.Sprintf("CREATE TABLE  %s (%s)", table.Name, desc)).Exec()
	return err
}

func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s", s.RefTable().Name)).Exec()
	return err
}

func (s *Session) HasTable() bool {
	sql, values := s.dialect.TableExistSQL(s.RefTable().Name)
	row := s.Raw(sql, values...).QueryRow()
	fmt.Println("sql", sql)
	fmt.Println("values", values)
	var tmp string
	err := row.Scan(&tmp)
	fmt.Println("tmp", tmp)
	fmt.Println("s.reftable.name", s.RefTable().Name)
	log.Error(err)
	return tmp == s.RefTable().Name
}
