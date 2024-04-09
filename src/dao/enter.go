package dao

import (
	"GraduationDesign/src/dao/mysql"
	"GraduationDesign/src/dao/redis/query"
)

type group struct {
	Mysql mysql.DB
	Redis *query.Queries
}

var Group = new(group)
