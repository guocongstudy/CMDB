package host

import (
	"context"
)

type Service interface {
	//定义存储的接口
	SaveHost(context.Context, *Host) (*Host, error)
	//定义接口的查询分页的请求
	QueryHost(context.Context, *QueryHostRequest) (*HostSet, error)
	//查询详情(从前端来讲最直观的就是列表的详情页)
	DescribeHost(context.Context, *DescribeHostRequest) (*HostSet, error)
	//支持删除(这里其实是有冗余的，可以共用一个id接口)
	DeleteHost(context.Context, *DeleteHostRequest) (*Host, error)
	//更新数据接口
	UpdateHost(context.Context, *UpdateHostRequest) (*Host, error)
}

type QueryHostRequest struct {
	//查询主机需要的参数
	PageSize   uint64 `json:"page_size,omitempty"` //一页20条，omitempty 表示传入的值为空则不输出
	PageNumber uint64 `json:"page_number,omitempty"`
}

type HostSet struct {
	Items []*Host `json:"items"`
	Total int     `json:"total"`
}

type DescribeHostRequest struct {
	//利用主机Id
	Id string `json:"id"`
}

type DeleteHostRequest struct {
	Id string `json:"id" validate:"required"` //参数校验，必传
}

type UpdateMode int

const (
	PUT UpdateMode = iota
	PATCH
)

type UpdateHostData struct {
	*Resource
	*Describe
}

type UpdateHostRequest struct {
	Id             string          `json:"id" validate:"required"` //参数校验，必传
	UpdateMode     UpdateMode      `json:"update_mode"`
	UpdateHostDate *UpdateHostData `json:"date" validate:"required"` //参数校验，必传
}

//1：22：08y
