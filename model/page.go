/*go**************************************************************************
 File            : page.go
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

type TPage struct {
	Id           string
	Author       string
	Source       string
	Date         string
	HitsCount    string
	CommentCount string
}
