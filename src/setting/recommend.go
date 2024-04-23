package setting

import (
	"GraduationDesign/src/dao"
	"context"
)

type recommend struct {
}

func (r recommend) Init() {
	dao.Group.Redis.ReadRE(context.Background())
}
