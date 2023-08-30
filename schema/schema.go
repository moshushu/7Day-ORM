package schema

import (
	"bytes"
	"geeorm/dialect"
	"go/ast"
	"reflect"
	"unicode"
)

type Field struct {
	Name string // 字段名
	Type string // 字段类型
	Tag  string // 字段约束条件
}

type Schema struct {
	Model      interface{}       // 结构体（表结构）
	Name       string            // 表名
	Fields     []*Field          // 字段
	FieldNames []string          // 所有字段名
	fieldMap   map[string]*Field // 记录字段名与Field的映射关系
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

// Parse 将任意的对象解析为Schema实例
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     toSnakeCase(modelType.Name()),
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: toSnakeCase(p.Name),
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("geeorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, toSnakeCase(p.Name))
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}

func toSnakeCase(input string) string {
	var res bytes.Buffer
	for i, r := range input {
		if unicode.IsUpper(r) {
			if i > 0 {
				res.WriteRune('_')
			}
			res.WriteRune(unicode.ToLower(r))
		} else {
			res.WriteRune(r)
		}
	}
	return res.String()
}
