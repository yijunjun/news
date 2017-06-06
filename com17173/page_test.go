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

import "testing"

/*
func TestLOLPage(t *testing.T) {
	page, err := LOLPage("http://lol.17173.com/news/02212017/160148806.shtml")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(page.Id, page.Author, page.Source, page.Date, page.HitsCount, page.CommentCount)
}

func TestOWPage(t *testing.T) {
	// page, err := OWPage("http://bbs.17173.com/thread-10307183-1-1.html")
	page, err := OWPage("http://bbs.17173.com/thread-10310586-1-1.html")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(page.Id, page.Author, page.Source, page.Date, page.HitsCount, page.CommentCount)
}

func TestMEPage(t *testing.T) {
	page, err := MEPage("http://news.17173.com/z/pvp/content/06062017/105422411.shtml")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(page.Id, page.Author, page.Source, page.Date, page.HitsCount, page.CommentCount)
}
*/

func TestDOTA2Page(t *testing.T) {
	page, err := DOTA2Page("http://dota2.17173.com/news/04292017/150138726.shtml")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(page.Id, page.Author, page.Source, page.Date, page.HitsCount, page.CommentCount)
}

/*
func TestCSGOPage(t *testing.T) {
	page, err := CSGOPage("http://csgo.15w.com/yxzx/5606954825.html")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(page.Id, page.Author, page.Source, page.Date, page.HitsCount, page.CommentCount)
}
*/
