package setting

import (
	"GraduationDesign/src/dao"
	"GraduationDesign/src/dao/mysql"
	"GraduationDesign/src/dao/redis"
	"GraduationDesign/src/global"
)

type mDao struct {
}

func (m mDao) Init() {
	dao.Group.Mysql = mysql.Init(global.PvSettings.Mysql.DriverName, global.PvSettings.Mysql.SourceName)
	dao.Group.Redis = redis.Init(
		global.PvSettings.Redis.Address,
		global.PvSettings.Redis.Password,
		global.PvSettings.Redis.PoolSize,
		global.PvSettings.Redis.DB,
	)
}
