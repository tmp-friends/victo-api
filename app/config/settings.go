package config

import (
	"os"
)

type appConfig struct {
	HTTPInfo  *HTTPInfo
	MySQLInfo *MySQLInfo
}

type HTTPInfo struct {
	Addr string
}

type MySQLInfo struct {
	User     string
	Pass string
	Addr     string
	DBName   string
}

func LoadConfig() *appConfig {
	// HTTPInfo
	addr := ":" + os.Getenv("ADDR")

	httpInfo := &HTTPInfo{
		Addr: addr,
	}

	// MySQLInfo
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPass := os.Getenv("MYSQL_PASS")
	mysqlAddr := os.Getenv("MYSQL_ADDR")
	mysqlDBBame := os.Getenv("MYSQL_DBNAME")

	mysqlInfo := &MySQLInfo{
		User:     mysqlUser,
		Pass: mysqlPass,
		Addr:     mysqlAddr,
		DBName:   mysqlDBBame,
	}

	// appConfig
	config := appConfig{
		HTTPInfo:  httpInfo,
		MySQLInfo: mysqlInfo,
	}

	return &config
}
