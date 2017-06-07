/*go**************************************************************************
 File            : download.go
 Subsystem       : common
 Author          : yijunjun
 Date&Time       : 2017-06-07
 Description     : 公共功能-下载网页
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package common

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
)

func Download(raw_url string, opts ...string) (*goquery.Document, error) {
	if len(opts) == 0 {
		return goquery.NewDocument(raw_url)
	}

	resp, err := http.Get(raw_url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 第一个可选参数为网页编码(非utf8)
	gbk := mahonia.NewDecoder(opts[0])

	return goquery.NewDocumentFromReader(gbk.NewReader(resp.Body))
}
