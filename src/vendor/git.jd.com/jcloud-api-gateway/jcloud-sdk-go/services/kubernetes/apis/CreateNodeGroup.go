// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/core"
    kubernetes "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/kubernetes/models"
)

type CreateNodeGroupRequest struct {

    core.JDCloudRequest

    /* 地域 ID  */
    RegionId string `json:"regionId"`

    /* 名称（同一用户的 cluster 内部唯一）  */
    Name string `json:"name"`

    /* 描述 (Optional) */
    Description *string `json:"description"`

    /* node group所属的cluster  */
    ClusterId string `json:"clusterId"`

    /* 节点组配置  */
    NodeConfig *kubernetes.NodeConfigSpec `json:"nodeConfig"`

    /* nodeGroup初始化大小  */
    InitialNodeCount int `json:"initialNodeCount"`

    /* k8s运行的vpc  */
    VpcId string `json:"vpcId"`

    /* k8s的node的cidr  */
    NodeCidr string `json:"nodeCidr"`

    /* 是否开启 node group 的自动修复，默认关闭 (Optional) */
    AutoRepair *bool `json:"autoRepair"`
}

/*
 * param regionId: 地域 ID (Required)
 * param name: 名称（同一用户的 cluster 内部唯一） (Required)
 * param clusterId: node group所属的cluster (Required)
 * param nodeConfig: 节点组配置 (Required)
 * param initialNodeCount: nodeGroup初始化大小 (Required)
 * param vpcId: k8s运行的vpc (Required)
 * param nodeCidr: k8s的node的cidr (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateNodeGroupRequest(
    regionId string,
    name string,
    clusterId string,
    nodeConfig *kubernetes.NodeConfigSpec,
    initialNodeCount int,
    vpcId string,
    nodeCidr string,
) *CreateNodeGroupRequest {

	return &CreateNodeGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/nodeGroups",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Name: name,
        ClusterId: clusterId,
        NodeConfig: nodeConfig,
        InitialNodeCount: initialNodeCount,
        VpcId: vpcId,
        NodeCidr: nodeCidr,
	}
}

/*
 * param regionId: 地域 ID (Required)
 * param name: 名称（同一用户的 cluster 内部唯一） (Required)
 * param description: 描述 (Optional)
 * param clusterId: node group所属的cluster (Required)
 * param nodeConfig: 节点组配置 (Required)
 * param initialNodeCount: nodeGroup初始化大小 (Required)
 * param vpcId: k8s运行的vpc (Required)
 * param nodeCidr: k8s的node的cidr (Required)
 * param autoRepair: 是否开启 node group 的自动修复，默认关闭 (Optional)
 */
func NewCreateNodeGroupRequestWithAllParams(
    regionId string,
    name string,
    description *string,
    clusterId string,
    nodeConfig *kubernetes.NodeConfigSpec,
    initialNodeCount int,
    vpcId string,
    nodeCidr string,
    autoRepair *bool,
) *CreateNodeGroupRequest {

    return &CreateNodeGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/nodeGroups",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Name: name,
        Description: description,
        ClusterId: clusterId,
        NodeConfig: nodeConfig,
        InitialNodeCount: initialNodeCount,
        VpcId: vpcId,
        NodeCidr: nodeCidr,
        AutoRepair: autoRepair,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateNodeGroupRequestWithoutParam() *CreateNodeGroupRequest {

    return &CreateNodeGroupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/nodeGroups",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 ID(Required) */
func (r *CreateNodeGroupRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param name: 名称（同一用户的 cluster 内部唯一）(Required) */
func (r *CreateNodeGroupRequest) SetName(name string) {
    r.Name = name
}

/* param description: 描述(Optional) */
func (r *CreateNodeGroupRequest) SetDescription(description string) {
    r.Description = &description
}

/* param clusterId: node group所属的cluster(Required) */
func (r *CreateNodeGroupRequest) SetClusterId(clusterId string) {
    r.ClusterId = clusterId
}

/* param nodeConfig: 节点组配置(Required) */
func (r *CreateNodeGroupRequest) SetNodeConfig(nodeConfig *kubernetes.NodeConfigSpec) {
    r.NodeConfig = nodeConfig
}

/* param initialNodeCount: nodeGroup初始化大小(Required) */
func (r *CreateNodeGroupRequest) SetInitialNodeCount(initialNodeCount int) {
    r.InitialNodeCount = initialNodeCount
}

/* param vpcId: k8s运行的vpc(Required) */
func (r *CreateNodeGroupRequest) SetVpcId(vpcId string) {
    r.VpcId = vpcId
}

/* param nodeCidr: k8s的node的cidr(Required) */
func (r *CreateNodeGroupRequest) SetNodeCidr(nodeCidr string) {
    r.NodeCidr = nodeCidr
}

/* param autoRepair: 是否开启 node group 的自动修复，默认关闭(Optional) */
func (r *CreateNodeGroupRequest) SetAutoRepair(autoRepair bool) {
    r.AutoRepair = &autoRepair
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateNodeGroupRequest) GetRegionId() string {
    return r.RegionId
}

type CreateNodeGroupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateNodeGroupResult `json:"result"`
}

type CreateNodeGroupResult struct {
    NodeGroupId string `json:"nodeGroupId"`
}