/*go**************************************************************************
 File            : list.go
 Subsystem       : pcgames
 Author          : yijunjun
 Date&Time       : 2017-06-05
 Description     : pcgames-列表页
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package pcgames

import (
	"errors"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	. "github.com/yijunjun/news/model"
)

func LOLList(list_url string) (*List, error) {
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

	art_list := doc.Find("#topnewsContent .media-body")

	list := &List{
		Infos: make([]PageInfo, art_list.Length()),
	}

	art_list.Each(func(i int, s *goquery.Selection) {
		a_node := s.Find(".media-body-title a")

		// 详细页网址
		if href, has := a_node.Attr("href"); has {
			list.Infos[i].Url = href
		}

		// 标题
		list.Infos[i].Title = strings.TrimSpace(a_node.Text())

		// 发布时间
		list.Infos[i].Date = strings.TrimSpace(s.Find(".date").Text())
	})

	next_pages(list, doc)

	return list, nil
}

func DOTA2List(list_url string) (*List, error) {
	return LOLList(list_url)
}

func CSGOList(list_url string) (*List, error) {
	return DOTA2List(list_url)
}

func next_pages(list *List, doc *goquery.Document) {
	url_list := doc.Find("#page a[href]")

	list.Urls = make([]string, url_list.Length())

	url_list.Each(func(i int, s *goquery.Selection) {
		list.Urls[i] = s.AttrOr("href", "")
	})
}

func NewList(list_url string) (*List, error) {
	var game_list = map[string]func(string) (*List, error){
		"http://wangyou.":                    LOLList,
		"http://fight.pcgames.com.cn/dota2/": DOTA2List,
		"http://fight.pcgames.com.cn/cs":     CSGOList,
	}
	for prefix, handler := range game_list {
		if strings.HasPrefix(list_url, prefix) {
			return handler(list_url)
		}
	}
	return nil, errors.New("can not support " + list_url)
}
