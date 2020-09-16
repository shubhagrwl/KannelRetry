package dataservices

import (
	"fmt"

	"github.com/KannelRetry/constants"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Connect opens a connection to the DB
func (pc *MysqlClient) Connect() {
	config := dbConfig()
	var err error
	//username:password@tcp(host:port)/dbname?
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config[constants.DBUSER], config[constants.DBPASS], config[constants.DBHOST],
		config[constants.DBPORT], config[constants.DBNAME])

	pc.DB, err = gorm.Open("mysql", mysqlInfo)

	if err != nil {
		log.Fatal("Fatal error encountered: ", err.Error())
	}
	err = pc.DB.DB().Ping()
	if err != nil {
		log.Fatal("Unable to ping database. Shutting down server.")
	}
	pc.DB.DB().SetMaxOpenConns(15)
	pc.DB.DB().SetMaxIdleConns(3)
	log.Info("Successfully connected to database.")
}

func dbConfig() map[string]string {
	var host, port, user, password, name string
	conf := make(map[string]string)

	host = viper.GetString(constants.KANNELRETRYDB + "." + constants.DBHOST)
	if host == "" {
		panic("DBHOST variable required but not set")
	}
	port = viper.GetString(constants.KANNELRETRYDB + "." + constants.DBPORT)
	if port == "" {
		panic("DBPORT variable required but not set")
	}
	user = viper.GetString(constants.KANNELRETRYDB + "." + constants.DBUSER)
	if user == "" {
		panic("DBUSER variable required but not set")
	}
	password = viper.GetString(constants.KANNELRETRYDB + "." + constants.DBPASS)
	if password == "" {
		panic("DBPASS variable required but not set")
	}
	name = viper.GetString(constants.KANNELRETRYDB + "." + constants.DBNAME)
	if name == "" {
		panic("DBNAME variable required but not set")
	}
	conf[constants.DBHOST] = host
	conf[constants.DBPORT] = port
	conf[constants.DBUSER] = user
	conf[constants.DBPASS] = password
	conf[constants.DBNAME] = name
	return conf
}

func (pc *MysqlClient) GetDBInstance() *gorm.DB {
	return pc.DB
}
