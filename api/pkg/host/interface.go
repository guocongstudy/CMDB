package host

import "context"

type Service interface {
	//定义存储的接口
	SaveHost(context.Context,*Host)(*Host,error)
}
//1：22：08