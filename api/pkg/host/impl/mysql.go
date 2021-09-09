package impl

import "github.com/infraboard/mcube/logger"

var (
	//Service 服务实例
	Service =&service{}
)
type service struct {
	log logger.Logger
}
