package main

import (
	"fmt"

	_ "github.com/yijunjun/news/com15w"
	_ "github.com/yijunjun/news/com17173"
	_ "github.com/yijunjun/news/com178"
	_ "github.com/yijunjun/news/dadianjing"
	_ "github.com/yijunjun/news/duowan"
	_ "github.com/yijunjun/news/ooqiu"
	_ "github.com/yijunjun/news/pcgames"
	_ "github.com/yijunjun/news/tgbus"
)

func main() {
	fmt.Println("hello,world!")
}
