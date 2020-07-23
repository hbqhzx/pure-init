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

type GetResourceQuotaByObjectRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /*   */
    QueryVo *quota.ResourceQuotaReqVo `json:"queryVo"`
}

/*
 * param regionId: Region ID (Required)
 * param queryVo:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetResourceQuotaByObjectRequest(
    regionId string,
    queryVo *quota.ResourceQuotaReqVo,
) *GetResourceQuotaByObjectRequest {

	return &GetResourceQuotaByObjectRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/resourceQuota",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        QueryVo: queryVo,
	}
}

/*
 * param regionId: Region ID (Required)
 * param queryVo:  (Required)
 */
func NewGetResourceQuotaByObjectRequestWithAllParams(
    regionId string,
    queryVo *quota.ResourceQuotaReqVo,
) *GetResourceQuotaByObjectRequest {

    return &GetResourceQuotaByObjectRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/resourceQuota",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        QueryVo: queryVo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetResourceQuotaByObjectRequestWithoutParam() *GetResourceQuotaByObjectRequest {

    return &GetResourceQuotaByObjectRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/resourceQuota",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *GetResourceQuotaByObjectRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param queryVo: (Required) */
func (r *GetResourceQuotaByObjectRequest) SetQueryVo(queryVo *quota.ResourceQuotaReqVo) {
    r.QueryVo = queryVo
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetResourceQuotaByObjectRequest) GetRegionId() string {
    return r.RegionId
}

type GetResourceQuotaByObjectResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetResourceQuotaByObjectResult `json:"result"`
}

type GetResourceQuotaByObjectResult struct {
    Data quota.ResourceQuotaVo `json:"data"`
}