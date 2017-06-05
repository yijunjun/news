/*go**************************************************************************
 File            : site.go
 Subsystem       : common
 Author          : yijunjun
 Date&Time       : 2017-06-05
 Description     : 公共功能-站点分发处理
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package common

import (
	"errors"
	"net/url"
	"strings"
)

type SiteHandle func(string, *url.URL) error

var g_site_handles = map[string]SiteHandle{}

func RegSiteHandle(site string, handle SiteHandle) {
	g_site_handles[site] = handle
}

func FetchSite(page_url string) error {
	u, err := url.Parse(page_url)
	if err != nil {
		return err
	}

	for s, h := range g_site_handles {
		if strings.Contains(u.Host, s) {
			return h(page_url, u)
		}
	}

	return NewSelfError("can not support: " + page_url)
}
