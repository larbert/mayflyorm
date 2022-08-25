package mayflyorm

import (
	"database/sql"

	"github.com/larbert/mayflylog"
	"github.com/larbert/mayflyorm/session"
)

type Engine struct {
	db *sql.DB
}

func New(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		mayflylog.Error(err)
		return
	}
	if err = db.Ping(); err != nil {
		mayflylog.Error(err)
		return
	}
	e = &Engine{db: db}
	mayflylog.Info("数据库连接成功")
	return
}

func (e *Engine) Close() {
	err := e.db.Close()
	if err != nil {
		mayflylog.Error(err)
		return
	}
	mayflylog.Info("数据库关闭成功")
}

func (e *Engine) NewSession() *session.Session {
	return session.New(e.db)
}
