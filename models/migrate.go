package models

import (
	"log"
	"mygo/gin_test/pools"
	"os"
	"strconv"
	"strings"
)

// 默认表名为 `users`
// 但是 db.SingularTable(true) 可以是否增加s后缀
// 使用结构体 struct 创建数据库表和字段
//type User struct{
//	// 如果需要使用结构体创建 并使用grom的model 默认为ID组件(帮你创建主键 自增长?)  如果自己则需要字段ID默认为组件 小写不行??
//	// 所有字段在数据库表中都是小写
//	//gorm.Model
//	ID int
//	Username string 	`gorm:"column:name;index;not null;unique"`
//	Passwoed string 	`gorm:"column:pswd;size:233;not null"`
//	Address  string		`gorm:"column:addr"`
//
//
//	// 经过测试 结构体创建模型表 其中的字段必须为大写开头 小写不会写入  驼峰命名会在两个大写之间使用下划线(都小写)做字段名
//	//CeShi string
//	//zaiciCeShi string
//	//hhhh string
//
//}

// 将user模型表的表名改为 返回值为在数据库创建的表名
// 如果不使用这个方法 根据在创建连接池之前进行的操作 db.SingularTable(true) 是否小写并增加s后缀
// 创建的表名是结构体小写 不加任何后缀
//func (User) TableName()string{
//	return "test_user"
//}


func Migrate(){

	// 检查环境变量中GIN_MODE是否为debug 不是说明是生成模式 直接返回
	var dvl string
	dvl = strings.ToLower("debug")
	if os.Getenv("GIN_MODE") != dvl{
		log.Println("生产环境!!!")
		return
	}

	// 自动迁移模式 不需要下面的方法 						不会删除?? 只会增加??
	pools.DB.AutoMigrate(&User{})
	// 根据结构体struct的方法创建的TableName 来操作表名
	//pools.DB.CreateTable(&User{})
	// 如果直接使用 db.table 指定表名 则上面的方法对这个操作不起作用
	//DB.Table("test_user").CreateTable(&User{})
	//fmt.Println("database table 创建完成!")




	// 如果存在则删除表
	//pools.DB.DropTableIfExists(&User{})

	// 寻找where 条件符合的 在user模型创建的表中删除 Delete
	//pools.DB.Where("name = ?","xiaoming").Delete(&User{})

	// sql语句   select * from user;
	//var u []User			// 创建一个能接受返回数据的容器,默认为该模型创建的切片对象
	//pools.DB.Model(&User{}).Find(&u) 	// 全部 find      first()  单个


	// 添加数据
	i := 0
	for i=0;i<50;i++{
		u:= User{
			Username: "xiaohong" + strconv.Itoa(i),
			Passwoed: "zymssb" + strconv.Itoa(i),
			Address:  "127.0.0." + strconv.Itoa(i),
		}
		pools.DB.Create(&u)			// 必须使用引用
	}


}