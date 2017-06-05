/*go**************************************************************************
 File            : list_test.go
 Subsystem       : dadianjing
 Author          : yijunjun
 Date&Time       : 2017-06-05
 Description     : dadianjing-列表页-测试
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package dadianjing

import (
	"testing"

	. "github.com/yijunjun/news/model"
)

func run(t *testing.T, handler func(string) (*List, error), url string) {
	list, err := handler(url)
	if err != nil {
		t.Error(err.Error())
		return
	}

	for _, info := range list.Infos {
		t.Log(info.Url, info.Title, info.ImgSrc, info.Date)
	}

	for _, u := range list.Urls {
		t.Log(u)
	}
}

func TestLOLList(t *testing.T) {
	run(t, LOLList, "http://www.dadianjing.cn/index.php?m=Index&a=xhrList&cid=5&page=1")
}

func TestOWList(t *testing.T) {
	run(t, OWList, "http://www.dadianjing.cn/index.php?m=Index&a=xhrList&cid=6&page=1")
}

func TestCSGOList(t *testing.T) {
	run(t, CSGOList, "http://www.dadianjing.cn/index.php?m=Index&a=xhrList&cid=26&page=1")
}
