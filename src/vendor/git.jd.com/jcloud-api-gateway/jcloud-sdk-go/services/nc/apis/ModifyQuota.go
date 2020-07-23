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

type ModifyQuotaRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* resourceType - 资源类型，支持多个[container, secret, pod]
  */
    ResourceType string `json:"resourceType"`

    /* 上限  */
    Limit int `json:"limit"`
}

/*
 * param regionId: Region ID (Required)
 * param resourceType: resourceType - 资源类型，支持多个[container, secret, pod]
 (Required)
 * param limit: 上限 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyQuotaRequest(
    regionId string,
    resourceType string,
    limit int,
) *ModifyQuotaRequest {

	return &ModifyQuotaRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/quotas",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ResourceType: resourceType,
        Limit: limit,
	}
}

/*
 * param regionId: Region ID (Required)
 * param resourceType: resourceType - 资源类型，支持多个[container, secret, pod]
 (Required)
 * param limit: 上限 (Required)
 */
func NewModifyQuotaRequestWithAllParams(
    regionId string,
    resourceType string,
    limit int,
) *ModifyQuotaRequest {

    return &ModifyQuotaRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/quotas",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ResourceType: resourceType,
        Limit: limit,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyQuotaRequestWithoutParam() *ModifyQuotaRequest {

    return &ModifyQuotaRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/quotas",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *ModifyQuotaRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param resourceType: resourceType - 资源类型，支持多个[container, secret, pod]
(Required) */
func (r *ModifyQuotaRequest) SetResourceType(resourceType string) {
    r.ResourceType = resourceType
}

/* param limit: 上限(Required) */
func (r *ModifyQuotaRequest) SetLimit(limit int) {
    r.Limit = limit
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyQuotaRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyQuotaResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyQuotaResult `json:"result"`
}

type ModifyQuotaResult struct {
}