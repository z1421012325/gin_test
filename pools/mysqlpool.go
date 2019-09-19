package pools

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

	// 输入sql日志
	db.LogMode(true)

	//gorm查找struct名对应数据库中的表名的时候会默认把你的struct中的大写字母转换为小写并加上“s”，
	// 所以可以加上db.SingularTable(true) 让grom转义struct名字的时候不用加上s
	db.SingularTable(true)

	// 设置可重用连接的最大时间量 如果d<=0，则永远重用连接
	db.DB().SetConnMaxLifetime(time.Hour)

	//设置到数据库的最大打开连接数 如果n<=0，则不限制打开的连接数 默认值为0
	db.DB().SetMaxOpenConns(0)

	// 设置空闲中的最大连接数 默认最大空闲连接数当前为2 如果n<=0，则不保留空闲连接
	db.DB().SetMaxIdleConns(20)


	err = db.DB().Ping()
	if err != nil {
		log.Fatal(err)
	}

	DB = db


}

