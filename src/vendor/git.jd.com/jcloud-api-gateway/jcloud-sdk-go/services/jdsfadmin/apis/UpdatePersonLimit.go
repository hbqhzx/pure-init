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

type UpdatePersonLimitRequest struct {

    core.JDCloudRequest

    /* 区域id  */
    RegionId string `json:"regionId"`

    /* 主键Id  */
    Id int `json:"id"`

    /* 限额  */
    LimitCount int `json:"limitCount"`

    /* 限额单位（ALL_REGION,SINGLE_REGION）  */
    ProductUnit string `json:"productUnit"`
}

/*
 * param regionId: 区域id (Required)
 * param id: 主键Id (Required)
 * param limitCount: 限额 (Required)
 * param productUnit: 限额单位（ALL_REGION,SINGLE_REGION） (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdatePersonLimitRequest(
    regionId string,
    id int,
    limitCount int,
    productUnit string,
) *UpdatePersonLimitRequest {

	return &UpdatePersonLimitRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/person_resource_limits/{id}",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Id: id,
        LimitCount: limitCount,
        ProductUnit: productUnit,
	}
}

/*
 * param regionId: 区域id (Required)
 * param id: 主键Id (Required)
 * param limitCount: 限额 (Required)
 * param productUnit: 限额单位（ALL_REGION,SINGLE_REGION） (Required)
 */
func NewUpdatePersonLimitRequestWithAllParams(
    regionId string,
    id int,
    limitCount int,
    productUnit string,
) *UpdatePersonLimitRequest {

    return &UpdatePersonLimitRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/person_resource_limits/{id}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Id: id,
        LimitCount: limitCount,
        ProductUnit: productUnit,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdatePersonLimitRequestWithoutParam() *UpdatePersonLimitRequest {

    return &UpdatePersonLimitRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/person_resource_limits/{id}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域id(Required) */
func (r *UpdatePersonLimitRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param id: 主键Id(Required) */
func (r *UpdatePersonLimitRequest) SetId(id int) {
    r.Id = id
}

/* param limitCount: 限额(Required) */
func (r *UpdatePersonLimitRequest) SetLimitCount(limitCount int) {
    r.LimitCount = limitCount
}

/* param productUnit: 限额单位（ALL_REGION,SINGLE_REGION）(Required) */
func (r *UpdatePersonLimitRequest) SetProductUnit(productUnit string) {
    r.ProductUnit = productUnit
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdatePersonLimitRequest) GetRegionId() string {
    return r.RegionId
}

type UpdatePersonLimitResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdatePersonLimitResult `json:"result"`
}

type UpdatePersonLimitResult struct {
    UpdateResult string `json:"updateResult"`
}