package tx

import (
	db "GraduationDesign/src/dao/mysql/sqlc"
	"GraduationDesign/src/pkg/tool"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (store *SqlStore) UpdateUserAvatarTx(c *gin.Context, fileKey string, fileUrl string, userId int64) error {
	return store.execTx(c, func(queries *db.Queries) error {
		var err error
		err = tool.DoThat(err, func() error {
			err = queries.CreateFile(c, db.CreateFileParams{
				Filename: strconv.Itoa(int(userId)) + "avatar",
				FileKey:  fileKey,
				Url:      fileUrl,
				Userid:   userId,
			})
			return err
		})
		err = tool.DoThat(err, func() error {
			err = queries.UpdateUserAvatar(c, db.UpdateUserAvatarParams{
				Avatar: fileUrl,
				ID:     userId,
			})
			return err
		})
		return err
	})
}
