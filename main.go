package main

func main() {
	server := NewHttpServer("server1") //创建服务

	server.Route("GET", "/", func(ctx *Context) {
		ctx.WriteJsonOK("hello world")
	}) //注册服务的路由信息
	server.Start(":8080") //启动服务
}
