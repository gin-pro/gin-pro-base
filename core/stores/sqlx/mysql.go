package sqlx

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type MysqlConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	ShowSQL  bool   `json:"showsql"`
}

func NewMysql(host, port, database, userName, passed string, showSql ...bool) *MysqlConf {
	sSql := false
	if showSql != nil && len(showSql) > 0 {
		sSql = showSql[0]
	}
	return &MysqlConf{
		Host:     host,
		Port:     port,
		Database: database,
		Username: userName,
		Password: passed,
		ShowSQL:  sSql,
	}
}

func (c MysqlConf) defaultDB() (*xorm.Engine, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Database)
	engine, err := xorm.NewEngine("mysql", url)
	if err != nil {
		return nil, err
	}
	engine.ShowSQL(c.ShowSQL)
	return engine, err
}
