package main

import "net/http"

type HandlerBasedOnMap struct {
	//key 是 method+url
	//value 是对应的处理函数
	handlers map[string]func(ctx *Context)
}
type Routable interface {
	Route(method, pattern string, handelFunc func(ctx *Context))
}
type Handler interface { //接口组合
	ServeHTTP(c *Context)
	Routable
}

func (h *HandlerBasedOnMap) Route(method, pattern string, handelFunc func(ctx *Context)) {
	key := h.key(method, pattern)
	h.handlers[key] = handelFunc
}

// ServeHTTP implements http.Handler.
func (h *HandlerBasedOnMap) ServeHTTP(c *Context) {
	key := h.key(c.R.Method, c.R.URL.Path) //生成对应请求方法和路由的key

	if hander, ok := h.handlers[key]; ok { //已经注册的处理函数
		hander(NewContext(c.W, c.R)) //执行注册的方法
	} else { //未注册就404处理
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("404 not found"))
	}
	panic("unimplemented")
}
func (h *HandlerBasedOnMap) key(method, pattern string) string { //根据请求对象生成key
	return method + "#" + pattern
}
func NewHandler() Handler {
	return &HandlerBasedOnMap{
		handlers: make(map[string]func(ctx *Context), 128), //初始化map
	}
}
