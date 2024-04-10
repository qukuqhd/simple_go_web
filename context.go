package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Context struct { //context结构体，表示当前请求的上下文
	W http.ResponseWriter
	R *http.Request
}

func (c *Context) ReadJson(obj interface{}) error { //读取请求体的json数据，反序列化处理
	body, err := io.ReadAll(c.R.Body) //读取请求体数据
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, c.R)
	if err != nil {
		return err
	}

	return nil
}
func (c *Context) WriteJson(status_code int, resp interface{}) error { //对任意的一个对象进行序列化然后写入响应

	c.W.WriteHeader(status_code)          //设置响应状态码
	resp_json, err1 := json.Marshal(resp) //序列化响应数据
	if err1 != nil {
		return err1
	}
	_, err2 := c.W.Write(resp_json) //写入响应数据
	if err2 != nil {
		return err2
	}
	return nil
}
func (c *Context) WriteJsonOK(resp interface{}) error {
	return c.WriteJson(http.StatusOK, resp)
}
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{W: w, R: r}
}
