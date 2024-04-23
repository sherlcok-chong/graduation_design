package tx

import (
	db "GraduationDesign/src/dao/mysql/sqlc"
	"GraduationDesign/src/model/reply"
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
)

type TXer interface {
	// UpdateUserAvatarTx  更新用户头像
	UpdateUserAvatarTx(c *gin.Context, fileKey string, fileUrl string, userId int64) error
	// GetUserLendProductTx 获取用户出租商品
	GetUserLendProductTx(c *gin.Context, userId int64) ([]reply.ProductInfo, error)
	// GetUserNeedProductTx 需求
	GetUserNeedProductTx(c *gin.Context, userId int64) ([]reply.ProductInfo, error)
	// GetProductInfoTx 需求
	GetProductInfoTx(c *gin.Context, offset, limit int32) ([]reply.ProductInfo, error)
	// GetProductDetailsTX 详情
	GetProductDetailsTX(c *gin.Context, pid, uID int64) (reply.Product, error)
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
