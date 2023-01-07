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
	Password string
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
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlAddr := os.Getenv("MYSQL_Addr")
	mysqlDBBame := os.Getenv("MYSQL_DBNAME")

	mysqlInfo := &MySQLInfo{
		User:     mysqlUser,
		Password: mysqlPassword,
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
