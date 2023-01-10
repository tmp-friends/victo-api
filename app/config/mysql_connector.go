package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const driverName = "mysql"

type MySQLConnector struct {
	Conn *sql.DB
}

func NewMySQLConnector() *MySQLConnector {
	conf := LoadConfig()
	dsn := createDSN(*conf.MySQLInfo)

	conn, err := sql.Open(driverName, dsn)
	if err != nil {
		panic(err)
	}

	return &MySQLConnector{
		Conn: conn,
	}
}

func createDSN(mysqlInfo MySQLInfo) string {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true&parseTime=true",
		mysqlInfo.User,
		mysqlInfo.Password,
		mysqlInfo.Addr,
		mysqlInfo.DBName,
	)

	return dataSourceName
}
