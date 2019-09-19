package v1

import (
	"github.com/gin-gonic/gin"
	"mygo/gin_test/models"
	"mygo/gin_test/serializer"
	"net/http"
)


type pro struct {
	Name string `form:"name" binding:"required"`
	Password  string `form:"password" binding:"required"`
}










// 确定服务器生死
func Ping(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"msg":"ok",
	})
}






func GetUserinfo(c *gin.Context){

	userid := c.Param("user_id")
	if userid == ""{
		c.JSON(
			http.StatusOK,
			serializer.Response{
				Status: 0,
				Data:   nil,
				Msg:    "",
				Error:  "user_id not in",
			}, )
		return
	}

	var u models.User
	u,_ = models.UserInfo(userid)

	c.JSON(
		http.StatusOK,
		serializer.Response{
			Status: 0,
			Data:   u,
			Msg:    "success",
			Error:  "",
		})
}




















