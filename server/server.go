package server

import (
	"github.com/gin-gonic/gin"
	"mygo/gin_test/service/v1"
	"os"

)


func GOServer(){
	server := gin.Default()

	// 所有路由都需要经过这个中间件
	//server.Use(middlerware.Session())
	//server.Use(middlerware.Cors())
	//server.Use(middlerware.CurrentUser())

	ApiOne := server.Group("/api/v1")
	{
		ApiOne.GET("ping",v1.Ping)


		// 将查找用户信息的api使用中间件 防止其他恶意用户调用
		ApiOne.Use()
		{
			ApiOne.GET("userinfo/:user_id",v1.GetUserinfo)
		}

	}



	addr := os.Getenv("SERVER_PORT")
	server.Run(addr)

}


