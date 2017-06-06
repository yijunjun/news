/*go**************************************************************************
 File            : list_test.go
 Subsystem       : com15w
 Author          : yijunjun
 Date&Time       : 2017-06-02
 Description     : 电竞头条-列表页-测试
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package com15w

import (
	"testing"

	. "github.com/yijunjun/news/model"
)

func run(t *testing.T, handler func(string) (*TList, error), url string) {
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
	run(t, LOLList, "http://lol.15w.com/zx/index_535.html")
}

func TestOWList(t *testing.T) {
	run(t, OWList, "http://ow.15w.com/xw/index_4.html")
}

func TestMEList(t *testing.T) {
	run(t, MEList, "http://me.15w.com/wzry/zx/index_1.html")
}

func TestDOTA2List(t *testing.T) {
	run(t, DOTA2List, "http://dota2.15w.com/news/index_3.html")
}

func TestCSGOList(t *testing.T) {
	run(t, CSGOList, "http://csgo.15w.com/news/index_4.html")
}
