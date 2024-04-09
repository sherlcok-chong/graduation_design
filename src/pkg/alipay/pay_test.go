package alipay_test

import (
	"GraduationDesign/src/global"
	pay "GraduationDesign/src/pkg/alipay"
	"log"
	"net/http"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/smartwalle/xid"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

var AliPayClient *pay.Client

func ExampleInit() {
	AliPayClient = pay.Init(pay.Config{
		KAppID:               global.PvSettings.AliPay.KAppID,
		KPrivateKey:          global.PvSettings.AliPay.KPrivateKey,
		IsProduction:         global.PvSettings.AliPay.IsProduction,
		AppPublicCertPath:    global.PvSettings.AliPay.AppPublicCertPath,
		AliPayRootCertPath:   global.PvSettings.AliPay.AliPayRootCertPath,
		AliPayPublicCertPath: global.PvSettings.AliPay.AliPayPublicCertPath,
		NotifyURL:            global.PvSettings.AliPay.NotifyURL,
		ReturnURL:            global.PvSettings.AliPay.ReturnURL,
	})
	var s = gin.Default()
	s.GET("/alipay", payUrl)
	s.GET("/callback", callback)
	s.POST("/notify", notify)
	if err := s.Run(":8080"); err != nil {
		panic(err)
	}
}
func TestMain(m *testing.M) {
	ExampleInit()
}
func payUrl(c *gin.Context) {
	orderID := strconv.FormatInt(xid.Next(), 10)
	url, err := AliPayClient.Pay(pay.Order{
		ID:          orderID,
		Subject:     "ttms购票:" + orderID,
		TotalAmount: 30,
		Code:        pay.LaptopWebPay,
	})
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, "系统错误")
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func callback(c *gin.Context) {
	_ = c.Request.ParseForm() // 解析form
	orderID, err := AliPayClient.VerifyForm(c.Request.Form)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, "校验失败")
		return
	}
	c.JSON(http.StatusOK, "支付成功:"+orderID)
}

func notify(c *gin.Context) {
	_ = c.Request.ParseForm() // 解析form
	orderID, err := AliPayClient.VerifyForm(c.Request.Form)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("支付成功:" + orderID)
	// 做自己的事
}
