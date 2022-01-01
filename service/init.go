package service

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
)

var DbEngine *xorm.Engine

// CONSTANT VALUES
const (
	DBTYPE = "mysql"
	SCHEMA = "%s:%s@tcp(mysql:3306)/%s?charset=utf8&parseTime=True&loc=Local"
)

var (
	Client   *gorm.DB
	username = "root"
	password = "mysql"
	dbName   = "gin"

	datasourceName = fmt.Sprintf(SCHEMA, username, password, dbName)
)

func init() {
	var err error
	// user:password@/db_name -> docker.compose.yml - mysql service
	Client, err = gorm.Open(DBTYPE, datasourceName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("database successfully configure")
}
