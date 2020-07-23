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

type DeleteInstanceRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* JDW 实例ID，唯一标识一个JDW实例  */
    InstanceId string `json:"instanceId"`
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceId: JDW 实例ID，唯一标识一个JDW实例 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteInstanceRequest(
    regionId string,
    instanceId string,
) *DeleteInstanceRequest {

	return &DeleteInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域代码 (Required)
 * param instanceId: JDW 实例ID，唯一标识一个JDW实例 (Required)
 */
func NewDeleteInstanceRequestWithAllParams(
    regionId string,
    instanceId string,
) *DeleteInstanceRequest {

    return &DeleteInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteInstanceRequestWithoutParam() *DeleteInstanceRequest {

    return &DeleteInstanceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *DeleteInstanceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: JDW 实例ID，唯一标识一个JDW实例(Required) */
func (r *DeleteInstanceRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteInstanceRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteInstanceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteInstanceResult `json:"result"`
}

type DeleteInstanceResult struct {
}