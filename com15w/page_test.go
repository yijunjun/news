/*go**************************************************************************
 File            : list.go
 Subsystem       : com15w
 Author          : yijunjun
 Date&Time       : 2017-06-02
 Description     : 电竞头条-详情页-测试
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

func test_page(t *testing.T, handler func(string) (*TPage, error), target string) {
	page, err := handler(target)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(page.Id, page.Author, page.Source, page.Date, page.HitsCount, page.CommentCount)
}

func TestLOLPage(t *testing.T) {
	test_page(t, LOLPage, "http://lol.15w.com/ss/5606964485.html")
}

func TestOWPage(t *testing.T) {
	test_page(t, OWPage, "http://ow.15w.com/news/5606974205.html")
}

func TestMEPage(t *testing.T) {
	test_page(t, MEPage, "http://me.15w.com/dj/5605406665.html")
}

func TestDOTA2Page(t *testing.T) {
	test_page(t, DOTA2Page, "http://dota2.15w.com/guofuxinwen/5606947805.html")
}

func TestCSGOPage(t *testing.T) {
	test_page(t, CSGOPage, "http://csgo.15w.com/yxzx/5606954825.html")
}
