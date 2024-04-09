package logic

import (
	"GraduationDesign/src/dao"
	"GraduationDesign/src/global"
	"context"
	"github.com/0RAJA/Rutils/pkg/goroutine/task"
)

type auto struct {
}

func (auto) Work() {
	ctx := context.Background()
	deleteExpiredFileTask := task.Task{
		Name:            "deleteExpiredFile",
		Ctx:             ctx,
		TaskDuration:    global.PbSettings.Auto.DeleteExpiredFileDuration,
		TimeoutDuration: global.PbSettings.Server.DefaultContextTimeout,
		F:               DeleteExpiredFile(),
	}
	startTask(deleteExpiredFileTask)
}

func startTask(tasks ...task.Task) {
	for i := range tasks {
		task.NewTickerTask(tasks[i])
	}
}

func DeleteExpiredFile() task.DoFunc {
	return func(parentCtx context.Context) {
		global.Logger.Info("auto task run : deleteExpiredFile")
		ctx, cancel := context.WithTimeout(parentCtx, global.PbSettings.Server.DefaultContextTimeout)
		defer cancel()
		data, err := dao.Group.Mysql.GetExpiredFileID(ctx)
		if err != nil {
			global.Logger.Error(err.Error())
			return
		}
		for _, v := range data {
			_, err := global.OSS.DeleteFile(v.FileKey)
			if err != nil {
				global.Logger.Error(err.Error())
				return
			}
			err = dao.Group.Mysql.DeleteFileByID(ctx, v.ID)
			if err != nil {
				global.Logger.Error(err.Error())
				return
			}
		}
	}
}
