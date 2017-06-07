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
	test_page(t, LOLPage, "http://lol.178.com/201705/289795432613.html")
}

func TestOWPage(t *testing.T) {
	test_page(t, OWPage, "http://ow.178.com/201706/290311805555.html")
}

func TestMEPage(t *testing.T) {
	test_page(t, MEPage, "http://shouyou.178.com/201611/32904103302.html")
}

func TestDOTA2Page(t *testing.T) {
	test_page(t, DOTA2Page, "http://dota2.178.com/201705/289523076814.html")
}

func TestCSGOPage(t *testing.T) {
	test_page(t, CSGOPage, "http://csgo.178.com/201706/290714969209.html")
}
