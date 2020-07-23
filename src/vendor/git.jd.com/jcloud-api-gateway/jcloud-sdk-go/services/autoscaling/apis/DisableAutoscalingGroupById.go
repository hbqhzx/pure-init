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

type DisableAutoscalingGroupByIdRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 组uuid  */
    AsGroupUuid string `json:"asGroupUuid"`
}

/*
 * param regionId: 地域 Id (Required)
 * param asGroupUuid: 组uuid (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDisableAutoscalingGroupByIdRequest(
    regionId string,
    asGroupUuid string,
) *DisableAutoscalingGroupByIdRequest {

	return &DisableAutoscalingGroupByIdRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/autoscalingGroups/{asGroupUuid}:disable",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AsGroupUuid: asGroupUuid,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param asGroupUuid: 组uuid (Required)
 */
func NewDisableAutoscalingGroupByIdRequestWithAllParams(
    regionId string,
    asGroupUuid string,
) *DisableAutoscalingGroupByIdRequest {

    return &DisableAutoscalingGroupByIdRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/autoscalingGroups/{asGroupUuid}:disable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AsGroupUuid: asGroupUuid,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDisableAutoscalingGroupByIdRequestWithoutParam() *DisableAutoscalingGroupByIdRequest {

    return &DisableAutoscalingGroupByIdRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/autoscalingGroups/{asGroupUuid}:disable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DisableAutoscalingGroupByIdRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param asGroupUuid: 组uuid(Required) */
func (r *DisableAutoscalingGroupByIdRequest) SetAsGroupUuid(asGroupUuid string) {
    r.AsGroupUuid = asGroupUuid
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DisableAutoscalingGroupByIdRequest) GetRegionId() string {
    return r.RegionId
}

type DisableAutoscalingGroupByIdResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DisableAutoscalingGroupByIdResult `json:"result"`
}

type DisableAutoscalingGroupByIdResult struct {
}