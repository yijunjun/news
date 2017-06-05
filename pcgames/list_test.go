/*go**************************************************************************
 File            : list_test.go
 Subsystem       : pcgames
 Author          : yijunjun
 Date&Time       : 2017-06-05
 Description     : pcgames-列表页-测试
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package pcgames

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
	run(t, LOLList, "http://wangyou.pcgames.com.cn/zhuanti/lol/news/")
}

func TestDOTA2List(t *testing.T) {
	run(t, DOTA2List, "http://fight.pcgames.com.cn/dota2/news/")
}

func TestCSGOList(t *testing.T) {
	run(t, CSGOList, "http://fight.pcgames.com.cn/cs/news/")
}
