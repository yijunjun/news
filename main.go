package main

import (
	"fmt"
	"html"
	"net/url"
	"strings"

	_ "github.com/yijunjun/news/com15w"
	_ "github.com/yijunjun/news/com17173"
	_ "github.com/yijunjun/news/com178"
	_ "github.com/yijunjun/news/duowan"
)

func loc(info string) {
	info = html.UnescapeString(info)

	anchor_flag := "作者："
	anchor_flag_size := len(anchor_flag)
	fmt.Println("anchor_flag_size:", anchor_flag_size)

	source_flag := "来源："
	source_flag_size := len(source_flag)
	fmt.Println("source_flag_size:", source_flag_size)

	date_flag := "发布时间："
	date_flag_size := len(date_flag)
	fmt.Println("date_flag_size:", date_flag_size)

	anchor_flag_pos := strings.Index(info, anchor_flag)
	if anchor_flag_pos == -1 {
		fmt.Println("anchor_flag_pos == -1")
		return
	}

	source_flag_pos := strings.Index(info, source_flag)
	if source_flag_pos == -1 {
		fmt.Println("source_flag_pos == -1")
		return
	}

	fmt.Println("anchor:", strings.TrimSpace(info[anchor_flag_pos+anchor_flag_size:source_flag_pos]))

	date_flag_pos := strings.Index(info, date_flag)
	if date_flag_pos == -1 {
		fmt.Println("date_flag_pos == -1")
		return
	}

	fmt.Println("source:", strings.TrimSpace(info[source_flag_pos+source_flag_size:date_flag_pos]))

	fmt.Println("date:", strings.TrimSpace(info[date_flag_pos+date_flag_size:]))
}

func main() {
	loc("作者：三杯空城丶小黑&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp来源：玩加赛事&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp发布时间：2017-05-26 18:43:52")
	return

	u, err := url.Parse("http://lol.178.com/201705/289795432613.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(u.Path)
	fmt.Println(strings.Split(u.Path, "/"))
}
