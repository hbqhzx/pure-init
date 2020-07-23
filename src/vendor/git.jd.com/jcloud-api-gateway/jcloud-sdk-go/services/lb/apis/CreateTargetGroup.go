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
)

type CreateTargetGroupRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 服务器组名字,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符  */
    TargetGroupName string `json:"targetGroupName"`

    /* 负载均衡的id  */
    LoadBalancerId string `json:"loadBalancerId"`

    /* 描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional) */
    Description *string `json:"description"`

    /* 服务器组的类型，取值为instance或ip, 默认为instance。当前只有nlb支持ip类型的服务器组 (Optional) */
    Type *string `json:"type"`
}

/*
 * param regionId: Region ID (Required)
 * param targetGroupName: 服务器组名字,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Required)
 * param loadBalancerId: 负载均衡的id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateTargetGroupRequest(
    regionId string,
    targetGroupName string,
    loadBalancerId string,
) *CreateTargetGroupRequest {

	return &CreateTargetGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/targetGroups/",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        TargetGroupName: targetGroupName,
        LoadBalancerId: loadBalancerId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param targetGroupName: 服务器组名字,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Required)
 * param loadBalancerId: 负载均衡的id (Required)
 * param description: 描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional)
 * param type_: 服务器组的类型，取值为instance或ip, 默认为instance。当前只有nlb支持ip类型的服务器组 (Optional)
 */
func NewCreateTargetGroupRequestWithAllParams(
    regionId string,
    targetGroupName string,
    loadBalancerId string,
    description *string,
    type_ *string,
) *CreateTargetGroupRequest {

    return &CreateTargetGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/targetGroups/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        TargetGroupName: targetGroupName,
        LoadBalancerId: loadBalancerId,
        Description: description,
        Type: type_,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateTargetGroupRequestWithoutParam() *CreateTargetGroupRequest {

    return &CreateTargetGroupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/targetGroups/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateTargetGroupRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param targetGroupName: 服务器组名字,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符(Required) */
func (r *CreateTargetGroupRequest) SetTargetGroupName(targetGroupName string) {
    r.TargetGroupName = targetGroupName
}

/* param loadBalancerId: 负载均衡的id(Required) */
func (r *CreateTargetGroupRequest) SetLoadBalancerId(loadBalancerId string) {
    r.LoadBalancerId = loadBalancerId
}

/* param description: 描述,允许输入UTF-8编码下的全部字符，不超过256字符(Optional) */
func (r *CreateTargetGroupRequest) SetDescription(description string) {
    r.Description = &description
}

/* param type_: 服务器组的类型，取值为instance或ip, 默认为instance。当前只有nlb支持ip类型的服务器组(Optional) */
func (r *CreateTargetGroupRequest) SetType(type_ string) {
    r.Type = &type_
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateTargetGroupRequest) GetRegionId() string {
    return r.RegionId
}

type CreateTargetGroupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateTargetGroupResult `json:"result"`
}

type CreateTargetGroupResult struct {
    TargetGroupId string `json:"targetGroupId"`
}