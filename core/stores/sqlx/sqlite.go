package sqlx

import (
	"xorm.io/xorm"
)

type SqliteConf struct {
	Database string `json:"database"`
	ShowSQL  bool   `json:"showsql"`
}

func InitSqlite(conf *SqliteConf) (*xorm.Engine, error) {
	return DefaultDB("sqlite3", conf.Database)
}
