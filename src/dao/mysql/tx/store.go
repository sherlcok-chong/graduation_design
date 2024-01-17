package tx

import (
	db "GraduationDesign/src/dao/mysql/sqlc"
	"context"
	"database/sql"
	"fmt"
)

type TXer interface {
}
type SqlStore struct {
	*db.Queries
	DB *sql.DB
}

// 通过事务执行回调函数
func (store *SqlStore) execTx(ctx context.Context, fn func(queries *db.Queries) error) error {
	tx, err := store.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := store.WithTx(tx) // 使用开启的事务创建一个查询
	if err := fn(q); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err:%v,rb err:%v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
