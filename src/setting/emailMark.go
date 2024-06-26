package setting

import (
	"GraduationDesign/src/global"
	"GraduationDesign/src/pkg/mark"
	"github.com/0RAJA/Rutils/pkg/email"
)

type emailMark struct {
}

func (emailMark) Init() {
	autoConf := global.PbSettings.Rule
	emailConf := global.PvSettings.Email
	global.EmailMark = mark.New(mark.Config{
		UserMarkDuration: autoConf.UserMarkDuration,
		CodeMarkDuration: autoConf.CodeMarkDuration,
		SMTPInfo: email.SMTPInfo{
			Host:     emailConf.Host,
			Port:     emailConf.Port,
			IsSSL:    emailConf.IsSSL,
			UserName: emailConf.UserName,
			Password: emailConf.Password,
			From:     emailConf.From,
		},
		AppName: global.PbSettings.App.Name,
	})
}
