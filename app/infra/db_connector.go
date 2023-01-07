package infra

import (
	"database/sql"
	"fmt"

	"github.com/tmp-friends/victo-api/app/config"
)

const driverName = "mysql"

type DBConnector struct {
	Conn *sql.DB
}

func NewDBConnector() *DBConnector {
	conf := config.LoadConfig()
	dsn := createDSN(*conf.MySQLInfo)

	conn, err := sql.Open(driverName, dsn)
	if err != nil {
		panic(err)
	}

	return &DBConnector{
		Conn: conn,
	}
}

func createDSN(mysqlInfo config.MySQLInfo) string {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		mysqlInfo.User,
		mysqlInfo.Password,
		mysqlInfo.Addr,
		mysqlInfo.DBName,
	)

	return dataSourceName
}
