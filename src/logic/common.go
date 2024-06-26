package logic

import (
	"GraduationDesign/src/global"
	"GraduationDesign/src/model"
	"GraduationDesign/src/pkg/retry"
	"github.com/0RAJA/Rutils/pkg/token"
)

// 尝试重试
// 失败: 打印日志
func reTry(name string, f func() error) {
	go func() {
		d := global.PbSettings.Auto.Retry.Duration
		times := global.PbSettings.Auto.Retry.MaxTimes
		report := <-retry.NewTry(name, f, d, times).Run()
		global.Logger.Error(report.Error())
	}()
}

// 新建token
// 成功: 返回 token，*token.Payload
// 失败: 返回 nil，error
func newToken(t model.TokenType, id int64) (string, *token.Payload, error) {
	duration := global.PvSettings.Token.UserTokenDuration
	if t == model.AccountToken {
		duration = global.PvSettings.Token.AccountTokenDuration
	}
	data, err := model.NewTokenContent(t, id).Marshal()
	if err != nil {
		return "", nil, err
	}
	result, payload, err := global.Maker.CreateToken(data, duration)
	if err != nil {
		return "", nil, err
	}
	return result, payload, nil
}

// 将id从小到大排序返回
func sortID(id1, id2 int64) (_, _ int64) {
	if id1 > id2 {
		return id2, id1
	}
	return id1, id2
}
