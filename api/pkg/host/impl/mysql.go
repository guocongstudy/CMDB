package impl

import (
	"CMDB/api/pkg/host/conf"
	"database/sql"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	//Service 服务实例
	Service = &service{}
)

type service struct {
	db *sql.DB
	log  logger.Logger
	stmts map[string]sql.Stmt
}

//service 初始化
func (s *service) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	s.log = zap.L().Named("Host")
	s.db = db
	return nil
}
