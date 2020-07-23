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
    quota "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/quota/models"
)

type GetResourceQuotaDetailRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* id  */
    Id string `json:"id"`
}

/*
 * param regionId: Region ID (Required)
 * param id: id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetResourceQuotaDetailRequest(
    regionId string,
    id string,
) *GetResourceQuotaDetailRequest {

	return &GetResourceQuotaDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/resourceQuota/{id}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Id: id,
	}
}

/*
 * param regionId: Region ID (Required)
 * param id: id (Required)
 */
func NewGetResourceQuotaDetailRequestWithAllParams(
    regionId string,
    id string,
) *GetResourceQuotaDetailRequest {

    return &GetResourceQuotaDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/resourceQuota/{id}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Id: id,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetResourceQuotaDetailRequestWithoutParam() *GetResourceQuotaDetailRequest {

    return &GetResourceQuotaDetailRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/resourceQuota/{id}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *GetResourceQuotaDetailRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param id: id(Required) */
func (r *GetResourceQuotaDetailRequest) SetId(id string) {
    r.Id = id
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetResourceQuotaDetailRequest) GetRegionId() string {
    return r.RegionId
}

type GetResourceQuotaDetailResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetResourceQuotaDetailResult `json:"result"`
}

type GetResourceQuotaDetailResult struct {
    Data quota.ResourceQuotaVo `json:"data"`
}