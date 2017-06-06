/*go**************************************************************************
 File            : list.go
 Subsystem       : com15w
 Author          : yijunjun
 Date&Time       : 2017-06-02
 Description     : 电竞头条-详情页
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package com15w

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	gcom "github.com/blemobi/go-commons"
	"github.com/yijunjun/news/common"
	. "github.com/yijunjun/news/model"
)

func LOLPage(page_url string) (*TPage, error) {
	doc, err := goquery.NewDocument(page_url)
	if err != nil {
		return nil, err
	}

	info_ele := doc.Find(".headline-box .info span")
	if info_ele == nil {
		return nil, common.NewSelfError("span find failure:" + page_url)
	}

	page := &TPage{
		Author: info_ele.Eq(1).Text(),
		Source: info_ele.Eq(5).Text(),
		Date:   info_ele.Eq(6).Text(),
	}

	// 获取文章id
	aid, has := doc.Find("#hitscount").Attr("aid")
	if !has {
		return nil, common.NewSelfError("aid find failure:" + page_url)
	}

	page.Id = aid

	page.HitsCount, err = hits_count(page.Id)
	if err != nil {
		return nil, err
	}

	page.CommentCount, err = comment_count(page.Id)
	if err != nil {
		return nil, err
	}
	return page, nil
}

func OWPage(page_url string) (*TPage, error) {
	return LOLPage(page_url)
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
		return "", common.NewSelfError("can not found (:" + target)
	}

	right := bytes.IndexByte(bs, byte(')'))
	if right == -1 {
		return "", common.NewSelfError("can not found ):" + target)
	}

	if left >= right {
		return "", common.NewSelfError(fmt.Sprintf("( at pos %v >= ) at pos %v %v", left, right, target))
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

	bs := gcom.RemoteCall(target)
	if bs == nil {
		return "", common.NewSelfError("RemoteCall failure:" + target)
	}

	m := g_cmt_reg.FindSubmatch(bs)
	if m == nil {
		return "", common.NewSelfError("regexp failure:" + target)
	}

	return string(m[1]), nil
}

func CSGOPage(page_url string) (*TPage, error) {
	doc, err := goquery.NewDocument(page_url)
	if err != nil {
		return nil, err
	}

	info_ele := doc.Find(".mainHead p span")
	if info_ele == nil {
		return nil, common.NewSelfError("span find failure:" + page_url)
	}

	page := &TPage{
		Source: info_ele.Eq(3).Text(),
		Date:   info_ele.Eq(4).Text(),
	}

	// 获取文章id
	aid, has := doc.Find("#hitscount").Attr("aid")
	if !has {
		return nil, common.NewSelfError("aid find failure:" + page_url)
	}

	page.Id = aid

	page.HitsCount, err = hits_count(page.Id)
	if err != nil {
		return nil, err
	}

	page.CommentCount, err = comment_count(page.Id)
	if err != nil {
		return nil, err
	}
	return page, nil
}

func MEPage(page_url string) (*TPage, error) {
	doc, err := goquery.NewDocument(page_url)
	if err != nil {
		return nil, err
	}

	info_ele := doc.Find(".pv-title .ti-info span")
	if info_ele == nil {
		return nil, common.NewSelfError("span find failure:" + page_url)
	}

	source2list := strings.SplitN(info_ele.Eq(2).Text(), "：", 2)
	if len(source2list) != 2 {
		return nil, common.NewSelfError(info_ele.Eq(2).Text() + " SplitN failure")
	}

	date2list := strings.SplitN(info_ele.Eq(0).Text(), "：", 2)
	if len(date2list) != 2 {
		return nil, common.NewSelfError(info_ele.Eq(0).Text() + " SplitN failure")
	}

	page := &TPage{
		Source: source2list[1],
		Date:   date2list[1],
	}

	// 获取文章id
	sid, has := doc.Find("#changyan_count_unit").Attr("sid")
	if !has {
		return nil, common.NewSelfError("sid find failure:" + page_url)
	}

	id2list := strings.SplitN(sid, "/", 2)
	if len(id2list) != 2 {
		return nil, common.NewSelfError(sid + " SplitN failure")
	}

	page.Id = id2list[1]

	page.CommentCount, err = comment_count(page.Id)
	if err != nil {
		return nil, err
	}
	return page, nil
}

func DOTA2Page(page_url string) (*TPage, error) {
	doc, err := goquery.NewDocument(page_url)
	if err != nil {
		return nil, err
	}

	info := doc.Find(".c_bor2 .t").Text()

	date2list := strings.SplitN(info, "发布时间：", 2)
	if len(date2list) != 2 {
		return nil, common.NewSelfError(info + " SplitN failure")
	}

	return &TPage{
		Date: date2list[1],
	}, nil
}
