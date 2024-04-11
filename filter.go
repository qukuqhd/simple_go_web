package main

import (
	"fmt"
	"time"
)

type HandlerFunc func(c *Context)
type Filter func(c *Context)
type FilterBuilder func(next Filter) Filter

var _ FilterBuilder = MetricsFilterBuilder

func MetricsFilterBuilder(next Filter) Filter {
	return func(c *Context) {
		start := time.Now().Nanosecond()
		next(c)
		end := time.Now().Nanosecond()
		fmt.Printf("cost:%d\n", end-start)
	}
}
