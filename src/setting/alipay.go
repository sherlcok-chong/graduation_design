package setting

import (
	"GraduationDesign/src/global"
	pay "GraduationDesign/src/pkg/alipay"
)

type alipay struct {
}

func (alipay) Init() {
	global.AliPayClient = pay.Init(pay.Config{
		KAppID:               global.PvSettings.AliPay.KAppID,
		KPrivateKey:          global.PvSettings.AliPay.KPrivateKey,
		IsProduction:         global.PvSettings.AliPay.IsProduction,
		AppPublicCertPath:    global.PvSettings.AliPay.AppPublicCertPath,
		AliPayRootCertPath:   global.PvSettings.AliPay.AliPayRootCertPath,
		AliPayPublicCertPath: global.PvSettings.AliPay.AliPayPublicCertPath,
		NotifyURL:            global.PvSettings.AliPay.NotifyURL,
		ReturnURL:            global.PvSettings.AliPay.ReturnURL,
	})
}
