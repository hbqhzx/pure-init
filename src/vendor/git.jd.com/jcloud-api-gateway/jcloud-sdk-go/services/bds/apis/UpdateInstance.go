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
    bds "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/bds/models"
)

type UpdateInstanceRequest struct {

    core.JDCloudRequest

    /* 区域代码  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`

    /* 更新用户  */
    InstanceSpec *bds.InstanceAttribute `json:"instanceSpec"`
}

/*
 * param regionId: 区域代码 (Required)
 * param instanceId: 实例ID (Required)
 * param instanceSpec: 更新用户 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateInstanceRequest(
    regionId string,
    instanceId string,
    instanceSpec *bds.InstanceAttribute,
) *UpdateInstanceRequest {

	return &UpdateInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        InstanceSpec: instanceSpec,
	}
}

/*
 * param regionId: 区域代码 (Required)
 * param instanceId: 实例ID (Required)
 * param instanceSpec: 更新用户 (Required)
 */
func NewUpdateInstanceRequestWithAllParams(
    regionId string,
    instanceId string,
    instanceSpec *bds.InstanceAttribute,
) *UpdateInstanceRequest {

    return &UpdateInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        InstanceSpec: instanceSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateInstanceRequestWithoutParam() *UpdateInstanceRequest {

    return &UpdateInstanceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域代码(Required) */
func (r *UpdateInstanceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例ID(Required) */
func (r *UpdateInstanceRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param instanceSpec: 更新用户(Required) */
func (r *UpdateInstanceRequest) SetInstanceSpec(instanceSpec *bds.InstanceAttribute) {
    r.InstanceSpec = instanceSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateInstanceRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateInstanceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateInstanceResult `json:"result"`
}

type UpdateInstanceResult struct {
}