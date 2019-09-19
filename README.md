### web框架上手



> 框架为gin 

--------------


```
读取配置文件 默认读取当前目录env文件
godotenv.Load
```

```
redis连接池 os.getenv 得到加载的环境变量
addr := os.Getenv("REDIS_ADDR")
redispools(addr)


// redis连接池实例
var RedisPools *redis.Pool

// Redis连接池
func redispools(addr string){

	r := &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp",addr)
		},

		// 用来检测是否链接 ping
		//TestOnBorrow: func(c redis.Conn, t time.Time) error {},

		// 最大空闲
		MaxIdle:         50,
		//当为零时，池中的连接数没有限制。
		MaxActive:       0,
		// 超时时间
		IdleTimeout:     240 * time.Second,
		// 当超过最大连接数时直接返回错误 默认false
		Wait:            false,
		// 最大等待时间 默认为0 无限制
		MaxConnLifetime: 0,

	}

	RedisPools = r

}
```


```
mysql连接池
username := os.Getenv("MYSQL_NAME")
password := os.Getenv("MYSQL_PASSWORD")
host := os.Getenv("MYSQL_HOST")
port := os.Getenv("MYSQL_PORT")
mysql_db := os.Getenv("MYSQL_DB")

// uri格式  "user:password@/tcp(host:port)dbname?charset=utf8&parseTime=True&loc=Local"
connArgs := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
    username,password, host, port, mysql_db )

mysqlPool(connArgs)

-----------------------------
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

// mysql连接池
var DB *gorm.DB

func mysqlPool(connArgs string){

	db,err := gorm.Open("mysql",connArgs)
	if err != nil {
		log.Fatal(err)
	}

	//gorm查找struct名对应数据库中的表名的时候会默认把你的struct中的大写字母转换为小写并加上“s”，
	// 所以可以加上db.SingularTable(true) 让grom转义struct名字的时候不用加上s
	db.SingularTable(true)

	// 设置可重用连接的最大时间量 如果d<=0，则永远重用连接
	db.DB().SetConnMaxLifetime(time.Hour)
jie
	//设置到数据库的最大打开连接数 如果n<=0，则不限制打开的连接数 默认值为0
	db.DB().SetMaxOpenConns(0)

	// 设置空闲中的最大连接数 默认最大空闲连接数当前为2 如果n<=0，则不保留空闲连接
	db.DB().SetMaxIdleConns(20)

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	DB = db


}

```