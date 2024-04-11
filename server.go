package main

import (
	"net/http"
)

type Server interface { //web框架服务器需要实现的接口
	Routable //可以路由设置
	// Start handelFunc函数 服务器启动
	Start(addr string) error
}
type sdkHttpServer struct { //基于sdk的http库的服务结构体
	root    Filter
	Name    string
	handler Handler //方法类型路由和处理函数的映射
}

func (s *sdkHttpServer) Route(method, pattern string, handelFunc func(ctx *Context)) { //为sdkhttpserver结构体实现server接口的route方法
	//这里添加映射信息
	s.handler.Route(method, pattern, handelFunc) //委托添加路由
}
func (s *sdkHttpServer) Start(addr string) error { //为sdkhttpserver结构体实现server接口的start方法
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) { //注册根路由
		s.root(NewContext(writer, request))
	})

	return http.ListenAndServe(addr, nil) //启动服务监听注册好的路由
}
func NewHttpServer(serverName string, builders ...FilterBuilder) Server { //返回一个创建的实现server接口的对象（多态）
	handler := NewHandler()
	var rootFilter = handler.ServeHTTP
	for i := len(builders) - 1; i >= 0; i-- { //从后往前
		b := builders[i]
		rootFilter = b(rootFilter)
	}
	return &sdkHttpServer{
		root:    rootFilter,
		Name:    serverName,
		handler: NewHandler(),
	}
}
