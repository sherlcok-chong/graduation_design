package redis

import (
	"GraduationDesign/src/dao/redis/query"
	"context"
	"github.com/go-redis/redis/v8"
)

func Init(Addr, Password string, PoolSize, DB int) *query.Queries {
	rdb := redis.NewClient(&redis.Options{
		Addr:     Addr,     // ip:端口
		Password: Password, // 密码
		PoolSize: PoolSize, // 连接池
		DB:       DB,       // 默认连接数据库
	})
	_, err := rdb.Ping(context.Background()).Result() // 测试连接
	if err != nil {
		panic(err)
	}
	return query.New(rdb)
}
