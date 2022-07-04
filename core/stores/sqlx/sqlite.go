package sqlx

import (
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

type SqliteConf struct {
	Database string `json:"database"`
	ShowSQL  bool   `json:"showsql"`
}

func (c SqliteConf) DefaultDB() (*xorm.Engine, error) {
	return xorm.NewEngine("sqlite3", c.Database)
}
