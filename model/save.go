/*go**************************************************************************
 File            : save.go
 Subsystem       : model
 Author          : yijunjun
 Date&Time       : 2017-06-05
 Description     : 数据模型-保存
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package model

import (
	"strings"
)

func SaveList(_ *TList, err error) error {
	return err
}

func SavePage(tp *TPage, err error) error {
	tp.Author = strings.TrimSpace(tp.Author)
	tp.Source = strings.TrimSpace(tp.Source)
	tp.Date = strings.TrimSpace(tp.Date)
	return err
}
