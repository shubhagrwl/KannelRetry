package controllers

import (
	dataservices "github.com/KannelRetry/dataservices"
)

// DBClient - We'll need an instance of the Mysql client
var (
	DataService dataservices.IMysqlClient
)
