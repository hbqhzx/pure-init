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
    renewal "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/renewal/models"
    common "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/common/models"
)

type ResourceRelationRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /*   */
    ServiceCode string `json:"serviceCode"`

    /*   */
    AppCode string `json:"appCode"`

    /* resourceIdList
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId:  (Required)
 * param serviceCode:  (Required)
 * param appCode:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewResourceRelationRequest(
    regionId string,
    serviceCode string,
    appCode string,
) *ResourceRelationRequest {

	return &ResourceRelationRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/resourceRelation",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ServiceCode: serviceCode,
        AppCode: appCode,
	}
}

/*
 * param regionId:  (Required)
 * param serviceCode:  (Required)
 * param appCode:  (Required)
 * param filters: resourceIdList
 (Optional)
 */
func NewResourceRelationRequestWithAllParams(
    regionId string,
    serviceCode string,
    appCode string,
    filters []common.Filter,
) *ResourceRelationRequest {

    return &ResourceRelationRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/resourceRelation",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ServiceCode: serviceCode,
        AppCode: appCode,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewResourceRelationRequestWithoutParam() *ResourceRelationRequest {

    return &ResourceRelationRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/resourceRelation",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *ResourceRelationRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param serviceCode: (Required) */
func (r *ResourceRelationRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = serviceCode
}

/* param appCode: (Required) */
func (r *ResourceRelationRequest) SetAppCode(appCode string) {
    r.AppCode = appCode
}

/* param filters: resourceIdList
(Optional) */
func (r *ResourceRelationRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ResourceRelationRequest) GetRegionId() string {
    return r.RegionId
}

type ResourceRelationResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ResourceRelationResult `json:"result"`
}

type ResourceRelationResult struct {
    ListQueries []renewal.ListQuery `json:"listQueries"`
    TotalCount int `json:"totalCount"`
}