package logs


import (
	"log"
)



func main(){
	// 日志格式化输出
	// 2019/09/14 18:59:27 main.go:20: error
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)

}