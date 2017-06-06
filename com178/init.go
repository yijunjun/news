/*go**************************************************************************
 File            : init.go
 Subsystem       : com178
 Author          : yijunjun
 Date&Time       : 2017-06-06
 Description     : 178-初始化
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package com178

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/yijunjun/news/common"
	. "github.com/yijunjun/news/model"
)

var g_digit_reg = regexp.MustCompile(`\d+`)

type TSiteNode struct {
	ListPrefix string
	ListFun    func(string) (*TList, error)

	PagePrefix *regexp.Regexp
	PageFun    func(string) (*TPage, error)
}

func init() {
	common.RegSiteHandle("178.com", func(raw_url string, u *url.URL) error {
		for game, sn := range map[string]TSiteNode{
			"lol": TSiteNode{
				ListPrefix: "/list/",
				ListFun:    LOLList,
				PagePrefix: g_digit_reg,
				PageFun:    LOLPage,
			},
			"ow": TSiteNode{
				ListPrefix: "/list/",
				ListFun:    OWList,
				PagePrefix: g_digit_reg,
				PageFun:    OWPage,
			},
			"shouyou": TSiteNode{
				ListPrefix: "/list/",
				ListFun:    MEList,
				PagePrefix: g_digit_reg,
				PageFun:    MEPage,
			},
			"dota2": TSiteNode{
				ListPrefix: "/list/",
				ListFun:    DOTA2List,
				PagePrefix: g_digit_reg,
				PageFun:    DOTA2Page,
			},
			"csgo": TSiteNode{
				ListPrefix: "/list/",
				ListFun:    CSGOList,
				PagePrefix: g_digit_reg,
				PageFun:    CSGOPage,
			},
		} {
			if strings.HasPrefix(u.Host, game) {
				if strings.HasPrefix(u.Path, sn.ListPrefix) {
					return SaveList(sn.ListFun(raw_url))
				}

				plist := strings.Split(u.Path, "/")
				if len(plist) > 0 && sn.PagePrefix.MatchString(plist[0]) {
					return SavePage(sn.PageFun(raw_url))
				}
			}
		}
		return common.NewSelfError("can not support:" + raw_url)
	})
}
