package v1

import (
	"GraduationDesign/src/dao"
	"GraduationDesign/src/global"
	"GraduationDesign/src/model/request"
	pay "GraduationDesign/src/pkg/alipay"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type alipay struct {
}

const return_url = " // 根据支付状态进行业务处理\n  if (paymentStatus === 'TRADE_SUCCESS') {\n    // 支付成功逻辑，更新订单状态等\n\n    // 生成跳转页面的HTML内容\n    const html = `\n      <html>\n        <head>\n          <meta http-equiv=\"refresh\" content=\"3;url=http:localhost:3000/main/person\">\n        </head>\n        <body>\n          <h1>支付成功！正在跳转回原来的页面...</h1>\n        </body>\n      </html>\n    `;\n\n    res.send(html);\n  } else {\n    // 支付失败逻辑\n\n    // 生成跳转页面的HTML内容\n    const html = `\n      <html>\n        <head>\n          <meta http-equiv=\"refresh\" content=\"3;url=http:localhost:3000/main/person\">\n        </head>\n        <body>\n          <h1>支付失败！正在跳转回原来的页面...</h1>\n        </body>\n      </html>\n    `;\n\n    res.send(html);\n  }"

func (alipay) PayUrl(c *gin.Context) {
	p := &request.Pay{}
	if err := c.ShouldBindQuery(p); err != nil {
		c.JSON(http.StatusOK, "系统错误")
		return
	}
	url, err := global.AliPayClient.Pay(pay.Order{
		ID:          p.OrderID,
		Subject:     "闲置租赁:" + p.OrderID,
		TotalAmount: p.Price,
		Code:        pay.LaptopWebPay,
	})
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, "系统错误")
		return
	}
	c.JSON(http.StatusOK, url)
}

func (alipay) Callback(c *gin.Context) {
	_ = c.Request.ParseForm() // 解析form
	_, err := global.AliPayClient.VerifyForm(c.Request.Form)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, "校验失败")
		return
	}
	c.Redirect(http.StatusFound, "http://8.137.9.66:3000/main/person")
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
	err = dao.Group.Mysql.ChangeStatusByOrderID(c, orderID)
	if err != nil {
		global.Logger.Error(err.Error())
	}
	c.Redirect(http.StatusFound, "http://8.137.9.66:3000/main/person")
}
