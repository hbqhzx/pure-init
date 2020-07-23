package demo

import (
	"fmt"
	"testing"

	. "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/core"
	common "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/common/models"
	. "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/lb/apis"
	. "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/lb/client"
	"git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/lb/models"
)

var (
	backendid     string
	listenerId    string
	targetGroupId string
)

const (
	LOADBALANCER_ID        = "alb-3l868osj4z"
	LOADBALANCER_NAME      = "testlb111"
	LB_SUBNET_ID           = "subnet-6ksoujomny"
	AZ_NORTH_1             = "cn-north-1a"
	ASSOC_ELASTIC_IP_ID    = "fip-n5435qc5an"
	ASSOC_SECURITYGROUP_ID = "sg-63pe0qw6c8"
	LISTENER_NAME          = "test-listener"
	REGION_ID              = "cn-north-1"
	INSTANCE_ID            = "i-s3y06yl79n"
)

func initLbClient() *LbClient {
	accessKey := "59DF21C0FD18D4F1C19FD5A96E4F7CBE"
	secretKey := "8B635A090571F339C6DECA2E65F6F1C6"
	credentials := NewCredentials(accessKey, secretKey)

	return NewLbClient(credentials)
}

func TestLbGetLoadbalancers(t *testing.T) {
	client := initLbClient()
	req := NewDescribeLoadBalancersRequest("cn-north-1")
	req.SetPageNumber(1)
	req.SetPageSize(10)
	resp, err := client.DescribeLoadBalancers(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe loadbalancers failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.LoadBalancers)
	fmt.Println(resp.Result.TotalCount)
	fmt.Println(len(resp.Result.LoadBalancers))
}

func TestLbGetLoadbalancer(t *testing.T) {
	client := initLbClient()
	req := NewDescribeLoadBalancerRequest("cn-north-1", LOADBALANCER_ID)
	resp, err := client.DescribeLoadBalancer(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe loadbalancer failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.LoadBalancer)
}

func TestLbCreateLoadbalancer(t *testing.T) {
	client := initLbClient()
	azs := []string{AZ_NORTH_1}
	req := NewCreateLoadBalancerRequest("cn-north-1", LOADBALANCER_NAME, LB_SUBNET_ID, azs)
	req.SetType("alb")
	resp, err := client.CreateLoadBalancer(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("create loadbalancer failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.LoadBalancerId)
}

func TestLbModifyLoadbalancer(t *testing.T) {
	client := initLbClient()
	req := NewUpdateLoadBalancerRequest("cn-north-1", LOADBALANCER_ID)
	req.SetLoadBalancerName("sdaldj1")
	resp, err := client.UpdateLoadBalancer(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("update loadbalancer failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.RequestID)
}

func TestLbDeleteLoadbalancer(t *testing.T) {
	client := initLbClient()
	req := NewDeleteLoadBalancerRequest("cn-north-1", LOADBALANCER_ID)
	resp, err := client.DeleteLoadBalancer(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("delete loadbalancer failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.RequestID)
}

func TestLbAssociateElasticIp(t *testing.T) {
	client := initLbClient()
	req := NewAssociateElasticIpRequest("cn-north-1", LOADBALANCER_ID, ASSOC_ELASTIC_IP_ID)
	resp, err := client.AssociateElasticIp(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("associate elasticIp failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.RequestID)
}

func TestLbDisassociateElasticIp(t *testing.T) {
	client := initLbClient()
	req := NewDisassociateElasticIpRequest("cn-north-1", LOADBALANCER_ID, ASSOC_ELASTIC_IP_ID)
	resp, err := client.DisassociateElasticIp(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("disassociate elasticIp failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.RequestID)
}

func TestLbAssociateSecurityGroup(t *testing.T) {
	client := initLbClient()
	sgs := []string{ASSOC_SECURITYGROUP_ID}
	req := NewAssociateSecurityGroupRequest("cn-north-1", LOADBALANCER_ID, sgs)
	resp, err := client.AssociateSecurityGroup(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("associate securityGroup failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.RequestID)
}

func TestLbDisassociateSecurityGroup(t *testing.T) {
	client := initLbClient()
	sgs := []string{ASSOC_SECURITYGROUP_ID}
	req := NewDisassociateSecurityGroupRequest("cn-north-1", LOADBALANCER_ID, sgs)
	resp, err := client.DisassociateSecurityGroup(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("disassociate securityGroup failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.RequestID)
}

func TestLbGetQuotaForAlbTarget(t *testing.T) {
	client := initLbClient()
	req := NewDescribeQuotaRequest("cn-north-1", "loadbalancer")
	req.SetLbType("alb")
	req.SetType("target")
	req.SetParentResourceId("xxx")
	resp, err := client.DescribeQuota(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe quota failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.Quota)
}

func TestLbGetQuotaForNlbTarget(t *testing.T) {
	client := initLbClient()
	req := NewDescribeQuotaRequest("cn-north-1", "loadbalancer")
	req.SetLbType("nlb")
	req.SetType("target")
	req.SetParentResourceId("xxx")
	resp, err := client.DescribeQuota(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe quota failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.Quota)
}

func TestLbModifyQuotaForAlb(t *testing.T) {
	client := initLbClient()
	req := NewModifyQuotaRequestWithoutParam()
	req.SetLbType("alb")
	req.SetType("loadbalancer")
	req.SetMaxLimit(123)
	resp, err := client.ModifyQuota(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("modify quota failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.RequestID)
}

func TestLbModifyQuotaForNlbTarget(t *testing.T) {
	client := initLbClient()
	req := NewModifyQuotaRequestWithoutParam()
	req.SetLbType("nlb")
	req.SetType("target")
	req.SetParentResourceId("tg-ybfcxbonzn")
	req.SetMaxLimit(45)
	resp, err := client.ModifyQuota(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("modify quota failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.RequestID)
}


func TestLbGetBackends(t *testing.T) {
	client := initLbClient()
	req := NewDescribeBackendsRequest("cn-north-1")
	req.SetPageNumber(1)
	req.SetPageSize(10)
	resp, err := client.DescribeBackends(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe loadbalancers failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.Backends)
	fmt.Println(resp.Result.TotalCount)
	fmt.Println(len(resp.Result.Backends))
	for _, v := range resp.Result.Backends {
		if v.BackendName == "lsltest" {
			backendid = v.BackendId
		}
	}
}

func TestLbGetBackend(t *testing.T) {
	client := initLbClient()
	req := NewDescribeBackendRequest("cn-north-1", "backend-3fnml2zjot")
	resp, err := client.DescribeBackend(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe backend failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.Backend)
}

func TestLbDeleteBackend(t *testing.T) {
	client := initLbClient()
	req := NewDeleteBackendRequest("cn-north-1", backendid)
	resp, err := client.DeleteBackend(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("delete backend failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result)
}

func TestLbCreateBackend(t *testing.T) {
	client := initLbClient()
	httppath := "/"
	req := NewCreateBackendRequest("cn-north-1", "lsltest", "alb-6xgddfvfpx", "http", 81, &models.HealthCheckSpec{Protocol: "http", HttpPath: &httppath})
	resp, err := client.CreateBackend(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("create backends failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.BackendId)
	backendid = resp.Result.BackendId
}

func TestLbUpdateBackend(t *testing.T) {
	client := initLbClient()
	req := NewUpdateBackendRequest("cn-north-1", "backend-swuu5cmb72")
	//req.SetCloseHealthCheck(true)
	req.SetBackendName("lostit2")
	resp, err := client.UpdateBackend(req)
	fmt.Println(err, resp)
	if err != nil || resp.Error.Code != 0 {
		t.Error("update backends failed. ", err)
	}
}

func TestLbDescribeTargetHealth(t *testing.T) {
	client := initLbClient()
	req := NewDescribeTargetHealthRequest("cn-north-1", "backend-swuu5cmb72")
	resp, err := client.DescribeTargetHealth(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe quota failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.TargetHealths)
}

func TestLbCreateListener(t *testing.T) {
	client := initLbClient()
	req := NewCreateListenerRequest("cn-north-1", LISTENER_NAME, "tcp", 80, "backend-m7xblvk6bg", "alb-e20fbmbgxt")
	req.SetAction("Forward")
	req.SetConnectionIdleTimeSeconds(6000)
	req.SetDescription("wwc-test-listener")

	resp, err := client.CreateListener(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("create listeners failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.ListenerId)
	listenerId = resp.Result.ListenerId
}

func TestLbGetListener(t *testing.T) {
	client := initLbClient()
	req := NewDescribeListernerRequest("cn-north-1", listenerId)
	resp, err := client.DescribeListerner(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe listener failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.Listener)
}

func TestLbUpdateListener(t *testing.T) {
	client := initLbClient()
	req := NewUpdateListenerRequest("cn-north-1", listenerId)
	req.SetStatus("Off")
	req.SetConnectionIdleTimeSeconds(6600)
	req.SetDescription("wwc-listener-after-update")

	resp, err := client.UpdateListener(req)
	fmt.Println(err, resp)
	if err != nil || resp.Error.Code != 0 {
		t.Error("update listener failed. ", err)
	}
}

func TestLbGetListeners(t *testing.T) {
	client := initLbClient()
	req := NewDescribeListernersRequest("cn-north-1")
	req.SetPageNumber(1)
	req.SetPageSize(10)
	filter := common.Filter{}
	filter.Name = "loadBalancerId"
	filter.Values = []string{"alb-e20fbmbgxt"}
	filter2 := common.Filter{}
	filter2.Name = "listenerNames"
	filter2.Values = []string{LISTENER_NAME}
	req.Filters = append(req.Filters, filter)
	req.Filters = append(req.Filters, filter2)

	resp, err := client.DescribeListerners(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe listeners failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.Listeners)
	fmt.Println(resp.Result.TotalCount)
}

func TestLbDeleteListener(t *testing.T) {
	client := initLbClient()
	req := NewDeleteListenerRequest("cn-north-1", listenerId)
	resp, err := client.DeleteListener(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("delete listener failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result)
}
func TestLbCreateTargetGroup(t *testing.T) {
	client := initLbClient()
	req := NewCreateTargetGroupRequest(REGION_ID, "test-target-group-1", LOADBALANCER_ID)
	req.SetDescription("target-group-1")
	resp, err := client.CreateTargetGroup(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("create target group failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	targetGroupId = resp.Result.TargetGroupId
	fmt.Println(resp.Result.TargetGroupId)
}

func TestLbGetTargetGroups(t *testing.T) {
	client := initLbClient()
	req := NewDescribeTargetGroupsRequest(REGION_ID)
	req.SetPageNumber(1)
	req.SetPageSize(10)
	resp, err := client.DescribeTargetGroups(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe target groups failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.TargetGroups)
	fmt.Println(resp.Result.TotalCount)
	fmt.Println(len(resp.Result.TargetGroups))
}

func TestLbGetTargetGroup(t *testing.T) {
	client := initLbClient()
	req := NewDescribeTargetGroupRequest(REGION_ID, targetGroupId)
	resp, err := client.DescribeTargetGroup(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe target group failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.TargetGroup)
}

func TestLbUpdateTargetGroup(t *testing.T) {
	client := initLbClient()
	req := NewUpdateTargetGroupRequest(REGION_ID, targetGroupId)
	req.SetDescription("Test-update")
	resp, err := client.UpdateTargetGroup(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("update target group failed. resp ", err, resp)
	}
}

func TestLbUpdateTargets(t *testing.T) {
	client := initLbClient()

	port := 10
	weight := 10
	targetUpdateSpecs := []models.TargetUpdateSpec{
		{
			TargetId: "",
			Port:     &port,
			Weight:   &weight,
		},
	}

	req := NewUpdateTargetsRequest(targetGroupId, REGION_ID, targetUpdateSpecs)
	resp, err := client.UpdateTargets(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("update targets failed. resp ", err, resp)
	}
}

func TestLbRegisterTargets(t *testing.T) {
	client := initLbClient()
	port := 10
	weight := 10
	targetSpecs := []models.TargetSpec{
		{
			InstanceId: INSTANCE_ID,
			Port:       &port,
			Weight:     &weight,
		},
	}
	req := NewRegisterTargetsRequest(targetGroupId, REGION_ID, targetSpecs)
	resp, err := client.RegisterTargets(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("register targets failed. ", err, resp.Error.Code, resp.Error.Message)
	} else {
		fmt.Println("Register targets success")
	}

	fmt.Println(resp.RequestID)

}

func TestLbDeRegisterTargets(t *testing.T) {
	client := initLbClient()
	targetIds := []string{""}
	req := NewDeRegisterTargetsRequest(REGION_ID, targetGroupId, targetIds)
	resp, err := client.DeRegisterTargets(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("deregister targets failed. ", err, resp.Error.Code, resp.Error.Message)
	} else {
		fmt.Println("Deregister targets success")
	}
}

func TestLbDeleteTargetGroup(t *testing.T) {
	client := initLbClient()
	req := NewDeleteTargetGroupRequest(REGION_ID, targetGroupId)
	resp, err := client.DeleteTargetGroup(req)
	if err != nil {
		t.Error(err.Error())
		fmt.Println("error")
	}
	if resp == nil || resp.Error.Code != 0 {
		t.Error(fmt.Sprintf("delete target group failed, resp %v", resp))
	}
}
