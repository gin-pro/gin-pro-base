package sqlx

import (
	"xorm.io/xorm"
)

func DefaultDB(driverName, dataSourceName string) (*xorm.Engine, error) {
	return xorm.NewEngine(driverName, dataSourceName)
}
