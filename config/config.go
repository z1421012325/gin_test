package config

import (
	"github.com/joho/godotenv"
	"log"
)




//读取根目录配置env文件
func Init(){
	// 默认在当前
	err := godotenv.Load("./gintest.env")
	if err != nil {
		log.Fatal("读取环境配置文件出错!!!")
		//panic("读取环境配置文件出错!!!")
	}

}
