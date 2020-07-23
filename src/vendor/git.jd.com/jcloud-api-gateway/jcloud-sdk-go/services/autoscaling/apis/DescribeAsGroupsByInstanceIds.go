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

type DescribeAsGroupsByInstanceIdsRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 实例ID组，用,号分割  */
    InstanceIds string `json:"instanceIds"`
}

/*
 * param regionId: 地域 Id (Required)
 * param instanceIds: 实例ID组，用,号分割 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAsGroupsByInstanceIdsRequest(
    regionId string,
    instanceIds string,
) *DescribeAsGroupsByInstanceIdsRequest {

	return &DescribeAsGroupsByInstanceIdsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/autoscalingGroups/profile",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceIds: instanceIds,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param instanceIds: 实例ID组，用,号分割 (Required)
 */
func NewDescribeAsGroupsByInstanceIdsRequestWithAllParams(
    regionId string,
    instanceIds string,
) *DescribeAsGroupsByInstanceIdsRequest {

    return &DescribeAsGroupsByInstanceIdsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/autoscalingGroups/profile",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceIds: instanceIds,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAsGroupsByInstanceIdsRequestWithoutParam() *DescribeAsGroupsByInstanceIdsRequest {

    return &DescribeAsGroupsByInstanceIdsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/autoscalingGroups/profile",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DescribeAsGroupsByInstanceIdsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceIds: 实例ID组，用,号分割(Required) */
func (r *DescribeAsGroupsByInstanceIdsRequest) SetInstanceIds(instanceIds string) {
    r.InstanceIds = instanceIds
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAsGroupsByInstanceIdsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeAsGroupsByInstanceIdsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAsGroupsByInstanceIdsResult `json:"result"`
}

type DescribeAsGroupsByInstanceIdsResult struct {
    AsGroups interface{} `json:"asGroups"`
}