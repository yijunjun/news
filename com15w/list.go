/*go**************************************************************************
 File            : list.go
 Subsystem       : com15w
 Author          : yijunjun
 Date&Time       : 2017-06-02
 Description     : 电竞头条-列表页
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package com15w

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/yijunjun/news/common"
	. "github.com/yijunjun/news/model"
)

func LOLList(list_url string) (*TList, error) {
	// 列表页
	doc, err := common.Download(list_url)
	if err != nil {
		return nil, err
	}

	art_list := doc.Find(".indexcon .artlist")

	list := &TList{
		Infos: make([]TPageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		a_node := s.Find(".artli-rbox a")

		// 详细页网址
		if href, has := a_node.Attr("href"); has {
			list.Infos[i].Url = href
		}
		// 标题
		if title, has := a_node.Attr("title"); has {
			list.Infos[i].Title = title
		}

		// 封面
		img_node := s.Find(".pic img")
		if src, has := img_node.Attr("src"); has {
			list.Infos[i].ImgSrc = src
		}

		// 发布时间
		list.Infos[i].Date = s.Find(".art-time").Text()
	})

	next_pages(list, doc)

	return list, nil
}

func OWList(list_url string) (*TList, error) {
	return LOLList(list_url)
}

func MEList(list_url string) (*TList, error) {
	// 列表页
	doc, err := common.Download(list_url)
	if err != nil {
		return nil, err
	}

	art_list := doc.Find(".hot-newslist .hot-mp")

	list := &TList{
		Infos: make([]TPageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		a_node := s.Find(".r a")

		// 详细页网址
		if href, has := a_node.Attr("href"); has {
			list.Infos[i].Url = href
		}
		// 标题
		if title, has := a_node.Attr("title"); has {
			list.Infos[i].Title = title
		}

		// 封面
		img_node := s.Find("a img")
		if src, has := img_node.Attr("src"); has {
			list.Infos[i].ImgSrc = src
		}

		// 发布时间
		list.Infos[i].Date = s.Find(".r .date").Text()
	})

	next_pages(list, doc)

	return list, nil
}

func DOTA2List(list_url string) (*TList, error) {
	// 列表页
	doc, err := common.Download(list_url)
	if err != nil {
		return nil, err
	}

	art_list := doc.Find(".col1 li")

	list := &TList{
		Infos: make([]TPageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		a_node := s.ChildrenFiltered("a")

		// 详细页网址
		if href, has := a_node.Attr("href"); has {
			list.Infos[i].Url = href
		}

		// 标题
		list.Infos[i].Title = a_node.ChildrenFiltered(".fli").Text()

		// 发布时间
		list.Infos[i].Date = a_node.ChildrenFiltered(".date").Text()
	})

	next_pages(list, doc)

	return list, nil
}

func CSGOList(list_url string) (*TList, error) {
	// 列表页
	doc, err := common.Download(list_url)
	if err != nil {
		return nil, err
	}

	art_list := doc.Find(".mainTabCon .artlist")

	list := &TList{
		Infos: make([]TPageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		a_node := s.Find(".artli-rbox a")

		// 详细页网址
		if href, has := a_node.Attr("href"); has {
			list.Infos[i].Url = href
		}
		// 标题
		if title, has := a_node.Attr("title"); has {
			list.Infos[i].Title = title
		}

		// 封面
		img_node := s.Find("a img")
		if src, has := img_node.Attr("src"); has {
			list.Infos[i].ImgSrc = src
		}

		// 发布时间
		list.Infos[i].Date = s.Find(".art-time").Text()
	})

	next_pages(list, doc)

	return list, nil
}

func next_pages(list *TList, doc *goquery.Document) {
	url_list := doc.Find(".pages a")

	list.Urls = make([]string, url_list.Length())

	url_list.Each(func(i int, s *goquery.Selection) {
		list.Urls[i] = s.AttrOr("href", "")
	})
}
