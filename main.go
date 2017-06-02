package main

import (
	"fmt"
	"net/url"
	"path"
)

func main() {
	u, err := url.Parse("http://lol.duowan.com/tag/307577396279.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(u.Scheme+"//"+u.Host, u.Path, path.Dir(u.Path))
}
