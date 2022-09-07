package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

//当启动程序时，就初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲链接数
		MaxActive:   0,   // 表示和数据库的最大链接数， 0 表示没有限制
		IdleTimeout: 100, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化链接的代码， 链接哪个ip 的redis
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func Set(imgPath string) {
	//先从pool 取出一个链接
	conn := pool.Get()
	defer conn.Close()

	//同样的使用
	_, err := conn.Do("Set", "name", "汤姆猫~~")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	//取出
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	fmt.Println("r=", r)
	//如果我们要从pool 取出链接，一定保证链接池是没有关闭
	//defer pool.Close()

}
