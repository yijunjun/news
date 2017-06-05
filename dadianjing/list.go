/*go**************************************************************************
 File            : list.go
 Subsystem       : dadianjing
 Author          : yijunjun
 Date&Time       : 2017-06-05
 Description     : dadianjing-列表页
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package dadianjing

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/blemobi/go-commons"
	. "github.com/yijunjun/news/model"
)

type Art struct {
	Id    string
	Time  string
	Title string
	Thumb string
}

type ArtJson struct {
	Data struct {
		List []Art
		Page struct {
			PageCount int
			Page      int
		}
	}
}

func game_list(list_url string) (*List, error) {
	u, err := url.Parse(list_url)
	if err != nil {
		return nil, err
	}

	cid := u.Query().Get("cid")

	bs := common.RemoteCall(list_url)
	if bs == nil {
		return nil, errors.New("bs is nil")
	}
	aj := &ArtJson{}
	err = json.Unmarshal(bs, aj)
	if err != nil {
		return nil, err
	}

	list := &List{
		Infos: make([]PageInfo, len(aj.Data.List)),
	}

	for i, art := range aj.Data.List {
		list.Infos[i].Url = fmt.Sprintf("http://www.dadianjing.cn/info/%v.html", art.Id)
		list.Infos[i].Title = art.Title
		list.Infos[i].ImgSrc = art.Thumb
		list.Infos[i].Date = art.Time
	}

	nowPage := aj.Data.Page.Page
	allPage := aj.Data.Page.PageCount

	url_fmt := "http://www.dadianjing.cn/index.php?m=Index&a=xhrList&cid=%v&page=%v"

	list.Urls = []string{}

	if nowPage > 1 {
		// 上一页
		list.Urls = append(list.Urls, fmt.Sprintf(url_fmt, cid, nowPage-1))
	}

	if nowPage < allPage {
		// 下一页
		list.Urls = append(list.Urls, fmt.Sprintf(url_fmt, cid, nowPage+1))
	}

	return list, nil
}

func LOLList(list_url string) (*List, error) {
	return game_list(list_url)
}

func OWList(list_url string) (*List, error) {
	return game_list(list_url)
}

func CSGOList(list_url string) (*List, error) {
	return game_list(list_url)
}

func NewList(list_url string) (*List, error) {
	u, err := url.Parse(list_url)
	if err != nil {
		return nil, err
	}

	str2list := strings.SplitN(path.Base(u.Path), ".", 2)
	if len(str2list) != 2 {
		return nil, errors.New(path.Base(u.Path) + " failure")
	}

	return game_list(fmt.Sprintf(
		"http://www.dadianjing.cn/index.php?m=Index&a=xhrList&cid=%v&page=1",
		str2list[1],
	))
}
