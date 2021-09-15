package http

import (
	"CMDB/api/pkg"
	"CMDB/api/pkg/host"
	"fmt"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/julienschmidt/httprouter"
)

var (
	api = &handler{}
)

type handler struct {
	service host.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named("Host")
	if pkg.Host == nil {
		return fmt.Errorf("dependence service host not ready")
	}
	h.service = pkg.Host
	return nil
}

func RegistAPI(r *httprouter.Router) {
	api.Config()
	r.GET("/hosts", api.QueryHost)
	r.POST("/hosts", api.CreateHost)
	r.GET("/hosts/:id", api.DescribeHost)
	r.DELETE("/hosts/:id", api.DeleteHost)
	r.PUT("/hosts/:id", api.PutHost)
	r.PATCH("/hosts/:id", api.PatchHost)
}
