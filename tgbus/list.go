/*go**************************************************************************
 File            : list.go
 Subsystem       : tgbus
 Author          : yijunjun
 Date&Time       : 2017-06-05
 Description     : tgbus-列表页
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package tgbus

import (
	"errors"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	. "github.com/yijunjun/news/model"

	"path"
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

	// 列表页
	doc, err := goquery.NewDocument(list_url)
	if err != nil {
		return nil, err
	}

	art_list := doc.Find(".list li")

	list := &List{
		Infos: make([]PageInfo, 0),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		a_node := s.ChildrenFiltered("a[title]")

		// 标题
		title, has := a_node.Attr("title")
		if !has || title == "" {
			return
		}

		info := PageInfo{Title: title}

		// 详细页网址
		if href, has := a_node.Attr("href"); has {
			info.Url = href
		}

		// 封面
		img_node := s.Find("a img")
		if src, has := img_node.Attr("src"); has {
			info.ImgSrc = src
		}

		// 发布时间
		date2list := strings.SplitN(s.Find(".fz12").Text(), "发布时间：", 2)
		if len(date2list) != 2 {
			return
		}
		info.Date = date2list[1]

		list.Infos = append(list.Infos, info)
	})

	next_pages(list, doc, base)

	return list, nil
}

func OWList(list_url string) (*List, error) {
	base, err := url_dir(list_url)
	if err != nil {
		return nil, err
	}

	// 列表页
	doc, err := goquery.NewDocument(list_url)
	if err != nil {
		return nil, err
	}

	art_list := doc.Find(".neirong dl")

	list := &List{
		Infos: make([]PageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		a_node := s.Find("dd h6 a")

		// 详细页网址
		if href, has := a_node.Attr("href"); has {
			list.Infos[i].Url = href
		}

		// 标题
		if title, has := a_node.Attr("title"); has {
			list.Infos[i].Title = title
		}

		// 封面
		img_node := s.Find("dt a img")
		if src, has := img_node.Attr("src"); has {
			list.Infos[i].ImgSrc = src
		}

		// 发布时间
		list.Infos[i].Date = strings.Trim(s.Find("em font font").Text(), "[]")
	})

	next_pages(list, doc, base)

	return list, nil
}

func DOTA2List(list_url string) (*List, error) {
	// 列表页
	doc, err := goquery.NewDocument(list_url)
	if err != nil {
		return nil, err
	}

	art_list := doc.Find(".newslist li")

	list := &List{
		Infos: make([]PageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		a_node := s.Find(".info .title a")

		// 详细页网址
		if href, has := a_node.Attr("href"); has {
			list.Infos[i].Url = href
		}

		// 标题
		list.Infos[i].Title = a_node.Text()

		// 封面
		img_node := s.Find(".pic a img")
		if src, has := img_node.Attr("src"); has {
			list.Infos[i].ImgSrc = src
		}

		// 发布时间
		date2list := strings.SplitN(s.Find(".time").Text(), "：", 2)
		if len(date2list) == 2 {
			list.Infos[i].Date = date2list[1]
		}
	})

	url_list := doc.Find(".paging a")

	list.Urls = make([]string, url_list.Length())

	url_list.Each(func(i int, s *goquery.Selection) {
		list.Urls[i] = s.AttrOr("href", "")
	})

	return list, nil
}

func CSGOList(list_url string) (*List, error) {
	base, err := url_dir(list_url)
	if err != nil {
		return nil, err
	}

	// 列表页
	doc, err := goquery.NewDocument(list_url)
	if err != nil {
		return nil, err
	}

	art_list := doc.Find(".newslist li")

	list := &List{
		Infos: make([]PageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		a_node := s.Find(".info .title a")

		// 详细页网址
		if href, has := a_node.Attr("href"); has {
			list.Infos[i].Url = href
		}

		// 标题
		list.Infos[i].Title = a_node.Text()

		// 封面
		img_node := s.Find(".pic a img")
		if src, has := img_node.Attr("src"); has {
			list.Infos[i].ImgSrc = src
		}

		// 发布时间
		list.Infos[i].Date = s.Find(".time").Text()
	})

	url_list := doc.Find(".fanye a")

	list.Urls = make([]string, url_list.Length())

	url_list.Each(func(i int, s *goquery.Selection) {
		list.Urls[i] = s.AttrOr("href", "")
		if list.Urls[i] != "" {
			list.Urls[i] = base + list.Urls[i]
		}
	})

	return list, nil
}

func next_pages(list *List, doc *goquery.Document, base string) {
	url_list := doc.Find("#pager a")

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
	}
	for prefix, handler := range game_list {
		if strings.HasPrefix(list_url, prefix) {
			return handler(list_url)
		}
	}
	return nil, errors.New("can not support " + list_url)
}
