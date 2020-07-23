package demo

import (
	"fmt"
	"testing"
	"time"
	. "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/core"
	. "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/vm/apis"
	. "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/vm/client"
	"git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/common/models"
	vm "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/vm/models"
	vpc "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/vpc/models"
)

func initVmClient() *VmClient {
	accessKey := "35DDDCFFB86CF2D494F0F3B6B0B3EF68"
	secretKey := "93C107EF1F3A0C46C6329C04F561A29E"
	credentials := NewCredentials(accessKey, secretKey)

	return NewVmClient(credentials)
}

func TestVmGetInstances(t *testing.T) {
	client := initVmClient()
	req := NewDescribeInstancesRequestWithoutParam()
	req.SetRegionId("cn-north-1")
	resp, err := client.DescribeInstances(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe instance failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.Instances)
	fmt.Println(resp.Result.TotalCount)
	fmt.Println(len(resp.Result.Instances))
}

func TestVmGetInstancesByPage(t *testing.T) {
	client := initVmClient()
	req := NewDescribeInstancesRequestWithoutParam()
	req.SetRegionId("cn-north-1")
	req.SetPageNumber(1)
	req.SetPageSize(10)
	resp, err := client.DescribeInstances(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe instance failed. ", err, resp.Error.Code, resp.Error.Message)
	}

	if len(resp.Result.Instances) > 10 {
		t.Error("describe instances by paging failed.", len(resp.Result.Instances))
	}
	fmt.Println(resp.Result.Instances)
	fmt.Println(resp.Result.TotalCount)
	fmt.Println(len(resp.Result.Instances))
}

func TestVmGetInstanceTypes(t *testing.T) {
	client := initVmClient()
	req := NewDescribeInstanceTypesRequestWithAllParams("cn-north-1", nil)
	resp, err := client.DescribeInstanceTypes(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe instance types failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.InstanceTypes)
}

func TestVmCreateInstance(t *testing.T) {
	instanceId := createInstance(t)
	testVmGetInstancesWithFilter(t, instanceId)
	try(10, 5*time.Second, stopInstance, t, instanceId)
	try(10, 5*time.Second, deleteInstance, t, instanceId)
}

func createInstance(t *testing.T) string {
	client := initVmClient()

	az := "cn-north-1a"
	diskCategory := "local"
	description  := "golang-sdk"
	instanceType := "g.n2.medium"
    imageId      := "98d44a0f-88c1-451a-8971-f1f769073b6c"
	spec := vm.InstanceSpec{
		Az:           &az,
		InstanceType: &instanceType,
		ImageId:      &imageId,
		Name:         "golang-sdk-test",
		PrimaryNetworkInterface: &vm.InstanceNetworkInterfaceAttachmentSpec{
			NetworkInterface: &vpc.NetworkInterfaceSpec{SubnetId: "subnet-j0yka3pkcu", Az: &az},
		},
		SystemDisk:  &vm.InstanceDiskAttachmentSpec{DiskCategory: &diskCategory},
		Description: &description,
	}
	req := NewCreateInstancesRequestWithoutParam()
	req.SetRegionId("cn-north-1")
	req.SetInstanceSpec(&spec)
	req.SetMaxCount(1)
	resp, err := client.CreateInstances(req)
	if err != nil {
		t.Error(err.Error())
		return ""
	}

	if resp.Error.Code != 0 {
		t.Error(fmt.Sprintf("create instance failed, code=%d, message=%s", resp.Error.Code, resp.Error.Message))
	}

	return resp.Result.InstanceIds[0]
}

func testVmGetInstancesWithFilter(t *testing.T, instanceId string) {
	client := initVmClient()
	op := "eq"
	filter := models.Filter{"instanceId", &op, []string{instanceId}}
	req := NewDescribeInstancesRequestWithoutParam()
	req.SetRegionId("cn-north-1")
	req.SetFilters([]models.Filter{filter})
	resp, err := client.DescribeInstances(req)
	if err != nil || resp.Result.TotalCount != 1 {
		t.Error("describe instance with filter failed. ", err, resp.Error.Message)
	}
	fmt.Println(resp.Result.Instances)
}

func stopInstance(instanceId string) bool {
	req := NewStopInstanceRequestWithAllParams("cn-north-1", instanceId)
	client := initVmClient()
	resp, err := client.StopInstance(req)
	return err == nil && resp.Error.Code == 0
}

func deleteInstance(instanceId string) bool {
	req := NewDeleteInstanceRequestWithAllParams("cn-north-1", instanceId)
	client := initVmClient()
	resp, err := client.DeleteInstance(req)
	return err == nil && resp.Error.Code == 0
}

func try(count int, wait time.Duration, op func(instanceId string) bool, t *testing.T, instanceId string) {
	for count > 0 {
		result := op(instanceId)
		if result {
			return
		}

		time.Sleep(wait)
		count--
		fmt.Println("Could try times: ", count)
	}
}
