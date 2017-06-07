/*go**************************************************************************
 File            : list.go
 Subsystem       : com178
 Author          : yijunjun
 Date&Time       : 2017-06-02
 Description     : 178-详情页
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package com178

import (
	"bytes"
	"errors"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/yijunjun/news/common"
	. "github.com/yijunjun/news/model"
)

func LOLPage(page_url string) (*TPage, error) {
	doc, err := common.Download(page_url)
	if err != nil {
		return nil, err
	}

	info := html.UnescapeString(doc.Find(".article_lf_con_author").Text())

	author_flag := "作者："
	source_flag := "来源："
	date_flag := "发布时间："

	anchor_flag_pos := strings.Index(info, author_flag)
	if anchor_flag_pos == -1 {
		return nil, common.NewSelfError("can not find author:" + info)
	}

	source_flag_pos := strings.Index(info, source_flag)
	if source_flag_pos == -1 {
		return nil, common.NewSelfError("can not find source:" + info)
	}

	date_flag_pos := strings.Index(info, date_flag)
	if date_flag_pos == -1 {
		return nil, common.NewSelfError("can not find date:" + info)
	}

	return &TPage{
		Author: info[anchor_flag_pos+len(author_flag) : source_flag_pos],
		Source: info[source_flag_pos+len(source_flag) : date_flag_pos],
		Date:   info[date_flag_pos+len(date_flag):],
	}, nil
}

func OWPage(page_url string) (*TPage, error) {
	doc, err := common.Download(page_url)
	if err != nil {
		return nil, err
	}

	info_ele := doc.Find(".arc-text h2 span")
	if info_ele == nil {
		return nil, common.NewSelfError("span find failure:" + page_url)
	}

	author2list := strings.SplitN(
		info_ele.Eq(0).Text(),
		"：",
		2,
	)
	if len(author2list) != 2 {
		return nil, common.NewSelfError("author splitN failure:" + info_ele.Eq(0).Text())
	}

	source2list := strings.SplitN(
		info_ele.Eq(1).Text(),
		"：",
		2,
	)
	if len(source2list) != 2 {
		return nil, common.NewSelfError("source splitN failure:" + info_ele.Eq(1).Text())
	}

	date2list := strings.SplitN(
		info_ele.Eq(2).Text(),
		"：",
		2,
	)
	if len(date2list) != 2 {
		return nil, common.NewSelfError("date splitN failure:" + info_ele.Eq(2).Text())
	}

	return &TPage{
		Author: author2list[1],
		Source: source2list[1],
		Date:   date2list[1],
	}, nil
}

func MEPage(page_url string) (*TPage, error) {
	doc, err := common.Download(page_url)
	if err != nil {
		return nil, err
	}

	info_ele := doc.Find(".author-item")
	if info_ele == nil {
		return nil, common.NewSelfError("author-item find failure:" + page_url)
	}

	return &TPage{
		Author: info_ele.Find("span[author]").AttrOr("author", ""),
		Date:   info_ele.ChildrenFiltered(".time").Text(),
	}, nil
}

func DOTA2Page(page_url string) (*TPage, error) {
	doc, err := goquery.NewDocument(page_url)
	if err != nil {
		return nil, err
	}

	info := doc.Find("div[class=info]")

	page := &TPage{
		Author: info.ChildrenFiltered(".author").Text(),
		Date:   info.ChildrenFiltered(".time").Text(),
	}

	return page, nil
}

func CSGOPage(page_url string) (*TPage, error) {
	doc, err := common.Download(page_url)
	if err != nil {
		return nil, err
	}

	info_ele := doc.Find(".author span")
	if info_ele == nil {
		return nil, common.NewSelfError("author find failure:" + page_url)
	}

	source2list := strings.SplitN(
		info_ele.Eq(0).Text(),
		"：",
		2,
	)
	if len(source2list) != 2 {
		return nil, common.NewSelfError("source splitN failure:" + info_ele.Eq(0).Text())
	}

	date2list := strings.SplitN(
		info_ele.Eq(2).Text(),
		"：",
		2,
	)
	if len(date2list) != 2 {
		return nil, common.NewSelfError("date splitN failure:" + info_ele.Eq(2).Text())
	}

	return &TPage{
		Source: source2list[1],
		Date:   date2list[1],
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
