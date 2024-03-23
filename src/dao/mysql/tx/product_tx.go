package tx

import (
	db "GraduationDesign/src/dao/mysql/sqlc"
	"GraduationDesign/src/model/reply"
	"GraduationDesign/src/pkg/tool"
	"github.com/gin-gonic/gin"
)

func (store *SqlStore) GetUserLendProductTx(c *gin.Context, userId int64) ([]reply.ProductInfo, error) {
	ps := make([]reply.ProductInfo, 0, 10)
	return ps, store.execTx(c, func(queries *db.Queries) error {
		var err error
		var data []db.GetUserProductInfoRow
		err = tool.DoThat(err, func() error {
			data, err = queries.GetUserProductInfo(c, userId)
			if err != nil {
				return err
			}
			for _, v := range data {
				p := reply.ProductInfo{
					ID:     v.ID,
					Name:   v.Name,
					Price:  v.Price,
					IsFree: v.IsFree,
				}
				err = tool.DoThat(err, func() error {
					d, err := queries.GetUserInfoById(c, v.UserID)
					if err != nil {
						return err
					}
					p.UserName = d.Name
					p.Avatar = d.Avatar
					return err
				})
				err = tool.DoThat(err, func() error {
					d, err := queries.GetProductFirstMedia(c, v.ID)
					if err != nil {
						return err
					}
					id := d.(int64)
					media, err := queries.GetFileByID(c, id)
					if err != nil {
						return err
					}
					p.Media = media
					return err
				})
				if err != nil {
					return err
				}
				ps = append(ps, p)
			}
			return err
		})
		return err
	})
}
func (store *SqlStore) GetUserNeedProductTx(c *gin.Context, userId int64) ([]reply.ProductInfo, error) {
	ps := make([]reply.ProductInfo, 0, 10)
	return ps, store.execTx(c, func(queries *db.Queries) error {
		var err error
		var data []db.GetUserNeedInfoRow
		err = tool.DoThat(err, func() error {
			data, err = queries.GetUserNeedInfo(c, userId)
			if err != nil {
				return err
			}
			for _, v := range data {
				p := reply.ProductInfo{
					ID:    v.ID,
					Name:  v.Name,
					Price: v.Price,
				}
				err = tool.DoThat(err, func() error {
					d, err := queries.GetUserInfoById(c, v.UserID)
					if err != nil {
						return err
					}
					p.UserName = d.Name
					p.Avatar = d.Avatar
					return err
				})
				err = tool.DoThat(err, func() error {
					d, err := queries.GetProductFirstMedia(c, v.ID)
					if err != nil {
						return err
					}
					id := d.(int64)
					media, err := queries.GetFileByID(c, id)
					if err != nil {
						return err
					}
					p.Media = media
					return err
				})
				if err != nil {
					return err
				}
				ps = append(ps, p)
			}
			return err
		})
		return err
	})
}
func (store *SqlStore) GetProductInfoTx(c *gin.Context, offset int32) ([]reply.ProductInfo, error) {
	ps := make([]reply.ProductInfo, 0, 10)
	return ps, store.execTx(c, func(queries *db.Queries) error {
		var err error
		var data []db.GetProductInfoRow
		err = tool.DoThat(err, func() error {
			data, err = queries.GetProductInfo(c, offset)
			if err != nil {
				return err
			}
			for _, v := range data {
				p := reply.ProductInfo{
					ID:    v.ID,
					Name:  v.Name,
					Price: v.Price,
				}
				err = tool.DoThat(err, func() error {
					d, err := queries.GetUserInfoById(c, v.UserID)
					if err != nil {
						return err
					}
					p.UserName = d.Name
					p.Avatar = d.Avatar
					return err
				})
				err = tool.DoThat(err, func() error {
					d, err := queries.GetProductFirstMedia(c, v.ID)
					if err != nil {
						return err
					}
					id := d.(int64)
					media, err := queries.GetFileByID(c, id)
					if err != nil {
						return err
					}
					p.Media = media
					return err
				})
				if err != nil {
					return err
				}
				ps = append(ps, p)
			}
			return err
		})
		return err
	})
}
func (store *SqlStore) GetProductDetailsTX(c *gin.Context, pID, uID int64) (reply.Product, error) {
	ps := reply.Product{}
	return ps, store.execTx(c, func(queries *db.Queries) error {
		var err error
		com, err := queries.GetProductByID(c, pID)
		if err != nil {
			return err
		}
		ps = reply.Product{
			ID:     com.ID,
			Name:   com.Name,
			UserID: com.UserID,
			Price:  com.Price,
			Texts:  com.Texts,
			IsFree: com.IsFree,
			IsLend: com.IsLend,
		}
		ids, err := queries.GetProductMediaId(c, pID)
		if err != nil {
			return err
		}
		for _, v := range ids {
			url, err := queries.GetFileByID(c, v)
			if err != nil {
				return err
			}
			ps.Media = append(ps.Media, url)
		}
		tags, err := queries.GetProductTags(c, pID)
		if err != nil {
			return err
		}
		t := make([]reply.Tags, len(tags))
		for _, v := range tags {
			d := reply.Tags{
				ID:  v.TagID,
				Tag: v.TagName,
			}
			t = append(t, d)
		}
		f, err := queries.CheckUserLike(c, db.CheckUserLikeParams{
			UserID:    uID,
			ProductID: com.ID,
		})
		if err != nil {
			return err
		}
		ps.IsLike = f
		ps.Tags = t
		return err
	})
}
