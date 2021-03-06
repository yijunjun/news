/*go**************************************************************************
 File            : list.go
 Subsystem       : com17173
 Author          : yijunjun
 Date&Time       : 2017-06-02
 Description     : 17173-详情页
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package com17173

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/yijunjun/news/common"
	. "github.com/yijunjun/news/model"
)

func LOLPage(page_url string) (*TPage, error) {
	doc, err := common.Download(page_url)
	if err != nil {
		return nil, err
	}

	info_ele := doc.Find(".gb-final-mod-info span")
	if info_ele == nil {
		return nil, common.NewSelfError("gb-final-mod-info find failure: " + page_url)
	}

	date2list := strings.SplitN(
		info_ele.Eq(0).Text(),
		"：",
		2,
	)
	if len(date2list) != 2 {
		return nil, common.NewSelfError("date splitN failure: " + info_ele.Eq(0).Text())
	}

	return &TPage{
		Author: info_ele.Eq(1).ChildrenFiltered("b").Text(),
		Source: info_ele.Eq(2).ChildrenFiltered("b").Text(),
		Date:   date2list[1],
	}, nil
}

func OWPage(page_url string) (*TPage, error) {
	return LOLPage(page_url)
}

func MEPage(page_url string) (*TPage, error) {
	return LOLPage(page_url)
}

func DOTA2Page(page_url string) (*TPage, error) {
	doc, err := common.Download(page_url, "gbk")
	if err != nil {
		return nil, err
	}

	info_ele := doc.Find(".info span")
	if info_ele == nil {
		return nil, errors.New("span find failure")
	}

	date2list := strings.SplitN(
		info_ele.Eq(0).Text(),
		"：",
		2,
	)
	if len(date2list) != 2 {
		return nil, common.NewSelfError("date splitN failure: " + info_ele.Eq(0).Text())
	}

	author2list := strings.SplitN(
		info_ele.Eq(1).Text(),
		"：",
		2,
	)
	if len(author2list) != 2 {
		return nil, common.NewSelfError("author splitN failure: " + info_ele.Eq(1).Text())
	}

	return &TPage{
		Author: author2list[1],
		Date:   date2list[1],
	}, nil
}

func CSGOPage(page_url string) (*TPage, error) {
	doc, err := common.Download(page_url)
	if err != nil {
		return nil, err
	}

	info_ele := doc.Find(".info span")
	if info_ele == nil {
		return nil, errors.New("span find failure")
	}

	date2list := strings.SplitN(
		info_ele.Eq(0).Text(),
		"：",
		2,
	)
	if len(date2list) != 2 {
		return nil, common.NewSelfError("date splitN failure: " + info_ele.Eq(0).Text())
	}

	author2list := strings.SplitN(
		info_ele.Eq(1).Text(),
		"：",
		2,
	)
	if len(author2list) != 2 {
		return nil, common.NewSelfError("author splitN failure: " + info_ele.Eq(1).Text())
	}

	return &TPage{
		Author: author2list[1],
		Date:   date2list[1],
	}, nil
}

func BBSPage(page_url string) (*TPage, error) {
	doc, err := common.Download(page_url)
	if err != nil {
		return nil, err
	}

	span_ele := doc.Find("#postlist .hm span[class=xi1]")

	date := doc.Find(".authi span[title]").Eq(0).AttrOr("title", "")
	if date == "" {
		date = doc.Find(".authi em[id]").Eq(0).Text()[len("发表于"):]
	}

	return &TPage{
		Author:       doc.Find(".authi a[class=xw1]").Eq(0).Text(),
		Date:         date,
		HitsCount:    span_ele.Eq(0).Text(),
		CommentCount: span_ele.Eq(1).Text(),
	}, nil
}

// 构造请求获取阅读次数
func hits_count(id string) (string, error) {
	unix_time := time.Now().Unix()
	target := fmt.Sprintf(
		"http://www.15w.com/index.php?callback=jsonp%v&tn=ajax&ac=cmshits&type=news&id=%v&h=%v",
		unix_time,
		id,
		unix_time,
	)
	resp, err := http.Get(target)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	left := bytes.IndexByte(bs, byte('('))
	if left == -1 {
		return "", errors.New("can not found (")
	}

	right := bytes.IndexByte(bs, byte(')'))
	if right == -1 {
		return "", errors.New("can not found )")
	}

	if left >= right {
		return "", errors.New(fmt.Sprintf("( at pos %v >= ) at pos %v", left, right))
	}

	return string(bs[left+1 : right]), nil
}

var g_cmt_reg = regexp.MustCompile(`"comments":(\d+)`)

// 利用搜狐的畅言做的评论系统
func comment_count(id string) (string, error) {
	target := fmt.Sprintf(
		"https://changyan.sohu.com/api/2/topic/count?client_id=cyqQ61pvv&topic_source_id=1/%v",
		id,
	)
	resp, err := http.Get(target)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	m := g_cmt_reg.FindSubmatch(bs)
	if m == nil {
		return "", errors.New("regexp failure")
	}

	return string(m[1]), nil
}
