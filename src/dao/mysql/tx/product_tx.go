package tx

import (
	db "GraduationDesign/src/dao/mysql/sqlc"
	"GraduationDesign/src/pkg/tool"
	"github.com/gin-gonic/gin"
)

type Product struct {
	ID     int64    `json:"id"`
	UserID int64    `json:"user_id"`
	Price  int32    `json:"price"`
	Texts  string   `json:"texts"`
	IsFree bool     `json:"is_free"`
	IsLend bool     `json:"is_lend"`
	Media  []string `json:"media"`
	Tags   []string `json:"tags"`
}

func (store *SqlStore) GetUserLendProductTx(c *gin.Context, userId int64) ([]Product, error) {
	ps := make([]Product, 0)
	return ps, store.execTx(c, func(queries *db.Queries) error {
		var err error
		var data []db.Commodity
		err = tool.DoThat(err, func() error {
			data, err = queries.GetUserLendProduct(c, userId)
			return err
		})
		err = tool.DoThat(err, func() error {
			for _, v := range data {
				p := Product{
					ID:     v.ID,
					UserID: v.UserID,
					Price:  v.Price,
					Texts:  v.Texts,
					IsFree: v.IsFree,
					IsLend: v.IsLend,
					Media:  nil,
					Tags:   make([]string, 0),
				}
				ids, err := queries.GetProductMediaId(c, v.ID)
				if err != nil {
					return err
				}
				for _, v := range ids {
					url, err := queries.GetFileByID(c, v)
					if err != nil {
						return err
					}
					p.Media = append(p.Media, url)
				}
				tags, err := queries.GetProductTags(c, v.ID)
				if err != nil {
					return err
				}
				p.Tags = tags
				ps = append(ps, p)
			}
			return err
		})
		return err
	})
}
