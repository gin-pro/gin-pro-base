package sqlx

import (
	"xorm.io/xorm"
)

type InitDB interface {
	DefaultDB() (*xorm.Engine, error)
}
