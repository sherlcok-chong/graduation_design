package setting

import (
	"GraduationDesign/src/dao"
	"GraduationDesign/src/dao/mysql"
	"GraduationDesign/src/global"
)

type mDao struct {
}

func (m mDao) Init() {
	dao.Group.Mysql = mysql.Init(global.PvSettings.Mysql.DriverName, global.PvSettings.Mysql.SourceName)
}
