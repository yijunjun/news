/*go**************************************************************************
 File            : list.go
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
	"errors"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/yijunjun/news/common"
	. "github.com/yijunjun/news/model"
)

func init() {
	common.RegSiteHandle("15w.com", func(raw_url string, u *url.URL) error {
		switch strings.SplitN(u.Host, ".", 2)[0] {
		case "lol":
			SaveList(LOLList)
		case "ow":
			SaveList(OWList)
		case "csgo":
			SaveList(CSGOList)
		case "me":
			SaveList(MEList)
		case "dota2":
			SaveList(DOTA2List)
		default:
			return common.NewSelfError("can not support:" + raw_url)
		}
	})
}
