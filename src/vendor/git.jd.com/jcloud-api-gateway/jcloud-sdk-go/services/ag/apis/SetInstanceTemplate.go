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

type SetInstanceTemplateRequest struct {

    core.JDCloudRequest

    /* 地域  */
    RegionId string `json:"regionId"`

    /* 可用组 ID  */
    AgId string `json:"agId"`

    /* 实力模板 id (Optional) */
    InstanceTemplateId *string `json:"instanceTemplateId"`
}

/*
 * param regionId: 地域 (Required)
 * param agId: 可用组 ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetInstanceTemplateRequest(
    regionId string,
    agId string,
) *SetInstanceTemplateRequest {

	return &SetInstanceTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/availabilityGroups/{agId}:setInstanceTemplate",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AgId: agId,
	}
}

/*
 * param regionId: 地域 (Required)
 * param agId: 可用组 ID (Required)
 * param instanceTemplateId: 实力模板 id (Optional)
 */
func NewSetInstanceTemplateRequestWithAllParams(
    regionId string,
    agId string,
    instanceTemplateId *string,
) *SetInstanceTemplateRequest {

    return &SetInstanceTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/availabilityGroups/{agId}:setInstanceTemplate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AgId: agId,
        InstanceTemplateId: instanceTemplateId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetInstanceTemplateRequestWithoutParam() *SetInstanceTemplateRequest {

    return &SetInstanceTemplateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/availabilityGroups/{agId}:setInstanceTemplate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域(Required) */
func (r *SetInstanceTemplateRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param agId: 可用组 ID(Required) */
func (r *SetInstanceTemplateRequest) SetAgId(agId string) {
    r.AgId = agId
}

/* param instanceTemplateId: 实力模板 id(Optional) */
func (r *SetInstanceTemplateRequest) SetInstanceTemplateId(instanceTemplateId string) {
    r.InstanceTemplateId = &instanceTemplateId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetInstanceTemplateRequest) GetRegionId() string {
    return r.RegionId
}

type SetInstanceTemplateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetInstanceTemplateResult `json:"result"`
}

type SetInstanceTemplateResult struct {
}