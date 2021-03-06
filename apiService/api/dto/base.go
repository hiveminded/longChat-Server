package dto

import (
	"fmt"
)

type BaseRsp struct {
	//StatusCode=0时表示请求成功
	StatusCode int
	//StatusCode !=0时，Error会被详细错误信息填充
	Error string
}

func PasswordNotMatchErrRsp() *BaseRsp {
	return &BaseRsp{
		StatusCode: 1,
		Error:      fmt.Sprintf("password or username not matched!"),
	}
}
func SessionNotFoudErrRsp() *BaseRsp {
	return &BaseRsp{
		StatusCode: 2,
		Error:      fmt.Sprintf("session not found!"),
	}
}
func PostDataErrRsp(dataName string) *BaseRsp {
	return &BaseRsp{
		StatusCode: 1000,
		Error:      fmt.Sprintf("invalid POST data(%s)", dataName),
	}
}

func ParameterErrRsp(params ...string) *BaseRsp {
	s := ""
	if len(params) > 0 {
		for i := range params {
			s += params[i] + ","
		}
		s = s[:len(s)-1]
	}

	return &BaseRsp{
		StatusCode: 1001,
		Error:      fmt.Sprintf("invalid url parameters(%s)", s),
	}
}

func InternalErrRsp() *BaseRsp {
	return &BaseRsp{StatusCode: 1002, Error: "internal server error"}
}

func SuccessRsp() *BaseRsp {
	var succRsp BaseRsp = BaseRsp{StatusCode: 0, Error: ""}
	return &succRsp
}
