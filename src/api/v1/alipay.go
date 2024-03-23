package v1

import (
	"GraduationDesign/src/global"
	pay "GraduationDesign/src/pkg/alipay"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/xid"
	"log"
	"net/http"
	"strconv"
)

type alipay struct {
}

func (alipay) PayUrl(c *gin.Context) {
	orderID := strconv.FormatInt(xid.Next(), 10)
	url, err := global.AliPayClient.Pay(pay.Order{
		ID:          orderID,
		Subject:     "闲置租赁:" + orderID,
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

func (alipay) Callback(c *gin.Context) {
	_ = c.Request.ParseForm() // 解析form
	orderID, err := global.AliPayClient.VerifyForm(c.Request.Form)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, "校验失败")
		return
	}
	c.JSON(http.StatusOK, "支付成功:"+orderID)
}

func (alipay) Notify(c *gin.Context) {
	_ = c.Request.ParseForm() // 解析form
	orderID, err := global.AliPayClient.VerifyForm(c.Request.Form)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("支付成功:" + orderID)
	// 做自己的事
}
