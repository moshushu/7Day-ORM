package handlers

import (
	"database/sql"
	"gee-orm/log"
	"gee-orm/session"
)

type Engine struct {
	db *sql.DB
}

// NewEngin 连接数据库，返回 *sql.DB
// 调用db.Ping()，检查数据库是否正常连接
func NewEngin(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	e = &Engine{db: db}
	log.Info("Connect database success")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db)
}
