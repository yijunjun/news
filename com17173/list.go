/*go**************************************************************************
 File            : list.go
 Subsystem       : com17173
 Author          : yijunjun
 Date&Time       : 2017-06-02
 Description     : 17173-列表页
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package com17173

import (
	"errors"
	"strings"

	"github.com/PuerkitoBio/goquery"
	. "github.com/yijunjun/news/model"

	"net/http"

	"github.com/axgle/mahonia"
)

func LOLList(list_url string) (*List, error) {
	// 列表页
	doc, err := goquery.NewDocument(list_url)
	if err != nil {
		return nil, err
	}

	art_list := doc.Find(".comm-list .art-item")

	list := &List{
		Infos: make([]PageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		a_node := s.Find("a")

		// 详细页网址
		if href, has := a_node.Attr("href"); has {
			list.Infos[i].Url = href
		}

		// 标题
		list.Infos[i].Title = a_node.Text()

		// 发布时间
		list.Infos[i].Date = s.Find(".time").Text()
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

	art_list := doc.Find(".list-article .list-item")

	list := &List{
		Infos: make([]PageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		a_node := s.Find(".art-item-c2 h3 a")

		// 详细页网址
		if href, has := a_node.Attr("href"); has {
			list.Infos[i].Url = href
		}

		// 标题
		list.Infos[i].Title = a_node.Text()

		// 封面
		img_node := s.Find(".art-item-c1 a img")
		if src, has := img_node.Attr("src"); has {
			list.Infos[i].ImgSrc = src
		}

		// 发布时间
		list.Infos[i].Date = s.Find(".art-item-c2 span").Eq(1).Text()
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

	art_list := doc.Find(".list-news .art-item")

	list := &List{
		Infos: make([]PageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		a_node := s.Find(".art-item-c2 h3 a")

		// 详细页网址
		if href, has := a_node.Attr("href"); has {
			list.Infos[i].Url = href
		}

		// 标题
		list.Infos[i].Title = a_node.Text()

		// 封面
		img_node := s.Find(".art-item-c1 a img")
		if src, has := img_node.Attr("src"); has {
			list.Infos[i].ImgSrc = src
		}

		// 发布时间
		list.Infos[i].Date = strings.SplitN(
			s.Find(".info .c1").Text(),
			"：",
			2,
		)[1]
	})

	next_pages(list, doc)

	return list, nil
}

func DOTA2List(list_url string) (*List, error) {
	// 因为服务器返回gb2312编码,但goquery默认采用utf8,所以采用mahonia转换
	resp, err := http.Get(list_url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gbk := mahonia.NewDecoder("gbk")

	// 列表页
	doc, err := goquery.NewDocumentFromReader(gbk.NewReader(resp.Body))
	if err != nil {
		return nil, err
	}

	art_list := doc.Find(".art-list-txt .art-item")

	list := &List{
		Infos: make([]PageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		a_node := s.Find(".art-item-c2 a")

		// 详细页网址
		if href, has := a_node.Attr("href"); has {
			list.Infos[i].Url = href
		}

		// 标题
		list.Infos[i].Title = a_node.Text()

		// 发布时间
		list.Infos[i].Date = s.Find(".art-item-c2 .time").Text()
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

	art_list := doc.Find(".list-news .art-item")

	list := &List{
		Infos: make([]PageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		a_node := s.Find(".art-item-c2 h3 a")

		// 详细页网址
		if href, has := a_node.Attr("href"); has {
			list.Infos[i].Url = href
		}

		// 标题
		list.Infos[i].Title = a_node.Text()

		// 封面
		img_node := s.Find(".art-item-c1 a img")
		if src, has := img_node.Attr("src"); has {
			list.Infos[i].ImgSrc = src
		}

		// 发布时间
		list.Infos[i].Date = strings.SplitN(
			s.Find(".info .c3").Text(),
			"：",
			2,
		)[1]
	})

	next_pages(list, doc)

	return list, nil
}

func next_pages(list *List, doc *goquery.Document) {
	url_list := doc.Find(".pagination li a")

	list.Urls = make([]string, url_list.Length())

	url_list.Each(func(i int, s *goquery.Selection) {
		list.Urls[i] = s.AttrOr("href", "")
	})
}

func NewList(list_url string) (*List, error) {
	var game_list = map[string]func(string) (*List, error){
		"http://lol.":   LOLList,
		"http://ow.":    OWList,
		"http://csgo.":  CSGOList,
		"http://news.":  MEList,
		"http://dota2.": DOTA2List,
	}
	for prefix, handler := range game_list {
		if strings.HasPrefix(list_url, prefix) {
			return handler(list_url)
		}
	}
	return nil, errors.New("can not support " + list_url)
}
