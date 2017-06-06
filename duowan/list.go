/*go**************************************************************************
 File            : list.go
 Subsystem       : duowan
 Author          : yijunjun
 Date&Time       : 2017-06-02
 Description     : duowan-列表页
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package duowan

import (
	"errors"
	"net/url"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
	. "github.com/yijunjun/news/model"
)

func LOLList(list_url string) (*TList, error) {
	u, err := url.Parse(list_url)
	if err != nil {
		return nil, err
	}

	prefix := u.Scheme + "//" + u.Host

	base := prefix + path.Dir(u.Path) + "/"

	// 列表页
	doc, err := goquery.NewDocument(list_url)
	if err != nil {
		return nil, err
	}

	art_list := doc.Find(".ZQ-page--news .m-list li")

	list := &TList{
		Infos: make([]TPageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		a_node := s.ChildrenFiltered("a")

		// 详细页网址
		if href, has := a_node.Attr("href"); has {
			list.Infos[i].Url = prefix + href
		}

		// 标题
		if title, has := a_node.Attr("title"); has {
			list.Infos[i].Title = title
		}

		// 发布时间
		list.Infos[i].Date = s.ChildrenFiltered(".date").Text()
	})

	next_pages(list, doc, base)

	return list, nil
}

func OWList(list_url string) (*TList, error) {
	u, err := url.Parse(list_url)
	if err != nil {
		return nil, err
	}

	prefix := u.Scheme + "//" + u.Host

	base := prefix + path.Dir(u.Path) + "/"

	// 列表页
	doc, err := goquery.NewDocument(list_url)
	if err != nil {
		return nil, err
	}

	art_list := doc.Find(".ch-list-sec a")

	list := &TList{
		Infos: make([]TPageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		// 详细页网址
		if href, has := s.Attr("href"); has {
			list.Infos[i].Url = prefix + href
		}

		// 标题
		list.Infos[i].Title = s.Find(".titles").Text()

		// 封面
		img_node := s.Find(".ch-img img")
		if src, has := img_node.Attr("src"); has {
			list.Infos[i].ImgSrc = src
		}

		// 发布时间
		list.Infos[i].Date = s.Find(".times").Text()
	})

	next_pages(list, doc, base)

	return list, nil
}

func MEList(list_url string) (*TList, error) {
	u, err := url.Parse(list_url)
	if err != nil {
		return nil, err
	}

	prefix := u.Scheme + "//" + u.Host

	base := prefix + path.Dir(u.Path)

	// 列表页
	doc, err := goquery.NewDocument(list_url)
	if err != nil {
		return nil, err
	}

	art_list := doc.Find(".news-list-page li a")

	list := &TList{
		Infos: make([]TPageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		// 详细页网址
		if href, has := s.Attr("href"); has {
			list.Infos[i].Url = href
		}

		// 标题
		if title, has := s.Attr("title"); has {
			list.Infos[i].Title = title
		}

		// 发布时间
		list.Infos[i].Date = s.ChildrenFiltered(".news-time").Text()
	})

	next_pages(list, doc, base)

	return list, nil
}

func DOTA2List(list_url string) (*TList, error) {
	u, err := url.Parse(list_url)
	if err != nil {
		return nil, err
	}

	prefix := u.Scheme + "//" + u.Host

	base := prefix + path.Dir(u.Path)

	// 列表页
	doc, err := goquery.NewDocument(list_url)
	if err != nil {
		return nil, err
	}

	art_list := doc.Find("#list li")

	list := &TList{
		Infos: make([]TPageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		a_node := s.ChildrenFiltered("a")

		// 详细页网址
		if href, has := a_node.Attr("href"); has {
			if !strings.HasPrefix(href, "http") {
				href = prefix + href
			}
			list.Infos[i].Url = href
		}

		// 标题
		if title, has := a_node.Attr("title"); has {
			list.Infos[i].Title = title
		}

		// 发布时间
		list.Infos[i].Date = s.ChildrenFiltered(".date").Text()
	})

	next_pages(list, doc, base)

	return list, nil
}

func CSGOList(list_url string) (*TList, error) {
	u, err := url.Parse(list_url)
	if err != nil {
		return nil, err
	}

	prefix := u.Scheme + "//" + u.Host

	base := prefix + path.Dir(u.Path)

	// 列表页
	doc, err := goquery.NewDocument(list_url)
	if err != nil {
		return nil, err
	}

	art_list := doc.Find(".list-artical li a")

	list := &TList{
		Infos: make([]TPageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		// 详细页网址
		if href, has := s.Attr("href"); has {
			list.Infos[i].Url = prefix + href
		}

		// 标题
		list.Infos[i].Title = s.Find(".item-cont h2").Text()

		// 封面
		img_node := s.Find(".item-cover img")
		if src, has := img_node.Attr("src"); has {
			list.Infos[i].ImgSrc = src
		}

		// 发布时间
		list.Infos[i].Date = s.Find(".item-time").Text()
	})

	next_pages(list, doc, base)

	return list, nil
}

func next_pages(list *TList, doc *goquery.Document, base string) {
	url_list := doc.Find("#pageNum span a")

	list.Urls = make([]string, url_list.Length())

	url_list.Each(func(i int, s *goquery.Selection) {
		tmp := s.AttrOr("href", "#")
		if tmp != "#" {
			list.Urls[i] = base + tmp
		}
	})
}

func NewList(list_url string) (*TList, error) {
	var game_list = map[string]func(string) (*TList, error){
		"http://lol.":   LOLList,
		"http://ow.":    OWList,
		"http://csgo.":  CSGOList,
		"http://wzry.":  MEList,
		"http://dota2.": DOTA2List,
	}
	for prefix, handler := range game_list {
		if strings.HasPrefix(list_url, prefix) {
			return handler(list_url)
		}
	}
	return nil, errors.New("can not support " + list_url)
}
