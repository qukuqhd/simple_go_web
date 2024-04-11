package main

import (
	"strings"
)

type HandlerBasedOnTree struct {
	root node //根节点
}

func NewNode(path string) *node {
	return &node{
		path:     path,
		children: make([]*node, 5),
	}
}

type node struct {
	path     string  //当前的路径线索
	children []*node //子节点
	handler  HandlerFunc
}

func (tree *HandlerBasedOnTree) ServeHTTP(c *Context) {

}
func (tree *HandlerBasedOnTree) Route(method, pattern string, handelFunc HandlerFunc) {
	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")
	cur := tree.root
	for _, path := range paths {
		matchChild, ok := tree.matchChild(&cur, path)
		if ok { //存在节点就覆盖
			matchChild.handler = handelFunc //
		} else { //不存在就创建子节点插入
			tree.insertNode(&cur, paths, handelFunc) //创建新的节点树插入当前到的节点的子节点
			break
		}
	}
}
func (tree *HandlerBasedOnTree) matchChild(root *node, path string) (*node, bool) {
	var wildcardNode *node = nil
	for _, child := range root.children {
		if child.path == path && child.path != "*" { //是直接命中了具体的路由
			return child, true
		} else if child.path == "*" { //命中了通配符
			wildcardNode = child
		}
	}
	return wildcardNode, wildcardNode != nil
}

// 添加路由树，根据现在还有的路径来往下生成一颗偏向树到目的的节点上
func (tree *HandlerBasedOnTree) insertNode(root *node, paths []string, handler HandlerFunc) {
	cur := root
	for _, path := range paths { //从根节点开始遍历
		nn := NewNode(path)
		cur.children = append(cur.children, nn) //添加到子节点切片
		cur = nn                                //移动到子节点
	}
	cur.handler = handler //设置处理函数
}
