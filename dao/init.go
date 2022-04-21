package dao

import (
	"errors"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	DB_TYPE = iota
	DB_HOST
	DB_PORT
	DB_NAME
	DB_USER
	DB_PASS
	DB_TZ
)

type DB struct {
	VurixDmsDBConn *gorm.DB
}

type Filter struct {
	Query interface{}
	Args  []interface{}
}

type DefaultParam struct {
	Vendors []string `json:"vendors"`
}

func ConnectDB(config []string) (*gorm.DB, error) {
	dbtype := config[DB_TYPE]
	if dbtype == "mysql" {
		return ConnectMySQL(config)
	}
	if dbtype == "postgres" {
		return ConnectPostgres(config)
	}
	return nil, errors.New("Unknown DB Type (" + dbtype + ")")
}

func ConnectMySQL(config []string) (*gorm.DB, error) {
	dbtype := config[DB_TYPE]
	// user:password@tcp(host:port)/dbname?charset=utf8&parseTime=True&loc=Local
	connstring := config[DB_USER] + ":" + config[DB_PASS] + "@tcp(" + config[DB_HOST] + ":" + config[DB_PORT] + ")/" + config[DB_NAME] + "?charset=utf8&parseTime=True&loc=Local"
	conn, err := gorm.Open(dbtype, connstring)
	if err != nil {
		log.Println("DB Connection Fail : ", config, ", Error : ", err)
		return nil, err
	}
	//conn = conn.Debug()
	return conn, nil
}

func ConnectPostgres(config []string) (*gorm.DB, error) {
	var conn *gorm.DB
	var err error
	var connstring string

	dbtype := config[DB_TYPE]
	// "host=myhost port=myport user=gorm dbname=gorm password=mypassword"
	connstring = "host=" + config[DB_HOST] + " port=" + config[DB_PORT] + " user=" + config[DB_USER] + " dbname=" + config[DB_NAME] + " password=" + config[DB_PASS] + " TimeZone=Asia/Seoul"
	conn, err = gorm.Open(dbtype, connstring)
	if err != nil {
		if strings.Contains(err.Error(), "SSL is not enabled on the server") {
			connstring = connstring + " sslmode=disable"
		}
		conn, err = gorm.Open(dbtype, connstring)
		if err != nil {
			log.Println("DB Connection Fail : ", config, ", Error : ", err)
			return nil, err
		}
	}
	//conn = conn.Debug()
	return conn, nil
}
