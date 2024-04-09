package v1

import (
	"GraduationDesign/src/logic"
	"github.com/0RAJA/Rutils/pkg/app"
	"github.com/gin-gonic/gin"
)

type tags struct {
}

func (tags) GetAllTags(c *gin.Context) {
	rly := app.NewResponse(c)
	data, err := logic.Group.Tags.GetAllTags(c)
	rly.Reply(err, data)
}
