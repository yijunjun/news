/*go**************************************************************************
 File            : list_test.go
 Subsystem       : com178
 Author          : yijunjun
 Date&Time       : 2017-06-02
 Description     : 178-列表页-测试
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package com178

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
	run(t, LOLList, "http://lol.178.com/list/info_5.html")
}

func TestOWList(t *testing.T) {
	run(t, OWList, "http://ow.178.com/list/230422432745_3.html")
}

func TestMEList(t *testing.T) {
	run(t, MEList, "http://shouyou.178.com/list/230864526697_7.html")
}

func TestDOTA2List(t *testing.T) {
	run(t, DOTA2List, "http://dota2.178.com/list/news_10.html")
}

func TestCSGOList(t *testing.T) {
	run(t, CSGOList, "http://csgo.178.com/list/252011760339.html")
}
