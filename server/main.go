package main

//启动接口

import (
	"txc-njl/server/routers"
)

func main() {
	//Default返回一个默认的路由引擎
	//DB.InitDB()
	r := routers.SetRouters()
	r.Run() // 可以从本地的127.0.0.1:8081访问，不填的话，默认是8080端口

}

//请求拦截
