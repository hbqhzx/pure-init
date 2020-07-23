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

type CheckInstanceAgentStatusRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 部署组名  */
    GroupId string `json:"groupId"`
}

/*
 * param regionId: 地域ID (Required)
 * param groupId: 部署组名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCheckInstanceAgentStatusRequest(
    regionId string,
    groupId string,
) *CheckInstanceAgentStatusRequest {

	return &CheckInstanceAgentStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instance:status",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        GroupId: groupId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param groupId: 部署组名 (Required)
 */
func NewCheckInstanceAgentStatusRequestWithAllParams(
    regionId string,
    groupId string,
) *CheckInstanceAgentStatusRequest {

    return &CheckInstanceAgentStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance:status",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        GroupId: groupId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCheckInstanceAgentStatusRequestWithoutParam() *CheckInstanceAgentStatusRequest {

    return &CheckInstanceAgentStatusRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance:status",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CheckInstanceAgentStatusRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param groupId: 部署组名(Required) */
func (r *CheckInstanceAgentStatusRequest) SetGroupId(groupId string) {
    r.GroupId = groupId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CheckInstanceAgentStatusRequest) GetRegionId() string {
    return r.RegionId
}

type CheckInstanceAgentStatusResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CheckInstanceAgentStatusResult `json:"result"`
}

type CheckInstanceAgentStatusResult struct {
    UnhealthyInstances []string `json:"unhealthyInstances"`
    TotalCount int `json:"totalCount"`
    UnhealthyCount int `json:"unhealthyCount"`
}