package dataservices

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MysqlClient stored reference to the DB
type MysqlClient struct {
	DB *gorm.DB
}

//IPostgresClient -
type IMysqlClient interface {
	Connect()

	GetDBInstance() *gorm.DB
}
