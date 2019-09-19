package pools

import (
	"os"
	"fmt"
)


// 连接池  读取配置文件
func Init(){
	// redis
	addr := os.Getenv("REDIS_ADDR")
	redispools(addr)


	//mysql
	username := os.Getenv("MYSQL_NAME")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	mysql_db := os.Getenv("MYSQL_DB")

	// uri格式  "user:password@/tcp(host:port)dbname?charset=utf8&parseTime=True&loc=Local"
	connArgs := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username,password, host, port, mysql_db )

	mysqlPool(connArgs)


}

