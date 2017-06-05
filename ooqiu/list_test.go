/*go**************************************************************************
 File            : list_test.go
 Subsystem       : ooqiu
 Author          : yijunjun
 Date&Time       : 2017-06-05
 Description     : ooqiu-列表页-测试
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package ooqiu

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
	run(t, LOLList, "http://www.ooqiu.com/lol/news/")
	run(t, LOLList, "http://www.ooqiu.com/lol/news/bg/")
}

func TestOWList(t *testing.T) {
	run(t, OWList, "http://ow.ooqiu.com/news/")
}

func TestMEList(t *testing.T) {
	run(t, MEList, "http://sy.ooqiu.com/wzry/rdnr/")
}

func TestDOTA2List(t *testing.T) {
	run(t, DOTA2List, "http://www.ooqiu.com/dota2/news/")
	run(t, DOTA2List, "http://www.ooqiu.com/dota2/news/saishi/")
}

func TestCSGOList(t *testing.T) {
	run(t, CSGOList, "http://www.ooqiu.com/csgo/news/")
	run(t, CSGOList, "http://www.ooqiu.com/csgo/news/saishi/")
}
