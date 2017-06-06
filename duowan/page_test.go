/*go**************************************************************************
 File            : list.go
 Subsystem       : duowan
 Author          : yijunjun
 Date&Time       : 2017-06-02
 Description     : duowan-详情页-测试
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package duowan

import "testing"

func TestLOLPage(t *testing.T) {
	page, err := LOLPage("http://lol.duowan.com/1706/360684186383.html")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(page.Id, page.Author, page.Source, page.Date, page.HitsCount, page.CommentCount)
}

func TestOWPage(t *testing.T) {
	page, err := LOLPage("http://ow.duowan.com/1703/354707581307.html")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(page.Id, page.Author, page.Source, page.Date, page.HitsCount, page.CommentCount)
}

func TestMEPage(t *testing.T) {
	page, err := MEPage("http://wzry.duowan.com/1706/360413827429.html")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(page.Id, page.Author, page.Source, page.Date, page.HitsCount, page.CommentCount)
}

func TestDOTA2Page(t *testing.T) {
	page, err := DOTA2Page("http://dota2.duowan.com/1604/325095526695.html")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(page.Id, page.Author, page.Source, page.Date, page.HitsCount, page.CommentCount)
}

func TestCSGOPage(t *testing.T) {
	page, err := CSGOPage("http://csgo.duowan.com/1704/356974010832.html")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(page.Id, page.Author, page.Source, page.Date, page.HitsCount, page.CommentCount)
}
