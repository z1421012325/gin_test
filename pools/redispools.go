package pools


import (
	"github.com/garyburd/redigo/redis"
	"time"
)


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

// redis.Values value 是返回多个数据 返回的是string 不是byte
//n,err := redis.String(pools.RedisPools.Get().Do("set","ppp","13187211252","ex","300"))
// redis.string 返回ok(如果成功)0和err
//if err != nil {
//	panic(err)
//}
//for _,value := range n{
//	fmt.Println(value)
//}
//n,_ := redis.String(pools.RedisPools.Get().Do("get","ppp"))
//if len(n) == 0 {
//	fmt.Println("没有找到!!")
//}