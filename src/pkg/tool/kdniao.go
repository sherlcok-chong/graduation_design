package tool

import (
	"GraduationDesign/src/global"
	"GraduationDesign/src/model/reply"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//物流状态：2-在途中,3-签收,4-问题件

const (
	ReqUrl             = "https://api.kdniao.com/Ebusiness/EbusinessOrderHandle.aspx"
	SHIP_CODE_YUNDA    = "YTO"
	SHIP_CODE_SHUNFENG = "SF"
	// You can add more, get code from https://view.officeapps.live.com/op/view.aspx?src=http://www.kdniao.com/file/ExpressCode.xls
)

type RequestData struct {
	ShipperCode  string
	LogisticCode string
}

type PostParams struct {
	RequestData string
	EBusinessID string
	RequestType string
	DataSign    string
	DataType    string
}

type TraceResult struct {
	Traces  []reply.TraceItem
	Success bool
	State   string
}

func KdnTraces(shipperCode string, logisticCode string) (traceResult *TraceResult, err error) {

	if requestDataJson, err := json.Marshal(&RequestData{
		ShipperCode:  shipperCode,
		LogisticCode: logisticCode,
	}); err != nil {
		fmt.Printf("KdnTraces request to json error:%v\n", err)
		return nil, err
	} else {
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(string(requestDataJson) + global.PvSettings.KDN.AppKey))
		b64 := base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(md5Ctx.Sum(nil))))

		resp, err := http.PostForm(ReqUrl,
			url.Values{
				"RequestData": {url.QueryEscape(string(requestDataJson))},
				"EBusinessID": {global.PvSettings.KDN.EBusinessID},
				"RequestType": {"1002"},
				"DataSign":    {url.QueryEscape(b64)},
				"DataType":    {"2"},
			})

		if err != nil {
			fmt.Printf("KdnTraces post error:%v\n", err)
			return nil, err
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("KdnTraces read body error:%v\n", err)
				return nil, err
			}

			fmt.Println(string(body))
			// Parser body
			traceResult := TraceResult{}
			json.Unmarshal(body, &traceResult)

			fmt.Printf("Trace result:%v\n", traceResult)

			return &traceResult, nil
		}
	}
}
