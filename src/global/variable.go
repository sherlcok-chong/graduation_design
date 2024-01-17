package global

import (
	"GraduationDesign/src/model/config"
	"GraduationDesign/src/pkg/mark"
	"github.com/0RAJA/Rutils/pkg/goroutine/work"
	"github.com/0RAJA/Rutils/pkg/logger"
	"github.com/0RAJA/Rutils/pkg/token"
)

var (
	Logger     *logger.Log    // 日志
	PbSettings config.Public  // Public配置
	PvSettings config.Private // Private配置
	Maker      token.Maker    // token
	Worker     *work.Worker
	EmailMark  *mark.Mark // 邮箱标记
)
