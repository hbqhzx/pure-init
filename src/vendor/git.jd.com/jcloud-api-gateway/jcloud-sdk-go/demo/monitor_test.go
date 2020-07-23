package demo

import (
	"fmt"
	"testing"
	. "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/core"
	. "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/monitor/client"
	"git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/monitor/apis"
)

func initMonitorClient() *MonitorClient {
	accessKey := "35DDDCFFB86CF2D494F0F3B6B0B3EF68"
	secretKey := "93C107EF1F3A0C46C6329C04F561A29E"
	credentials := NewCredentials(accessKey, secretKey)

	return NewMonitorClient(credentials)
}

func TestDescribeMetrics(t *testing.T) {
	req := apis.NewDescribeMetricsRequestWithoutParam()
	req.SetServiceCode("vm")
	client := initMonitorClient()
	resp, err := client.DescribeMetrics(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("monitor describe metrics failed.", err, resp.Error.Message)
	}
	fmt.Println(resp.Result.Metrics)
}
