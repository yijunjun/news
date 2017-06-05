/*go**************************************************************************
 File            : list.go
 Subsystem       : ooqiu
 Author          : yijunjun
 Date&Time       : 2017-06-05
 Description     : ooqiu-列表页
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package ooqiu

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	. "github.com/yijunjun/news/model"

	"path"

	"github.com/axgle/mahonia"
)

func url_dir(list_url string) (string, error) {
	u, err := url.Parse(list_url)
	if err != nil {
		return "", err
	}

	base := u.Scheme + "//" + u.Host

	if strings.HasSuffix(u.Path, "/") {
		base = base + u.Path
	} else {
		base = base + path.Dir(u.Path) + "/"
	}
	return base, nil
}

func LOLList(list_url string) (*List, error) {
	base, err := url_dir(list_url)
	if err != nil {
		return nil, err
	}

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

	art_list := doc.Find("#list .arc-list li")

	list := &List{
		Infos: make([]PageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		a_node := s.Find(".title a")

		// 详细页网址
		if href, has := a_node.Attr("href"); has {
			list.Infos[i].Url = href
		}

		// 标题
		if title, has := a_node.Attr("title"); has {
			list.Infos[i].Title = title
		}

		// 封面
		img_node := s.Find(".arc-img img")
		if src, has := img_node.Attr("data-echo"); has {
			list.Infos[i].ImgSrc = src
		}

		// 发布时间
		list.Infos[i].Date = s.Find(".time").Text()
	})

	next_pages(list, doc, base, "#page a[href]")

	return list, nil
}

func OWList(list_url string) (*List, error) {
	base, err := url_dir(list_url)
	if err != nil {
		return nil, err
	}

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

	art_list := doc.Find(".list-list li")

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
		img_node := a_node.ChildrenFiltered("img")
		if src, has := img_node.Attr("data-echo"); has {
			list.Infos[i].ImgSrc = src
		}

		// 发布时间
		list.Infos[i].Date = s.Find("em").Eq(2).Text()
	})

	next_pages(list, doc, base, ".mod_page a[href]")

	return list, nil
}

func MEList(list_url string) (*List, error) {
	return LOLList(list_url)
}

func DOTA2List(list_url string) (*List, error) {
	base, err := url_dir(list_url)
	if err != nil {
		return nil, err
	}

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

	art_list := doc.Find("li[class=yc01]")

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
		img_node := a_node.ChildrenFiltered("img")
		if src, has := img_node.Attr("data-echo"); has {
			list.Infos[i].ImgSrc = src
		}

		// 发布时间
		date2list := strings.SplitN(
			s.Find(".others span").Eq(2).Text(),
			"：",
			2,
		)
		if len(date2list) == 2 {
			list.Infos[i].Date = date2list[1]
		}
	})

	next_pages(list, doc, base, ".mod_page a[href]")

	return list, nil
}

func CSGOList(list_url string) (*List, error) {
	return DOTA2List(list_url)
}

func next_pages(list *List, doc *goquery.Document, base, selector string) {
	url_list := doc.Find(selector)

	list.Urls = make([]string, url_list.Length())

	url_list.Each(func(i int, s *goquery.Selection) {
		list.Urls[i] = s.AttrOr("href", "")
		if list.Urls[i] != "" {
			list.Urls[i] = base + list.Urls[i]
		}
	})
}

func NewList(list_url string) (*List, error) {
	var game_list = map[string]func(string) (*List, error){
		"http://lol.":   LOLList,
		"http://ow.":    OWList,
		"http://csgo.":  CSGOList,
		"http://dota2.": DOTA2List,
		"http://sy.":    MEList,
	}
	for prefix, handler := range game_list {
		if strings.HasPrefix(list_url, prefix) {
			return handler(list_url)
		}
	}
	return nil, errors.New("can not support " + list_url)
}
