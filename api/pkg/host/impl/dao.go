package impl

import (
	"CMDB/api/pkg/host"
	"context"
	"database/sql"
)

func (s *service) save(ctx context.Context, h *host.Host) error {
	var (
		stmt *sql.Stmt
		err  error
	)
	//开启一个事物
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
   //执行结果提交或者回滚事物
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	//生成描写信息的Hash
	if err :=h.GenHash();err !=nil{
		return err
	}

	_, err = stmt.Exec(
		h.Id, h.Vendor, h.Region, h.Zone, h.CreateAt, h.ExpireAt, h.Category, h.InstanceId,
		h.Name, h.Description, h.Status, h.UpdateAt, h.SyncAt, h.SyncAccount, h.PublicIP,
		h.PrivateIP, h.PayType, h.Description, h.ResourceHash,
	)
	if err != nil {
		return err
	}

	//避免SQL注入，请使用Prepare
	stmt, err = tx.Prepare(insertResourceSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		h.Id, h.Vendor, h.Region, h.Zone, h.CreateAt, h.ExpireAt, h.Category, h.InstanceId,
		h.Name, h.Description, h.Status, h.UpdateAt, h.SyncAt, h.SyncAccount, h.PublicIP,
		h.PrivateIP, h.PayType, h.DescribeHash, h.ResourceHash,
	)
	if err != nil {
		return err
	}

	return tx.Commit()

}
