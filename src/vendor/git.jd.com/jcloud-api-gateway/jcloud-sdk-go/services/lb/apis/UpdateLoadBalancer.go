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

type UpdateLoadBalancerRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* LB ID  */
    LoadBalancerId string `json:"loadBalancerId"`

    /* LoadBalancer的名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Optional) */
    LoadBalancerName *string `json:"loadBalancerName"`

    /* Start：开启，Stop：关闭 (Optional) */
    Action *string `json:"action"`

    /* LoadBalancer的描述信息,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional) */
    Description *string `json:"description"`

    /* 删除保护，取值为True(开启)或False(关闭)，默认为False (Optional) */
    DeleteProtection *bool `json:"deleteProtection"`
}

/*
 * param regionId: Region ID (Required)
 * param loadBalancerId: LB ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateLoadBalancerRequest(
    regionId string,
    loadBalancerId string,
) *UpdateLoadBalancerRequest {

	return &UpdateLoadBalancerRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/loadBalancers/{loadBalancerId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        LoadBalancerId: loadBalancerId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param loadBalancerId: LB ID (Required)
 * param loadBalancerName: LoadBalancer的名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Optional)
 * param action: Start：开启，Stop：关闭 (Optional)
 * param description: LoadBalancer的描述信息,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional)
 * param deleteProtection: 删除保护，取值为True(开启)或False(关闭)，默认为False (Optional)
 */
func NewUpdateLoadBalancerRequestWithAllParams(
    regionId string,
    loadBalancerId string,
    loadBalancerName *string,
    action *string,
    description *string,
    deleteProtection *bool,
) *UpdateLoadBalancerRequest {

    return &UpdateLoadBalancerRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/loadBalancers/{loadBalancerId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        LoadBalancerId: loadBalancerId,
        LoadBalancerName: loadBalancerName,
        Action: action,
        Description: description,
        DeleteProtection: deleteProtection,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateLoadBalancerRequestWithoutParam() *UpdateLoadBalancerRequest {

    return &UpdateLoadBalancerRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/loadBalancers/{loadBalancerId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *UpdateLoadBalancerRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param loadBalancerId: LB ID(Required) */
func (r *UpdateLoadBalancerRequest) SetLoadBalancerId(loadBalancerId string) {
    r.LoadBalancerId = loadBalancerId
}

/* param loadBalancerName: LoadBalancer的名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符(Optional) */
func (r *UpdateLoadBalancerRequest) SetLoadBalancerName(loadBalancerName string) {
    r.LoadBalancerName = &loadBalancerName
}

/* param action: Start：开启，Stop：关闭(Optional) */
func (r *UpdateLoadBalancerRequest) SetAction(action string) {
    r.Action = &action
}

/* param description: LoadBalancer的描述信息,允许输入UTF-8编码下的全部字符，不超过256字符(Optional) */
func (r *UpdateLoadBalancerRequest) SetDescription(description string) {
    r.Description = &description
}

/* param deleteProtection: 删除保护，取值为True(开启)或False(关闭)，默认为False(Optional) */
func (r *UpdateLoadBalancerRequest) SetDeleteProtection(deleteProtection bool) {
    r.DeleteProtection = &deleteProtection
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateLoadBalancerRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateLoadBalancerResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateLoadBalancerResult `json:"result"`
}

type UpdateLoadBalancerResult struct {
}