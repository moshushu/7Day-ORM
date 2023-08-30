package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

// Dialect 支持多种数据库
type Dialect interface {
	DataTypeOf(typ reflect.Value) string // 用于将Go语言的类型转换为该数据库的类型
	TableExistSQL(tableName string) (string, []interface{})
}

// RegisterDialect 注册dialect实例
// 需要新增数据库的支持时，调用注册全局即可
func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

// GetDialect 获取dialect实例
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
