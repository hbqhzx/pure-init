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

type AssociateSecurityGroupRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* LB ID  */
    LoadBalancerId string `json:"loadBalancerId"`

    /* 安全组 ID列表  */
    SecurityGroupIds []string `json:"securityGroupIds"`
}

/*
 * param regionId: Region ID (Required)
 * param loadBalancerId: LB ID (Required)
 * param securityGroupIds: 安全组 ID列表 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAssociateSecurityGroupRequest(
    regionId string,
    loadBalancerId string,
    securityGroupIds []string,
) *AssociateSecurityGroupRequest {

	return &AssociateSecurityGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/loadBalancers/{loadBalancerId}:associateSecurityGroup",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        LoadBalancerId: loadBalancerId,
        SecurityGroupIds: securityGroupIds,
	}
}

/*
 * param regionId: Region ID (Required)
 * param loadBalancerId: LB ID (Required)
 * param securityGroupIds: 安全组 ID列表 (Required)
 */
func NewAssociateSecurityGroupRequestWithAllParams(
    regionId string,
    loadBalancerId string,
    securityGroupIds []string,
) *AssociateSecurityGroupRequest {

    return &AssociateSecurityGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/loadBalancers/{loadBalancerId}:associateSecurityGroup",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        LoadBalancerId: loadBalancerId,
        SecurityGroupIds: securityGroupIds,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAssociateSecurityGroupRequestWithoutParam() *AssociateSecurityGroupRequest {

    return &AssociateSecurityGroupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/loadBalancers/{loadBalancerId}:associateSecurityGroup",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *AssociateSecurityGroupRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param loadBalancerId: LB ID(Required) */
func (r *AssociateSecurityGroupRequest) SetLoadBalancerId(loadBalancerId string) {
    r.LoadBalancerId = loadBalancerId
}

/* param securityGroupIds: 安全组 ID列表(Required) */
func (r *AssociateSecurityGroupRequest) SetSecurityGroupIds(securityGroupIds []string) {
    r.SecurityGroupIds = securityGroupIds
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AssociateSecurityGroupRequest) GetRegionId() string {
    return r.RegionId
}

type AssociateSecurityGroupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AssociateSecurityGroupResult `json:"result"`
}

type AssociateSecurityGroupResult struct {
}