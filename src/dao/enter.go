package dao

import (
	"GraduationDesign/src/dao/mysql"
)

type group struct {
	Mysql mysql.DB
}

var Group = new(group)
