package sdk

import (
	"fmt"
	"testing"

	"github.com/astaxie/beego/httplib"
)

func TestNewRongCloud(t *testing.T) {
	rcex := RongCloudExtra{
		RongCloudURI:    "http://172",
		RongCloudSMSURI: "123",
		TimeOut:         30,
	}
	rc := NewRongCloud(
		"输入用户app key",
		"输入用户app secret",
		&rcex,
	)
	fmt.Println(rc)
}

func TestRongCloud_FillHeader(t *testing.T) {
	rc := NewRongCloud(
		"输入用户app key",
		"输入用户app secret",
		nil,
	)
	req := httplib.Post(rc.RongCloudURI)
	rc.FillHeader(req)
	fmt.Println(req.GetRequest().Header)
}
