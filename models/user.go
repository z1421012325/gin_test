package models

import (
	"mygo/gin_test/pools"
)

// 根据数据库中的表生成的结构体
type User struct{
	ID int
	Username string 	`gorm:"column:name;index;not null;unique"`
	Passwoed string 	`gorm:"column:pswd;size:233;not null"`
	Address  string		`gorm:"column:addr"`
}

// user 表名
func (u *User) TableName() string{
	return "test_user"
}



// ---------------------------查询 --------------------------------

// 查询用户详情信息
func UserInfo (id interface{}) (User,error) {

	var u User

	err := pools.DB.Model(&User{}).Where("id = ?",id).First(&u).Error
	if err != nil {
		return u,err
	}

	return u,nil
}






// ---------------------------增加 --------------------------------

// 增加用户
func CreateUser (u *User) {
	pools.DB.Create(u)
}











// ---------------------------修改 --------------------------------
// ---------------------------删除 --------------------------------
