/*go**************************************************************************
 File            : list.go
 Subsystem       : com178
 Author          : yijunjun
 Date&Time       : 2017-06-02
 Description     : 178-列表页
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package com178

import (
	"errors"
	"strings"

	"github.com/PuerkitoBio/goquery"
	. "github.com/yijunjun/news/model"
)

func LOLList(list_url string) (*List, error) {
	// 列表页
	doc, err := goquery.NewDocument(list_url)
	if err != nil {
		return nil, err
	}

	art_list := doc.Find(".list-section-contents")

	list := &List{
		Infos: make([]PageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		a_node := s.Find("h2 a")

		// 详细页网址
		if href, has := a_node.Attr("href"); has {
			list.Infos[i].Url = href
		}

		// 标题
		if title, has := a_node.Attr("title"); has {
			list.Infos[i].Title = title
		}

		// 发布时间
		list.Infos[i].Date = s.ChildrenFiltered("h5").Text()
	})

	next_pages(list, doc)

	return list, nil
}

func OWList(list_url string) (*List, error) {
	// 列表页
	doc, err := goquery.NewDocument(list_url)
	if err != nil {
		return nil, err
	}

	art_list := doc.Find(".list-list dl")

	list := &List{
		Infos: make([]PageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		a_node := s.Find("dd .ds1 a")

		// 详细页网址
		if href, has := a_node.Attr("href"); has {
			list.Infos[i].Url = href
		}
		// 标题
		if title, has := a_node.Attr("title"); has {
			list.Infos[i].Title = title
		}

		// 封面
		a_img_node := s.Find("dt a img")
		if src, has := a_img_node.Attr("src"); has {
			list.Infos[i].ImgSrc = src
		}

		// 发布时间
		list.Infos[i].Date = strings.SplitN(
			s.Find("dd .ds2 span").Eq(1).Text(),
			"：",
			2,
		)[1]
	})

	next_pages(list, doc)

	return list, nil
}

func MEList(list_url string) (*List, error) {
	// 列表页
	doc, err := goquery.NewDocument(list_url)
	if err != nil {
		return nil, err
	}

	art_list := doc.Find(".list-section-contents")

	list := &List{
		Infos: make([]PageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		a_node := s.Find("h2 a")

		// 详细页网址
		if href, has := a_node.Attr("href"); has {
			list.Infos[i].Url = href
		}

		// 标题
		if title, has := a_node.Attr("title"); has {
			list.Infos[i].Title = title
		}

		// 发布时间
		list.Infos[i].Date = s.ChildrenFiltered("h5").Text()
	})

	next_pages(list, doc)

	return list, nil
}

func DOTA2List(list_url string) (*List, error) {
	// 列表页
	doc, err := goquery.NewDocument(list_url)
	if err != nil {
		return nil, err
	}

	art_list := doc.Find(".articlelist li")

	list := &List{
		Infos: make([]PageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		a_node := s.Find(".info a")

		// 详细页网址
		if href, has := a_node.Attr("href"); has {
			list.Infos[i].Url = href
		}

		// 标题
		list.Infos[i].Title = a_node.Text()

		// 封面
		img_node := s.Find(".pic img")
		if src, has := img_node.Attr("src"); has {
			list.Infos[i].ImgSrc = src
		}

		// 发布时间
		list.Infos[i].Date = s.Find(".time").Text()
	})

	next_pages(list, doc)

	return list, nil
}

func CSGOList(list_url string) (*List, error) {
	// 列表页
	doc, err := goquery.NewDocument(list_url)
	if err != nil {
		return nil, err
	}

	art_list := doc.Find(".conts li")

	list := &List{
		Infos: make([]PageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		a_node := s.ChildrenFiltered("a")

		// 详细页网址
		if href, has := a_node.Attr("href"); has {
			list.Infos[i].Url = href
		}
		// 标题
		if title, has := a_node.Attr("title"); has {
			list.Infos[i].Title = title
		}

		// 封面
		a_img_node := a_node.ChildrenFiltered("img")
		if src, has := a_img_node.Attr("src"); has {
			list.Infos[i].ImgSrc = src
		}

		// 发布时间
		list.Infos[i].Date = s.Find(".sp4").Text()
	})

	next_pages(list, doc)

	return list, nil
}

func next_pages(list *List, doc *goquery.Document) {
	url_list := doc.Find(".page a[class=cms_pages]")

	list.Urls = make([]string, url_list.Length())

	url_list.Each(func(i int, s *goquery.Selection) {
		list.Urls[i] = s.AttrOr("href", "")
	})
}

func NewList(list_url string) (*List, error) {
	var game_list = map[string]func(string) (*List, error){
		"http://lol.":     LOLList,
		"http://ow.":      OWList,
		"http://csgo.":    CSGOList,
		"http://shouyou.": MEList,
		"http://dota2.":   DOTA2List,
	}
	for prefix, handler := range game_list {
		if strings.HasPrefix(list_url, prefix) {
			return handler(list_url)
		}
	}
	return nil, errors.New("can not support " + list_url)
}
