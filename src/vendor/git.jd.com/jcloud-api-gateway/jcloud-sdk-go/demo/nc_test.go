package demo

import (
	"fmt"
	. "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/core"
	"git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/nc/apis"
	. "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/nc/client"
	"testing"
)

func initNcClient() *NcClient {
	accessKey := "AC78D6FE9F31B6AC24D1A199CA5863FA"
	secretKey := "A7F79E129A9BB36891A4BEAE0B60BEC9"
	credentials := NewCredentials(accessKey, secretKey)

	config := NewConfig()
	config.SetScheme(SchemeHttp)
	config.SetEndpoint("192.168.180.18")

	client := NewNcClient(credentials)
	client.SetConfig(config)
	return client
}

//func TestDescribeQuotaForOperator(t *testing.T) {
//	req := apis.NewDescribeQuotasRequestWithoutParam()
//	req.SetRegionId("cn-north-1")
//	req.AddHeader("x-jdcloud-pin", "jcloud_00")
//	req.AddHeader("x-jdcloud-erp", "DemoUser")
//
//	client := initNcClient()
//	resp, err := client.DescribeQuotas(req)
//	if err != nil || resp.Error.Code != 0 {
//		t.Error("describe quota failed. ", err, resp.Error.Message)
//	}
//	fmt.Println(resp.RequestID)
//	fmt.Println(resp.Result.Quotas)
//}

func TestDescribeContainers(t *testing.T) {
	req := apis.NewDescribeContainersRequestWithAllParams("cn-north-1", nil, nil, nil)
	req.AddHeader("x-jdcloud-pin", "jcloud_00")
	req.AddHeader("x-jdcloud-erp", "DemoUser")

	client := initNcClient()
	resp, err := client.DescribeContainers(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe Containers failed. ", err, resp.Error.Message)
	}
	fmt.Println(resp.RequestID)
	fmt.Println(resp.Result.Containers)
}

//func TestDescribeSecret(t *testing.T) {
//	req := apis.NewDescribeSecretRequest("cn-north-1", "汉")
//	req.AddHeader("x-jdcloud-pin", "jcloud_00")
//	req.AddHeader("x-jdcloud-erp", "DemoUser")
//
//	client := initNcClient()
//	resp, err := client.DescribeSecret(req)
//	if err != nil {
//		t.Error("describe secret failed. ", err)
//	}
//
//	if resp.Error.Code != 0 {
//		t.Error("describe secret failed. ", resp.Error.Message)
//	}
//
//	fmt.Println(resp.RequestID)
//}

func TestModifyQuota(t *testing.T) {
	req := apis.NewModifyQuotaRequestWithAllParams("cn-north-1", "container", 10)
	req.AddHeader("x-jdcloud-pin", "jcloud_00")
	req.AddHeader("x-jdcloud-erp", "DemoUser")

	client := initNcClient()
	resp, err := client.ModifyQuota(req)
	fmt.Println(resp)
	if err != nil {
		t.Error("modify Containers quota failed. ", err)
	}
	fmt.Println(resp.RequestID)
}

//func TestGetLogs(t *testing.T) {
//	req := apis.NewGetLogsRequest("cn-north-1", "c-6pvrwf3yyd", 100, 1000, 2000)
//	req.AddHeader("x-jdcloud-pin", "jcloud_00")
//	req.AddHeader("x-jdcloud-erp", "DemoUser")
//
//	client := initNcClient()
//	resp, err := client.GetLogs(req)
//	if err != nil || resp.Error.Code != 0 {
//		t.Error("Get container logs failed. ", err, resp.Error.Message)
//	}
//	fmt.Println(resp.RequestID)
//	fmt.Println(resp.Result.Logs)
//}
