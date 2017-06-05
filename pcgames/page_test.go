/*go**************************************************************************
 File            : list.go
 Subsystem       : pcgames
 Author          : yijunjun
 Date&Time       : 2017-06-05
 Description     : pcgames-详情页-测试
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package pcgames

import "testing"

func TestLOLPage(t *testing.T) {
	page, err := LOLPage("http://lol.15w.com/ss/5606964485.html")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(page.Id, page.Anthor, page.Source, page.Date, page.HitsCount, page.CommentCount)
}

func TestOWPage(t *testing.T) {
	page, err := LOLPage("http://ow.15w.com/ss/5606966705.html")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(page.Id, page.Anthor, page.Source, page.Date, page.HitsCount, page.CommentCount)
}

func TestCSGOPage(t *testing.T) {
	page, err := CSGOPage("http://csgo.15w.com/yxzx/5606954825.html")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(page.Id, page.Anthor, page.Source, page.Date, page.HitsCount, page.CommentCount)
}

func TestMEPage(t *testing.T) {

	page, err := MEPage("http://me.15w.com/dj/5605406665.html")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(page.Id, page.Anthor, page.Source, page.Date, page.HitsCount, page.CommentCount)
}

func TestDOTA2Page(t *testing.T) {

	page, err := DOTA2Page("http://dota2.15w.com/guofuxinwen/5606947805.html")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Log(page.Id, page.Anthor, page.Source, page.Date, page.HitsCount, page.CommentCount)
}
