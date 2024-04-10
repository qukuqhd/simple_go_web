package main

import (
	"net/http"
)

type Server interface { //web框架服务器需要实现的接口
	Route(method, pattern string, handelFunc func(ctx *Context)) //设置服务路由，请求命中路由时调用handelFunc函数
	//服务器启动
	Start(addr string) error
}
type sdkHttpServer struct { //基于sdk的http库的服务结构体

	Name    string
	handler HandlerBasedOnMap //方法类型路由和处理函数的映射
}

func (s *sdkHttpServer) Route(method, pattern string, handelFunc func(ctx *Context)) { //为sdkhttpserver结构体实现server接口的route方法
	//这里添加映射信息
	key := s.handler.key(method, pattern) //生成key
	s.handler.handlers[key] = handelFunc  //添加处理
}
func (s *sdkHttpServer) Start(addr string) error { //为sdkhttpserver结构体实现server接口的start方法
	http.Handle("/", s.handler) //具体把一系列的处理注册到http的handler

	return http.ListenAndServe(addr, nil) //启动服务监听注册好的路由
}
func NewHttpServer(server_name string) Server { //返回一个创建的实现server接口的对象（多态）
	return &sdkHttpServer{Name: server_name}
}
