/*go**************************************************************************
 File            : list_test.go
 Subsystem       : tgbus
 Author          : yijunjun
 Date&Time       : 2017-06-05
 Description     : tgbus-列表页-测试
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package tgbus

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
	test_list(t, LOLList, "http://lol.tgbus.com/news/")
	test_list(t, LOLList, "http://lol.tgbus.com/news/bgzt/")
	test_list(t, LOLList, "http://lol.tgbus.com/news/ssxw/")
}

func TestOWList(t *testing.T) {
	test_list(t, OWList, "http://ow.tgbus.com/wfxw-4763/")
}

func TestMEList(t *testing.T) {
	test_list(t, MEList, "http://news.17173.com/z/pvp/list/zxwz.shtml")
}

func TestDOTA2List(t *testing.T) {
	test_list(t, DOTA2List, "http://dota2.tgbus.com/news/")
}

func TestCSGOList(t *testing.T) {
	test_list(t, CSGOList, "http://csgo.tgbus.com/new/")
}
