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

func (c MysqlConf) DefaultDB() (*xorm.Engine, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
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
