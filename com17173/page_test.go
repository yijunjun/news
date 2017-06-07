/*go**************************************************************************
 File            : list.go
 Subsystem       : com17173
 Author          : yijunjun
 Date&Time       : 2017-06-02
 Description     : 17173-详情页-测试
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

func test_page(t *testing.T, fun func(string) (*TPage, error), target string) {
	page, err := fun(target)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(page.Id, page.Author, page.Source, page.Date, page.HitsCount, page.CommentCount)
}

func TestLOLPage(t *testing.T) {
	test_page(t, LOLPage, "http://lol.17173.com/news/02212017/160148806.shtml")
}

func TestOWPage(t *testing.T) {
	test_page(t, OWPage, "http://ow.17173.com/content/02072017/110817325.shtml")
}

func TestMEPage(t *testing.T) {
	test_page(t, MEPage, "http://news.17173.com/z/pvp/content/06062017/105422411.shtml")
}

func TestDOTA2Page(t *testing.T) {
	test_page(t, DOTA2Page, "http://dota2.17173.com/news/04292017/150138726.shtml")
}

func TestCSGOPage(t *testing.T) {
	test_page(t, CSGOPage, "http://csgo.17173.com/content/05312017/114851014.shtml")
}

func TestBBSPage(t *testing.T) {
	test_page(t, BBSPage, "http://bbs.17173.com/thread-10310586-1-1.html")
}
