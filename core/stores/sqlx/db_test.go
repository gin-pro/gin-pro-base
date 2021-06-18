package sqlx

import (
	"fmt"
	"testing"
)

func TestDB(t *testing.T) {
	c := &MysqlConf{
		Host:     "123456.com",
		Port:     "1234",
		Database: "123",
		Username: "root",
		Password: "123456",
		ShowSQL:  false,
	}
	_, err := InitMysql(c)
	if err != nil {
		fmt.Println(err)
	}
}
