package sqlx

import "xorm.io/xorm"

type DBI interface {
	defaultDB() (*xorm.Engine, error)
}

func DefaultDB[T DBI](db T) (*xorm.Engine, error) {
	return db.defaultDB()
}
