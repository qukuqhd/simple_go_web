package main

import "net/http"

type Server interface { //web框架服务器需要实现的接口
	Route(pattern string, handelFunc http.HandlerFunc) //设置服务路由，请求命中路由时调用handelFunc函数
	//服务器启动
	Start(addr string) error
}
