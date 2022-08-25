package session

import (
	"database/sql"
	"strings"

	"github.com/larbert/mayflylog"
)

type Session struct {
	db        *sql.DB
	sql       strings.Builder
	sqlParams []interface{}
}

func New(db *sql.DB) *Session {
	return &Session{db: db}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlParams = nil
}

func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlParams = append(s.sqlParams, values...)
	return s
}

func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	mayflylog.Info(s.sql.String(), s.sqlParams)
	result, err = s.DB().Exec(s.sql.String(), s.sqlParams...)
	if err != nil {
		mayflylog.Error(err)
	}
	return
}

func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	mayflylog.Info(s.sql.String(), s.sqlParams)
	return s.DB().QueryRow(s.sql.String(), s.sqlParams...)
}

func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	mayflylog.Info(s.sql.String(), s.sqlParams)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlParams...); err != nil {
		mayflylog.Error(err)
	}
	return
}
