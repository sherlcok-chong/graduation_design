package mysql

import (
	db "GraduationDesign/src/dao/mysql/sqlc"
	"GraduationDesign/src/dao/mysql/tx"
	"GraduationDesign/src/global"
	"context"
	"database/sql"
)

type DB interface {
	tx.TXer
	db.Querier
}

func Init(driverName, dataSourceName string) DB {
	conn, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), global.PbSettings.Server.DefaultContextTimeout)
	defer cancel()
	if err := conn.PingContext(ctx); err != nil {
		panic(err)
	}
	return &tx.SqlStore{Queries: db.New(conn), DB: conn}
}
