package sqlx

import (
	"fmt"
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

func InitMysql(conf *MysqlConf) (*xorm.Engine, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database)
	return DefaultDB("mysql", url)
}
