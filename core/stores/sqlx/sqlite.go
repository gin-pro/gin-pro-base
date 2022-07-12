package sqlx

import (
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

type SqliteConf struct {
	Database string `json:"database"`
	ShowSQL  bool   `json:"showsql"`
}

func NewSqlite(database string, showSql ...bool) *SqliteConf {
	sSql := false
	if showSql != nil && len(showSql) > 0 {
		sSql = showSql[0]
	}
	return &SqliteConf{
		Database: database,
		ShowSQL:  sSql,
	}
}
func (c SqliteConf) defaultDB() (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("sqlite3", c.Database)
	if err != nil {
		return nil, err
	}
	engine.ShowSQL(c.ShowSQL)
	return engine, err
}
