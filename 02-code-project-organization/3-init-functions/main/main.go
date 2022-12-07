package main

import (
	"fmt"

	"github.com/teivah/100-go-mistakes/02-code-project-organization/3-init-functions/redis"
)

// init会先从import开始，然后当前file可以多个init，顺序执行。
func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func main() {
	err := redis.Store("foo", "bar")
	_ = err
}
