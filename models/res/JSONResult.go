package res

import (
	"time"
)

type Result struct {
	Sucess    bool        `json:"sucess"`
	Message   string      `json:"message"`
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

func (res *Result) OK() Result {
	res.Sucess = true
	res.Message = "操作成功"
	res.Code = 200
	res.Data = nil
	res.Timestamp = time.Now().Unix()
	return *res
}

func (res *Result) Ok(message string) Result {
	res.Sucess = true
	res.Message = message
	res.Code = 200
	res.Data = nil
	res.Timestamp = time.Now().Unix()
	return *res
}

func (res *Result) SUCESS(param interface{}) Result {
	res.Sucess = true
	res.Message = "操作成功"
	res.Code = 200
	res.Data = param
	res.Timestamp = time.Now().Unix()
	return *res
}

func (res *Result) ERROR() Result {
	res.Sucess = false
	res.Message = "操作失败"
	res.Code = 500
	res.Data = nil
	res.Timestamp = time.Now().Unix()
	return *res
}

func (res *Result) Error(message string) Result {
	res.Sucess = false
	res.Message = message
	res.Code = 500
	res.Data = nil
	res.Timestamp = time.Now().Unix()
	return *res
}

func (res *Result) Fail(param interface{}) Result {
	res.Sucess = false
	res.Message = "操作失败"
	res.Code = 500
	res.Data = param
	res.Timestamp = time.Now().Unix()
	return *res
}
