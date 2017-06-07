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

import (
	"testing"

	. "github.com/yijunjun/news/model"
)

func test_page(t *testing.T, handler func(string) (*TPage, error), target string) {
	page, err := handler(target)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(page.Id, page.Author, page.Source, page.Date, page.HitsCount, page.CommentCount)
}

func TestLOLPage(t *testing.T) {
	test_page(t, LOLPage, "http://lol.duowan.com/1706/360684186383.html")
}

func TestOWPage(t *testing.T) {
	test_page(t, OWPage, "http://ow.duowan.com/1703/354707581307.html")
}

func TestMEPage(t *testing.T) {
	test_page(t, MEPage, "http://wzry.duowan.com/1706/360413827429.html")
}

func TestDOTA2Page(t *testing.T) {
	test_page(t, DOTA2Page, "http://dota2.duowan.com/1604/325095526695.html")
}

func TestCSGOPage(t *testing.T) {
	test_page(t, CSGOPage, "http://csgo.duowan.com/1704/356974010832.html")
}
