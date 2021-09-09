package impl

import (
	"CMDB/api/pkg/host"
	"context"
)

func (s *service) SaveHost(ctx context.Context,host *host.Host)(*host.Host,error){
	return nil,nil
}

func (s *service) QueryHost(context.Context, *host.QueryHostRequest) (*host.HostSet, error){
	return nil,nil
}

func (s *service) DescribeHost(context.Context, *host.DescribeHostRequest) (*host.HostSet, error){
	return nil,nil
}

func (s *service)DeleteHost(context.Context, *host.DeleteHostRequest) (*host.Host, error){
	return nil,nil
}