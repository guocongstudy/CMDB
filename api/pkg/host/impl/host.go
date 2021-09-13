package impl

import (
	"CMDB/api/pkg/host"
	"context"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/sqlbuilder"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
)

const (
	insertResourceSQL = `INSERT INTO resource (
		id,vendor,region,zone,create_at,expire_at,category,type,instance_id,
		name,description,status,update_at,sync_at,sync_accout,public_ip,
		private_ip,pay_type,describe_hash,resource_hash
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	insertHostSQL = `INSERT INTO host (
		resource_id,cpu,memory,gpu_amount,gpu_spec,os_type,os_name,
		serial_number,image_id,internet_max_bandwidth_out,
		internet_max_bandwidth_in,key_pair_name,security_groups
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?);`
	updateResourceSQL = `UPDATE resource SET 
		expire_at=?,category=?,type=?,name=?,description=?,
		status=?,update_at=?,sync_at=?,sync_accout=?,
		public_ip=?,private_ip=?,pay_type=?,describe_hash=?,resource_hash=?
	WHERE id = ?`
	updateHostSQL = `UPDATE host SET 
		cpu=?,memory=?,gpu_amount=?,gpu_spec=?,os_type=?,os_name=?,
		image_id=?,internet_max_bandwidth_out=?,
		internet_max_bandwidth_in=?,key_pair_name=?,security_groups=?
	WHERE resource_id = ?`

	queryHostSQL      = `SELECT * FROM resource as r LEFT JOIN host h ON r.id=h.resource_id`
	deleteHostSQL     = `DELETE FROM host WHERE resource_id = ?;`
	deleteResourceSQL = `DELETE FROM resource WHERE id = ?;`
)


func (s *service) SaveHost(ctx context.Context, h *host.Host) (*host.Host, error) {
	h.Id = xid.New().String()
	h.ResourceId=h.Id
	h.SyncAt=ftime.Now().Timestamp()
	if err :=s.save(ctx,h);err !=nil{
		return nil,err
	}
	return h,nil
}

func (s *service) QueryHost(ctx context.Context, req *host.QueryHostRequest) (*host.HostSet, error) {
	query :=sqlbuilder.NewQuery(queryHostSQL)

	querySQL,args :=query.Order("synv_at").Desc().Limit(req.Offset(),uint(req.PageSize)).BuildQuery()
    s.log.Debugf("sql:%s",querySQL)

	queryStmt,err :=s.db.Prepare(querySQL)
	if err !=nil{
		return nil, exception.NewInternalServerError("prepare query host error,%s",err.Error())
	}
	defer queryStmt.Close()

	rows,err :=queryStmt.Query(args...)
	if err !=nil{
		return nil,exception.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	set := host.NewHostSet()
	for rows.Next(){
		ins :=host.NewDefaultHost()
		err :=rows.Scan(

			)
	}

}

func (s *service) DescribeHost(context.Context, *host.DescribeHostRequest) (*host.HostSet, error) {
	return nil, nil
}

func (s *service) DeleteHost(context.Context, *host.DeleteHostRequest) (*host.Host, error) {
	return nil, nil
}
