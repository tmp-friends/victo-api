package config

type appConfig struct {
	HTTPInfo *HTTPInfo
	// MySQLInfo *MySQLInfo
}

type HTTPInfo struct {
	Addr string
}

// type MySQLInfo struct {
// 	User     string
// 	Password string
// 	Port     string
// 	DBName   string
// }

func LoadConfig() *appConfig {
	// HTTPInfo
	// addr := ":" + os.Getenv("ADDR")
	addr := ":3001"

	httpInfo := &HTTPInfo{
		Addr: addr,
	}

	// MySQLInfo
	//mysqlUser := os.Getenv("MYSQL_USER")
	//mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	//mysqlPort := os.Getenv("MYSQL_PORT")
	//mysqlDBBame := os.Getenv("MYSQL_DBNAME")

	//mysqlInfo := &MySQLInfo{
	//	User:     mysqlUser,
	//	Password: mysqlPassword,
	//	Port:     mysqlPort,
	//	DBName:   mysqlDBBame,
	//}

	// appConfig
	config := appConfig{
		HTTPInfo: httpInfo,
		// MySQLInfo: mysqlInfo,
	}

	return &config
}
