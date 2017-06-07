/*go**************************************************************************
 File            : list_test.go
 Subsystem       : com17173
 Author          : yijunjun
 Date&Time       : 2017-06-02
 Description     : 17173-列表页-测试
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package com17173

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
	test_list(t, LOLList, "http://lol.17173.com/list/zdbg/")
	test_list(t, LOLList, "http://lol.17173.com/list/zixun.shtml")
}

func TestOWList(t *testing.T) {
	test_list(t, OWList, "http://ow.17173.com/news/")
}

func TestMEList(t *testing.T) {
	test_list(t, MEList, "http://news.17173.com/z/pvp/list/zxwz.shtml")
}

func TestDOTA2List(t *testing.T) {
	test_list(t, DOTA2List, "http://dota2.17173.com/news/")
}

func TestCSGOList(t *testing.T) {
	test_list(t, CSGOList, "http://csgo.17173.com/z/")
}
