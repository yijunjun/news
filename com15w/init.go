/*go**************************************************************************
 File            : init.go
 Subsystem       : com15w
 Author          : yijunjun
 Date&Time       : 2017-06-02
 Description     : 电竞头条-初始化
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package com15w

import (
	"net/url"
	"strings"

	"github.com/yijunjun/news/common"
	. "github.com/yijunjun/news/model"
)

type TSiteNode struct {
	ListPrefix string
	ListFun    func(string) (*TList, error)

	PagePrefix string
	PageFun    func(string) (*TPage, error)
}

func init() {
	common.RegSiteHandle("15w.com", func(raw_url string, u *url.URL) error {
		for game, sn := range map[string]TSiteNode{
			"lol": TSiteNode{
				ListPrefix: "/zx/",
				ListFun:    LOLList,
				PagePrefix: "/ss/",
				PageFun:    LOLPage,
			},
			"ow": TSiteNode{
				ListPrefix: "/xw/",
				ListFun:    OWList,
				PagePrefix: "/news/",
				PageFun:    OWPage,
			},
			"me": TSiteNode{
				ListPrefix: "/wzry/",
				ListFun:    MEList,
				PagePrefix: "/dj/",
				PageFun:    MEPage,
			},
			"dota2": TSiteNode{
				ListPrefix: "/news/",
				ListFun:    DOTA2List,
				PagePrefix: "/guofuxinwen/",
				PageFun:    DOTA2Page,
			},
			"csgo": TSiteNode{
				ListPrefix: "/news/",
				ListFun:    CSGOList,
				PagePrefix: "/yxzx/",
				PageFun:    CSGOPage,
			},
		} {
			if strings.HasPrefix(u.Host, game) {
				if strings.HasPrefix(u.Path, sn.ListPrefix) {
					return SaveList(sn.ListFun(raw_url))
				}

				if strings.HasPrefix(u.Path, sn.PagePrefix) {
					return SavePage(sn.PageFun(raw_url))
				}
			}
		}
		return common.NewSelfError("can not support:" + raw_url)
	})
}
