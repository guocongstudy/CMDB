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
func (s *service) delete(ctx context.Context, req *host.DeleteHostRequest) error {
	var (
		stmt *sql.Stmt
		err  error
	)

	// 开启一个事物
	// 文档请参考: http://cngolib.com/database-sql.html#db-begintx
	// 关于事物级别可以参考文章: https://zhuanlan.zhihu.com/p/117476959
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// 执行结果提交或者回滚事务
	// 当使用sql.Tx的操作方式操作数据后，需要我们使用sql.Tx的Commit()方法显式地提交事务，
	// 如果出错，则可以使用sql.Tx中的Rollback()方法回滚事务，保持数据的一致性
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	stmt, err = tx.Prepare(deleteHostSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.Id)
	if err != nil {
		return err
	}

	stmt, err = s.db.Prepare(deleteResourceSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.Id)
	if err != nil {
		return err
	}

	return tx.Commit()
}
