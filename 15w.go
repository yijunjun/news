// 电竞头条网站处理
package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// 构造请求获取阅读次数
func HitsCount(id string) (string, error) {
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
func CmtSum(id string) (string, error) {
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

// 详情页,每个游戏各不相同,需要特别区分对待.
func main() {
	doc, err := goquery.NewDocument("http://me.15w.com/ss/5605408165.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	info_ele := doc.Find(".headline-box .info span")
	if info_ele == nil {
		fmt.Println("info_ele is nil")
		return
	}

	author := info_ele.Eq(1).Text()

	source := info_ele.Eq(5).Text()

	date := info_ele.Eq(6).Text()

	fmt.Println("attr", author, source, date)

	// 获取文章id
	aid, has := info_ele.Eq(8).Attr("aid")
	if has {
		hc, err := HitsCount(aid)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("hc", hc)

		cs, err := CmtSum(aid)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("cs", cs)
	}
}

/*
func main() {
	// 列表页
	doc, err := goquery.NewDocument("http://lol.15w.com/zx/")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Find the review items
	doc.Find(".indexcon .artlist").Each(func(i int, s *goquery.Selection) {
		a_node := s.ChildrenFiltered("a")
		// 详细页
		if href, has := a_node.Attr("href"); has {
			fmt.Println("href", href)
		}
		// 标题
		if title, has := a_node.Attr("title"); has {
			fmt.Println("title", title)
		}

		// 封面
		a_img_node := a_node.ChildrenFiltered("img")
		if src, has := a_img_node.Attr("src"); has {
			fmt.Println("img src", src)
		}

		// 发布时间
		art_time := s.Find(".art-time").Text()
		fmt.Println("art_time", art_time)
	})
}
*/
