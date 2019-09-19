package main

import (
	"mygo/gin_test/config"
	"mygo/gin_test/pools"
	"mygo/gin_test/server"
)

func main(){
	// 可放置在server.go中
	// 加载配置文件
	config.Init()

	// 加载redis,mysql连接池
	pools.Init()

	// 一般情况下是不会使用gorm来创建数据库表的,生成环境都是自带的 预设的
	//models.Migrate()

	// 启动服务器 端口在配置文件 env 中配置
	server.GOServer()





}