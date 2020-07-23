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

type DisableCCObserverModeRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 实例 ID  */
    InstanceId int `json:"instanceId"`
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: 实例 ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDisableCCObserverModeRequest(
    regionId string,
    instanceId int,
) *DisableCCObserverModeRequest {

	return &DisableCCObserverModeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:disableCCObserverMode",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: 实例 ID (Required)
 */
func NewDisableCCObserverModeRequestWithAllParams(
    regionId string,
    instanceId int,
) *DisableCCObserverModeRequest {

    return &DisableCCObserverModeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:disableCCObserverMode",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDisableCCObserverModeRequestWithoutParam() *DisableCCObserverModeRequest {

    return &DisableCCObserverModeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:disableCCObserverMode",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DisableCCObserverModeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例 ID(Required) */
func (r *DisableCCObserverModeRequest) SetInstanceId(instanceId int) {
    r.InstanceId = instanceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DisableCCObserverModeRequest) GetRegionId() string {
    return r.RegionId
}

type DisableCCObserverModeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DisableCCObserverModeResult `json:"result"`
}

type DisableCCObserverModeResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}