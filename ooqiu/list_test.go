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

func test_list(t *testing.T, handler func(string) (*TList, error), target string) {
	list, err := handler(target)
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
	test_list(t, LOLList, "http://www.ooqiu.com/lol/news/")
	test_list(t, LOLList, "http://www.ooqiu.com/lol/news/bg/")
}

func TestOWList(t *testing.T) {
	test_list(t, OWList, "http://ow.ooqiu.com/news/")
}

func TestMEList(t *testing.T) {
	test_list(t, MEList, "http://sy.ooqiu.com/wzry/rdnr/")
}

func TestDOTA2List(t *testing.T) {
	test_list(t, DOTA2List, "http://www.ooqiu.com/dota2/news/")
	test_list(t, DOTA2List, "http://www.ooqiu.com/dota2/news/saishi/")
}

func TestCSGOList(t *testing.T) {
	test_list(t, CSGOList, "http://www.ooqiu.com/csgo/news/")
	test_list(t, CSGOList, "http://www.ooqiu.com/csgo/news/saishi/")
}
