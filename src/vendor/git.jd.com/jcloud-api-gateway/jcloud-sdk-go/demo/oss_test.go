package demo

import (
	"testing"
	"fmt"
	. "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/core"
	"git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/oss/client"
	"git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/oss/apis"
)

func initOssClient() *client.OssClient {
	accessKey := "35DDDCFFB86CF2D494F0F3B6B0B3EF68"
	secretKey := "93C107EF1F3A0C46C6329C04F561A29E"
	credentials := NewCredentials(accessKey, secretKey)

	client := client.NewOssClient(credentials)
	config := NewConfig()
	config.SetScheme(SchemeHttp)
	client.SetConfig(config)
	return client
}

func TestListBuckets(t * testing.T) {
	req := apis.NewListBucketsRequestWithoutParam()
	req.SetRegionId("cn-north-1")
	client := initOssClient()
	resp, err := client.ListBuckets(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("oss list bucket failed.", err, resp.Error)
	}

	fmt.Println(resp.RequestID)
	fmt.Println(resp.Result.Buckets)
}

func TestHeadBucket(t *testing.T) {
	req := apis.NewHeadBucketRequestWithAllParams("cn-north-1", "go-sdk")
	client := initOssClient()
	resp, err := client.HeadBucket(req)
	if err != nil {
		t.Error("oss head bucket failed.", err)
	} else {
		fmt.Println(resp.RequestID)
	}
}