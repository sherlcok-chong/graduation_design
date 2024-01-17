package setting

import (
	"GraduationDesign/src/global"
	"github.com/0RAJA/Rutils/pkg/goroutine/work"
)

type worker struct {
}

func (worker) Init() {
	global.Worker = work.Init(work.Config{
		TaskChanCapacity:   global.PbSettings.Worker.TaskChanCapacity,
		WorkerChanCapacity: global.PbSettings.Worker.WorkerChanCapacity,
		WorkerNum:          global.PbSettings.Worker.WorkerNum,
	})
}
