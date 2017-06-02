/*go**************************************************************************
 File            : list.go
 Subsystem       : model
 Author          : yijunjun
 Date&Time       : 2017-06-02
 Description     : 数据模型-列表定义
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package model

type PageInfo struct {
	ImgSrc string
	Url    string
	Title  string
	Date   string
}

type List struct {
	Infos []PageInfo
	Urls  []string
}
