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
    vm "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/vm/models"
)

type ModifyInstanceDiskAttributeRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 云主机ID  */
    InstanceId string `json:"instanceId"`

    /* 云硬盘列表 (Optional) */
    DataDisks []vm.InstanceDiskAttribute `json:"dataDisks"`
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: 云主机ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyInstanceDiskAttributeRequest(
    regionId string,
    instanceId string,
) *ModifyInstanceDiskAttributeRequest {

	return &ModifyInstanceDiskAttributeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstanceDiskAttribute",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: 云主机ID (Required)
 * param dataDisks: 云硬盘列表 (Optional)
 */
func NewModifyInstanceDiskAttributeRequestWithAllParams(
    regionId string,
    instanceId string,
    dataDisks []vm.InstanceDiskAttribute,
) *ModifyInstanceDiskAttributeRequest {

    return &ModifyInstanceDiskAttributeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstanceDiskAttribute",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        DataDisks: dataDisks,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyInstanceDiskAttributeRequestWithoutParam() *ModifyInstanceDiskAttributeRequest {

    return &ModifyInstanceDiskAttributeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstanceDiskAttribute",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *ModifyInstanceDiskAttributeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 云主机ID(Required) */
func (r *ModifyInstanceDiskAttributeRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param dataDisks: 云硬盘列表(Optional) */
func (r *ModifyInstanceDiskAttributeRequest) SetDataDisks(dataDisks []vm.InstanceDiskAttribute) {
    r.DataDisks = dataDisks
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyInstanceDiskAttributeRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyInstanceDiskAttributeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyInstanceDiskAttributeResult `json:"result"`
}

type ModifyInstanceDiskAttributeResult struct {
}