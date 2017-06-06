/*go**************************************************************************
 File            : list.go
 Subsystem       : com178
 Author          : yijunjun
 Date&Time       : 2017-06-02
 Description     : 178-详情页-测试
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package com178

import "testing"

func TestLOLPage(t *testing.T) {
	page, err := LOLPage("http://lol.178.com/201705/289795432613.html")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(page.Id, page.Author, page.Source, page.Date, page.HitsCount, page.CommentCount)
}

func TestOWPage(t *testing.T) {
	page, err := OWPage("http://ow.178.com/201706/290311805555.html")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(page.Id, page.Author, page.Source, page.Date, page.HitsCount, page.CommentCount)
}

func TestMEPage(t *testing.T) {
	page, err := MEPage("http://shouyou.178.com/201611/32904103302.html")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(page.Id, page.Author, page.Source, page.Date, page.HitsCount, page.CommentCount)
}

func TestDOTA2Page(t *testing.T) {

	page, err := DOTA2Page("http://dota2.178.com/201705/289523076814.html")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(page.Id, page.Author, page.Source, page.Date, page.HitsCount, page.CommentCount)
}

func TestCSGOPage(t *testing.T) {
	page, err := CSGOPage("http://csgo.178.com/201706/290714969209.html")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(page.Id, page.Author, page.Source, page.Date, page.HitsCount, page.CommentCount)
}
