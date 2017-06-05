/*go**************************************************************************
 File            : selferr.go
 Subsystem       : common
 Author          : yijunjun
 Date&Time       : 2017-06-05
 Description     : 公共功能-自定义错误
 Revision        :

 History
 -------


 Copyright (c) Shenzhen Team Blemobi.
**************************************************************************go*/

package common

import (
	"errors"
)

type SelfErr struct {
	s string
}

func NewSelfError(reason string) *SelfErr {
	return &SelfErr{
		s: reason,
	}
}

func (this *SelfErr) Error() string {
	return this.s
}
