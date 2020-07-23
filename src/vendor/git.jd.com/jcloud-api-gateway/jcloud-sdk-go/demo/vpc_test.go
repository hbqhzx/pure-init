package demo

import (
	//	"time"
	"fmt"
	"testing"

	. "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/core"
	. "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/vpc/apis"
	. "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/vpc/client"
	vpc "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/vpc/models"
)

const (
	ELASTIC_IP_ID        = "fip-6j7ssgqr0o"
	NETWORK_INTERFACE_ID = "port-m2bbrmmaxz"
	SECONDARY_IP1        = "10.0.1.10"
	SECONDARY_IP2        = "10.0.1.11"
	SUBNET_ID            = "subnet-63ndensu4g"
	SECURITYGROUP_ID     = "sg-7lpiabk6lo"

	VPCPEERING_NAME       = "wwc_test_peering"
	VPCPEERING_NAME2      = "wwc_test_peering2"
	PEERING_LOCAL_VPC_ID  = "vpc-4utqaaxvm1"
	PEERING_REMOTE_VPC_ID = "vpc-9z7fy2dskf"
	VPC_ID                = "vpc-yzqu131my1"
	VPEPEERING_MOD_DEL_ID = "vpcpr-9otx9as779"
	SECURITYGROUP_NAME    = "test_sg1"
	VPC_NAME              = "test-vpcads"

	SG_RULE_PROTOCOL_ICMP = 1
	SG_RULE_PROTOCOL_UDP  = 6
	SG_RULE_PROTOCOL_TCP  = 17
	SG_RULE_PROTOCOL_ALL  = 300

	SG_RULE_DIRECTION_INGRESS = 0
	SG_RULE_DIRECTION_EGRESS  = 1
)

var vpcPeeringId = ""
var securityGroupId = ""
var sgRuleId1 = ""
var sgRuleId2 = ""
var vpcId = ""

func initVpcClient() *VpcClient {
	accessKey := "59DF21C0FD18D4F1C19FD5A96E4F7CBE"
	secretKey := "8B635A090571F339C6DECA2E65F6F1C6"
	credentials := NewCredentials(accessKey, secretKey)

	return NewVpcClient(credentials)
}

func TestVpcGetElasticIps(t *testing.T) {
	client := initVpcClient()
	req := NewDescribeElasticIpsRequestWithoutParam()
	req.SetRegionId("cn-north-1")
	req.SetPageNumber(1)
	req.SetPageSize(100)
	resp, err := client.DescribeElasticIps(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe elasticips failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.ElasticIps)
	fmt.Println(resp.Result.TotalCount)
	fmt.Println(len(resp.Result.ElasticIps))
}

func TestVpcGetElasticIp(t *testing.T) {
	client := initVpcClient()
	req := NewDescribeElasticIpRequestWithoutParam()
	req.SetRegionId("cn-north-1")
	req.SetElasticIpId(ELASTIC_IP_ID)
	resp, err := client.DescribeElasticIp(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe elasticip failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.ElasticIp)
}

func TestVpcCreateElasticIps(t *testing.T) {
	client := initVpcClient()
	bandWidth := 2
	provider := "bgp"
	spec := vpc.ElasticIpSpec{
		BandwidthMbps: bandWidth,
		Provider:      provider,
	}
	req := NewCreateElasticIpsRequestWithAllParams("cn-north-1", 1, nil, &spec)
	resp, err := client.CreateElasticIps(req)
	if err != nil {
		t.Error(err.Error())
		fmt.Println("error")
	}

	if resp.Error.Code != 0 {
		t.Error(fmt.Sprintf("create elasticIp failed, code=%d, message=%s", resp.Error.Code, resp.Error.Message))
	}

	//fmt.Println(resp.Result.ElasticIpIds[0])
}
func TestVpcDeleteElasticIp(t *testing.T) {
	client := initVpcClient()
	req := NewDeleteElasticIpRequestWithAllParams("cn-north-1", "fip-i7wpd87kzu")
	resp, err := client.DeleteElasticIp(req)
	if err != nil {
		t.Error(err.Error())
		fmt.Println("error")
	}
	if resp.Error.Code != 0 {
		t.Error(fmt.Sprintf("delete elasticIp failed, code=%d, message=%s", resp.Error.Code, resp.Error.Message))
	}

}

func TestVpcGetnetworkInterfaces(t *testing.T) {
	client := initVpcClient()
	req := NewDescribeNetworkInterfacesRequestWithoutParam()
	req.SetRegionId("cn-north-1")
	req.SetPageNumber(1)
	req.SetPageSize(100)
	resp, err := client.DescribeNetworkInterfaces(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe network interfaces failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.NetworkInterfaces)
	fmt.Println(resp.Result.TotalCount)
	fmt.Println(len(resp.Result.NetworkInterfaces))
}

func TestVpcGetnetworkInterface(t *testing.T) {
	client := initVpcClient()
	req := NewDescribeNetworkInterfaceRequestWithAllParams("cn-north-1", NETWORK_INTERFACE_ID)
	resp, err := client.DescribeNetworkInterface(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe network interface failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.NetworkInterface)
}

func TestVpcAssociateElasticIp(t *testing.T) {
	client := initVpcClient()
	req := NewAssociateElasticIpRequestWithoutParam()
	req.SetRegionId("cn-north-1")
	req.SetElasticIpId(NETWORK_INTERFACE_ID)
	eipId := ELASTIC_IP_ID
	req.ElasticIpId = &eipId
	resp, err := client.AssociateElasticIp(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("associate elastic ip failed. ", err, resp.Error.Code, resp.Error.Message)
	} else {
		fmt.Println("AssociateElasticIp success")
	}

	fmt.Println(resp.RequestID)

}

func TestVpcDisassociateElasticIp(t *testing.T) {
	client := initVpcClient()
	req := NewDisassociateElasticIpRequestWithAllParams("cn-north-1", NETWORK_INTERFACE_ID, nil, nil)
	resp, err := client.DisassociateElasticIp(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("disassociate elastic ip failed. ", err, resp.Error.Code, resp.Error.Message)
	} else {
		fmt.Println("DisassociateElasticIp success")
	}
}

func TestVpcAssignSecondaryIps(t *testing.T) {
	client := initVpcClient()
	req := NewAssignSecondaryIpsRequestWithoutParam()
	req.SetRegionId("cn-north-1")
	req.SetNetworkInterfaceId(NETWORK_INTERFACE_ID)
	req.SecondaryIps = []string{SECONDARY_IP1, SECONDARY_IP2}
	resp, err := client.AssignSecondaryIps(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("assign secondary ips failed. ", err, resp.Error.Code, resp.Error.Message)
	} else {
		fmt.Println("AssignSecondaryIps success")
	}

}

func TestVpcUnassignSecondaryIps(t *testing.T) {
	client := initVpcClient()
	req := NewUnassignSecondaryIpsRequestWithAllParams("cn-north-1", NETWORK_INTERFACE_ID, nil)
	req.SecondaryIps = []string{SECONDARY_IP1, SECONDARY_IP2}
	resp, err := client.UnassignSecondaryIps(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("unassign secondary ips failed. ", err, resp.Error.Code, resp.Error.Message)
	} else {
		fmt.Println("UnassignSecondaryIps success")
	}
}

func TestVpcDescribeSubnets(t *testing.T) {
	client := initVpcClient()
	req := NewDescribeSubnetsRequestWithoutParam()
	req.SetRegionId("cn-north-1")
	req.SetPageNumber(1)
	req.SetPageSize(100)
	resp, err := client.DescribeSubnets(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe subnets failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.Subnets)
	fmt.Println(resp.Result.TotalCount)
	fmt.Println(len(resp.Result.Subnets))
}

func TestVpcDescribeSubnet(t *testing.T) {
	client := initVpcClient()
	req := NewDescribeSubnetRequestWithAllParams("cn-north-1", SUBNET_ID)
	resp, err := client.DescribeSubnet(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe subnet failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.Subnet)
}

func TestVpcDescribeVpcPeerings(t *testing.T) {
	client := initVpcClient()
	req := NewDescribeVpcPeeringsRequestWithoutParam()
	req.SetRegionId("cn-north-1")
	req.SetPageNumber(1)
	req.SetPageSize(100)
	resp, err := client.DescribeVpcPeerings(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe vpcpeerings failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.VpcPeerings)
	fmt.Println(resp.Result.TotalCount)
	fmt.Println(len(resp.Result.VpcPeerings))
}

func TestVpcCreateVpcPeering(t *testing.T) {
	client := initVpcClient()
	req := NewCreateVpcPeeringRequestWithAllParams("cn-north-1", VPCPEERING_NAME, PEERING_LOCAL_VPC_ID, PEERING_REMOTE_VPC_ID, nil)
	resp, err := client.CreateVpcPeering(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("create vpcpeering failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.VpcPeering)
	vpcPeeringId = resp.Result.VpcPeering.VpcPeeringId
}

func TestVpcModifyVpcPeering(t *testing.T) {
	client := initVpcClient()
	req := NewModifyVpcPeeringRequestWithoutParam()
	req.SetRegionId("cn-north-1")
	req.SetVpcPeeringId(vpcPeeringId)
	req.SetVpcPeeringName(VPCPEERING_NAME2)
	resp, err := client.ModifyVpcPeering(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("modify vpcpeering failed. ", err, resp.Error.Code, resp.Error.Message, resp.RequestID)
	} else {
		fmt.Println("ModifyVpcPeering success")
	}
}

func TestVpcDescribeVpcPeering(t *testing.T) {
	client := initVpcClient()
	req := NewDescribeVpcPeeringRequestWithAllParams("cn-north-1", vpcPeeringId)
	resp, err := client.DescribeVpcPeering(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe vpcpeering failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.VpcPeering)
}

func TestVpcDeleteVpcPeering(t *testing.T) {
	client := initVpcClient()
	req := NewDeleteVpcPeeringRequestWithAllParams("cn-north-1", vpcPeeringId)
	resp, err := client.DeleteVpcPeering(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("delete vpcpeering failed. ", err, resp.Error.Code, resp.Error.Message)
	} else {
		fmt.Println("DeleteVpcPeering success")
	}
}

func TestVpcDescribeVpcs(t *testing.T) {
	client := initVpcClient()
	req := NewDescribeVpcsRequestWithoutParam()
	req.SetRegionId("cn-north-1")
	req.SetPageNumber(1)
	req.SetPageSize(100)
	resp, err := client.DescribeVpcs(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe vpcs failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.Vpcs)
	fmt.Println(resp.Result.TotalCount)
	fmt.Println(len(resp.Result.Vpcs))
}

func TestVpcDescribeVpc(t *testing.T) {
	client := initVpcClient()
	req := NewDescribeVpcRequestWithAllParams("cn-north-1", VPC_ID)
	resp, err := client.DescribeVpc(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe vpc failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.Vpc)
}

func TestVpcCreateSecurityGroup(t *testing.T) {
	client := initVpcClient()
	req := NewCreateNetworkSecurityGroupRequest("cn-north-1", VPC_ID, SECURITYGROUP_NAME)
	desc := "sg11111111"
	req.Description = &desc
	resp, err := client.CreateNetworkSecurityGroup(req)
	if err != nil {
		t.Error(err.Error())
		fmt.Println("error")
	}

	if resp.Error.Code != 0 {
		t.Error(fmt.Sprintf("create security group failed, code=%d, message=%s", resp.Error.Code, resp.Error.Message))
	}
	securityGroupId = resp.Result.NetworkSecurityGroupId
	fmt.Println("create security group  success, sgId=%s", securityGroupId)
}

func TestVpcAddSecurityGroupRules(t *testing.T) {
	client := initVpcClient()
	var rule1, rule2 vpc.AddSecurityGroupRules
	rule1.Protocol = SG_RULE_PROTOCOL_TCP
	rule1.Direction = SG_RULE_DIRECTION_INGRESS
	rule1.AddressPrefix = "192.168.3.0/24"
	fport := 1
	tport := 65535
	rule1.FromPort = &fport
	rule1.ToPort = &tport
	r1des := "rule1111111"
	rule1.Description = &r1des

	rule2.Protocol = SG_RULE_PROTOCOL_ICMP
	rule2.Direction = SG_RULE_DIRECTION_EGRESS
	rule2.AddressPrefix = "0.0.0.0/24"
	r2des := "rule2222222"
	rule2.Description = &r2des

	var ruleSpecs []vpc.AddSecurityGroupRules
	ruleSpecs = append(ruleSpecs, rule1)
	ruleSpecs = append(ruleSpecs, rule2)

	req := NewAddNetworkSecurityGroupRulesRequest("cn-north-1", securityGroupId, ruleSpecs)
	resp, err := client.AddNetworkSecurityGroupRules(req)
	if err != nil {
		t.Error(err.Error())
		fmt.Println("error")
	}

	if resp.Error.Code != 0 {
		t.Error(fmt.Sprintf("add security group rules failed, code=%d, message=%s", resp.Error.Code, resp.Error.Message))
	}

	req2 := NewDescribeNetworkSecurityGroupRequest("cn-north-1", securityGroupId)
	resp2, err := client.DescribeNetworkSecurityGroup(req2)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe network securitygroup failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	for _, rule := range resp2.Result.NetworkSecurityGroup.SecurityGroupRules {
		if rule.Protocol == SG_RULE_PROTOCOL_TCP {
			sgRuleId1 = rule.RuleId
		}
		if rule.Protocol == SG_RULE_PROTOCOL_ICMP {
			sgRuleId2 = rule.RuleId
		}
	}
	fmt.Println("add security group rules success, rule1=%s, rule2=%s", sgRuleId1, sgRuleId2)
}

func TestVpcDescribeNetworkSecurityGroups(t *testing.T) {
	client := initVpcClient()
	req := NewDescribeNetworkSecurityGroupsRequestWithoutParam()
	req.SetRegionId("cn-north-1")
	req.SetPageNumber(1)
	req.SetPageSize(100)
	resp, err := client.DescribeNetworkSecurityGroups(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe network securitygroups failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.NetworkSecurityGroups)
	fmt.Println(resp.Result.TotalCount)
	fmt.Println(len(resp.Result.NetworkSecurityGroups))
}

func TestVpcDescribeNetworkSecurityGroup(t *testing.T) {
	client := initVpcClient()
	req := NewDescribeNetworkSecurityGroupRequestWithoutParam()
	req.SetRegionId("cn-north-1")
	req.SetNetworkSecurityGroupId(securityGroupId)
	resp, err := client.DescribeNetworkSecurityGroup(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe network securitygroup failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.NetworkSecurityGroup)
}

func TestVpcModifySecurityGroupRules(t *testing.T) {
	client := initVpcClient()

	var rule1, rule2 vpc.ModifySecurityGroupRules
	rule1.RuleId = sgRuleId1
	r1proto := SG_RULE_PROTOCOL_ICMP
	rule1.Protocol = &r1proto
	r1prefix := "0.0.0.0/24"
	rule1.AddressPrefix = &r1prefix
	r1des := "rule1111111"
	rule1.Description = &r1des

	rule2.RuleId = sgRuleId2
	r2proto := SG_RULE_PROTOCOL_TCP
	rule2.Protocol = &r2proto
	r2prefix := "192.168.3.0/24"
	rule2.AddressPrefix = &r2prefix
	fport := 1
	tport := 65535
	rule2.FromPort = &fport
	rule2.ToPort = &tport
	r2des := "rule222222"
	rule2.Description = &r2des

	var ruleSpecs []vpc.ModifySecurityGroupRules
	ruleSpecs = append(ruleSpecs, rule1)
	ruleSpecs = append(ruleSpecs, rule2)

	req := NewModifyNetworkSecurityGroupRulesRequest("cn-north-1", securityGroupId, ruleSpecs)
	resp, err := client.ModifyNetworkSecurityGroupRules(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("modify security group rules failed ", err, resp.Error.Code, resp.Error.Message, resp.RequestID)
	} else {
		fmt.Println("ModifySecurityGroupRules success")
	}
}

func TestVpcRemoveSecurityGroupRules(t *testing.T) {
	client := initVpcClient()
	var ruleIds []string
	ruleIds = append(ruleIds, sgRuleId1)
	ruleIds = append(ruleIds, sgRuleId2)
	req := NewRemoveNetworkSecurityGroupRulesRequest("cn-north-1", securityGroupId, ruleIds)

	resp, err := client.RemoveNetworkSecurityGroupRules(req)
	if err != nil {
		t.Error(err.Error())
		fmt.Println("error")
	}

	if resp.Error.Code != 0 {
		t.Error(fmt.Sprintf("remove security group rules failed, code=%d, message=%s", resp.Error.Code, resp.Error.Message))
	}
	fmt.Println("remove security group rules success")
}

func TestVpcModifySecurityGroup(t *testing.T) {
	client := initVpcClient()
	req := NewModifyNetworkSecurityGroupRequest("cn-north-1", securityGroupId)
	name := "test_sg1_update"
	req.NetworkSecurityGroupName = &name
	desc := "sgsgsgsg"
	req.Description = &desc
	resp, err := client.ModifyNetworkSecurityGroup(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("modify security group failed ", err, resp.Error.Code, resp.Error.Message, resp.RequestID)
	} else {
		fmt.Println("ModifySecurityGroup success")
	}
}

func TestVpcDeleteSecurityGroup(t *testing.T) {
	client := initVpcClient()
	req := NewDeleteNetworkSecurityGroupRequest("cn-north-1", securityGroupId)
	resp, err := client.DeleteNetworkSecurityGroup(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("delete security group failed ", err, resp.Error.Code, resp.Error.Message, resp.RequestID)
	} else {
		fmt.Println("DeleteNetworkSecurityGroup success")
	}
}

func TestVpcDescribeQuota(t *testing.T) {
	client := initVpcClient()
	req := NewDescribeQuotaRequestWithoutParam()
	req.SetRegionId("cn-north-1")
	req.SetType("subnet")
	req.SetParentResourceId("wp-20180712-01")
	resp, err := client.DescribeQuota(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("describe quota failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.Result.Quota)
}

func TestVpcModifyQuota(t *testing.T) {
	client := initVpcClient()
	req := NewModifyQuotaRequestWithoutParam()
	req.SetRegionId("cn-north-1")
	req.SetType("subnet")
	req.SetParentResourceId(VPC_ID)
	req.AddHeader("x-jdcloud-erp", "wangpan74")
	req.SetMaxLimit(152)
	resp, err := client.ModifyQuota(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("modify quota failed. ", err, resp.Error.Code, resp.Error.Message)
	} else {
		fmt.Println("ModifyQuota success")
	}
}

func TestCreateVpc(t *testing.T) {
	client := initVpcClient()
	req := NewCreateVpcRequest("cn-north-1", VPC_NAME)
	req.SetAddressPrefix("172.16.0.0/18")
	resp, err := client.CreateVpc(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("create vpc failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	vpcId = resp.Result.VpcId
	fmt.Println(resp.Result.VpcId)
	fmt.Println(resp.RequestID)
}

func TestDeleteVpc(t *testing.T) {
	client := initVpcClient()
	req := NewDeleteVpcRequest("cn-north-1", vpcId)
	resp, err := client.DeleteVpc(req)
	if err != nil || resp.Error.Code != 0 {
		t.Error("create vpc failed. ", err, resp.Error.Code, resp.Error.Message)
	}
	fmt.Println(resp.RequestID)
}
