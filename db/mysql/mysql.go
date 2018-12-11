package mysql

import (
	"fmt"

	"github.com/s3dteam/go-toolkit/db/mysqldao"
)

// DB mysql DB
var DB *mysqldao.RdsService

// Init init mysql db with options
func Init(options *mysqldao.MysqlConifg) {
	var err error
	dbconfig := mysqldao.MysqlConifg{
		Hostname:    options.Hostname,
		Port:        options.Port,
		User:        options.User,
		Password:    options.Password,
		DbName:      options.DbName,
		TablePrefix: options.TablePrefix,
		Debug:       options.Debug,
	}

	newrdsdb, err := mysqldao.NewRdsService(dbconfig)
	if err != nil {
		panic(fmt.Sprintf("init mysql db err: %v", err))
	}
	DB = newrdsdb

	var tables []interface{}

	// create tables if not exists
	DB.RegistTables(tables)
}
