/*go**************************************************************************
 File            : init.go
 Subsystem       : com17173
 Author          : yijunjun
 Date&Time       : 2017-06-06
 Description     : com17173-初始化
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package com17173

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/yijunjun/news/common"
	. "github.com/yijunjun/news/model"
)

var g_dota2_list_reg = regexp.MustCompile(`/news/index_\d+.shtml`)

var g_dota2_page_reg = regexp.MustCompile(`/news/\d+/\d+.shtml`)

func init() {
	common.RegSiteHandle("17173.com", func(raw_url string, u *url.URL) error {
		switch strings.SplitN(u.Host, ".", 2)[0] {
		case "lol":
			if strings.HasPrefix(u.Path, "/list/") {
				return SaveList(LOLList(raw_url))
			}

			if strings.HasPrefix(u.Path, "/news/") {
				return SavePage(LOLPage(raw_url))
			}

		case "ow":
			return SaveList(OWList(raw_url))
		case "bbs":
			return SavePage(OWPage(raw_url))

		case "news":
			if strings.Contains(u.Path, "/list/") {
				return SaveList(MEList(raw_url))
			}

			if strings.Contains(u.Path, "/content/") {
				return SavePage(MEPage(raw_url))
			}

		case "dota2":
			if g_dota2_list_reg.MatchString(u.Path) {
				return SaveList(DOTA2List(raw_url))
			}

			if g_dota2_page_reg.MatchString(u.Path) {
				return SavePage(DOTA2Page(raw_url))
			}

		case "csgo":
			if strings.HasPrefix(u.Path, "/z/") {
				return SaveList(CSGOList(raw_url))
			}

			if strings.HasPrefix(u.Path, "/content/") {
				return SavePage(CSGOPage(raw_url))
			}
		}
		return common.NewSelfError("can not support:" + raw_url)
	})
}
