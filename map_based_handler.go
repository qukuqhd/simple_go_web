package main

import "net/http"

type HandlerBasedOnMap struct {
	//key 是 method+url
	//value 是对应的处理函数
	handlers map[string]func(ctx *Context)
}

// ServeHTTP implements http.Handler.
func (h HandlerBasedOnMap) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := h.key(request.Method, request.URL.Path) //生成对应请求方法和路由的key

	if hander, ok := h.handlers[key]; ok { //已经注册的处理函数
		hander(NewContext(writer, request)) //执行注册的方法
	} else { //未注册就404处理
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("404 not found"))
	}
	panic("unimplemented")
}
func (h *HandlerBasedOnMap) key(method, pattern string) string { //根据请求对象生成key
	return method + "#" + pattern
}
