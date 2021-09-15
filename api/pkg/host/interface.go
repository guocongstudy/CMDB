package host

import (
	"context"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)
//全局生成一个validate的实例
var (
	validate = validator.New()
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
func NewQueryHostRequestFromHTTP(r *http.Request) *QueryHostRequest {
	qs := r.URL.Query()

	ps := qs.Get("page_size")
	pn := qs.Get("page_number")
	kw := qs.Get("keywords")

	psUint64, _ := strconv.ParseUint(ps, 10, 64)
	pnUint64, _ := strconv.ParseUint(pn, 10, 64)

	if psUint64 == 0 {
		psUint64 = 20 //pagesize 默认是20
	}
	if pnUint64 == 0 {
		pnUint64 = 1
	}
	return &QueryHostRequest{
		PageSize:   psUint64,
		PageNumber: pnUint64,
		Keywords:   kw,
	}
}

type QueryHostRequest struct {
	//查询主机需要的参数
	PageSize   uint64 `json:"page_size,omitempty"` //一页20条，omitempty 表示传入的值为空则不输出
	PageNumber uint64 `json:"page_number,omitempty"`
	Keywords   string `json:"keywords"`
}

//pageNumber 第一页 第二页。。。
func (req *QueryHostRequest) Offset() int64 {
	return int64(req.PageSize) *int64(req.PageNumber-1)  //从0开始
}



func NewDescribeHostRequestWithID(id string) *DescribeHostRequest {
	return &DescribeHostRequest{
		Id: id,
	}
}


type DescribeHostRequest struct {
	//利用主机Id
	Id string `json:"id" validate:"required"`
}

func NewDeleteHostRequestWithID(id string) *DeleteHostRequest {
	return &DeleteHostRequest{Id: id}
}

type DeleteHostRequest struct {
	Id string `json:"id" validate:"required"` //参数校验，必传
}

type UpdateMode int

const (
	PUT UpdateMode = iota
	PATCH
)

func NewUpdateHostRequest(id string) *UpdateHostRequest {
	return &UpdateHostRequest{
		Id:             id,
		UpdateMode:     PUT,
		UpdateHostData: &UpdateHostData{},
	}
}

type UpdateHostData struct {
	*Resource
	*Describe
}

type UpdateHostRequest struct {
	Id             string          `json:"id" validate:"required"` //参数校验，必传
	UpdateMode     UpdateMode      `json:"update_mode"`
	UpdateHostData *UpdateHostData `json:"date" validate:"required"` //参数校验，必传
}

func (req *UpdateHostRequest) Validate() error {
	//用validate校验结构体
	return validate.Struct(req)
}


