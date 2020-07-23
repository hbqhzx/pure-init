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

type UpdateWafUpdatesRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* waf 实例id  */
    InstanceId string `json:"instanceId"`
}

/*
 * param regionId: 地域 Id (Required)
 * param instanceId: waf 实例id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateWafUpdatesRequest(
    regionId string,
    instanceId string,
) *UpdateWafUpdatesRequest {

	return &UpdateWafUpdatesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/wafUpdates",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param instanceId: waf 实例id (Required)
 */
func NewUpdateWafUpdatesRequestWithAllParams(
    regionId string,
    instanceId string,
) *UpdateWafUpdatesRequest {

    return &UpdateWafUpdatesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/wafUpdates",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateWafUpdatesRequestWithoutParam() *UpdateWafUpdatesRequest {

    return &UpdateWafUpdatesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/wafUpdates",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *UpdateWafUpdatesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: waf 实例id(Required) */
func (r *UpdateWafUpdatesRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateWafUpdatesRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateWafUpdatesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateWafUpdatesResult `json:"result"`
}

type UpdateWafUpdatesResult struct {
}