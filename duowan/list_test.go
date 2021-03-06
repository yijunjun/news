/*go**************************************************************************
 File            : list_test.go
 Subsystem       : duowan
 Author          : yijunjun
 Date&Time       : 2017-06-02
 Description     : duowan-列表页-测试
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package duowan

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
	test_list(t, LOLList, "http://lol.duowan.com/tag/307577396279.html")
}

func TestOWList(t *testing.T) {
	test_list(t, OWList, "http://ow.duowan.com/tag/309977280409.html")
}

func TestMEList(t *testing.T) {
	test_list(t, MEList, "http://wzry.duowan.com/tag/327319645493.html")
}

func TestDOTA2List(t *testing.T) {
	test_list(t, DOTA2List, "http://dota2.duowan.com/1302/m_225195792481.html")
}

func TestCSGOList(t *testing.T) {
	test_list(t, CSGOList, "http://csgo.duowan.com/tag/319717425248.html")
}
