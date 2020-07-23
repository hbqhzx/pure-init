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
    settlement "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/settlement/models"
)

type QueryPermissionListRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* erp (Optional) */
    Erp *string `json:"erp"`

    /* 业务系统 (Optional) */
    SystemCode *string `json:"systemCode"`

    /* 页数（默认1） (Optional) */
    PageIndex *int `json:"pageIndex"`

    /* 页条数（默认10） (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryPermissionListRequest(
    regionId string,
) *QueryPermissionListRequest {

	return &QueryPermissionListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/permission:list",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId:  (Required)
 * param erp: erp (Optional)
 * param systemCode: 业务系统 (Optional)
 * param pageIndex: 页数（默认1） (Optional)
 * param pageSize: 页条数（默认10） (Optional)
 */
func NewQueryPermissionListRequestWithAllParams(
    regionId string,
    erp *string,
    systemCode *string,
    pageIndex *int,
    pageSize *int,
) *QueryPermissionListRequest {

    return &QueryPermissionListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/permission:list",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Erp: erp,
        SystemCode: systemCode,
        PageIndex: pageIndex,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryPermissionListRequestWithoutParam() *QueryPermissionListRequest {

    return &QueryPermissionListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/permission:list",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *QueryPermissionListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param erp: erp(Optional) */
func (r *QueryPermissionListRequest) SetErp(erp string) {
    r.Erp = &erp
}

/* param systemCode: 业务系统(Optional) */
func (r *QueryPermissionListRequest) SetSystemCode(systemCode string) {
    r.SystemCode = &systemCode
}

/* param pageIndex: 页数（默认1）(Optional) */
func (r *QueryPermissionListRequest) SetPageIndex(pageIndex int) {
    r.PageIndex = &pageIndex
}

/* param pageSize: 页条数（默认10）(Optional) */
func (r *QueryPermissionListRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryPermissionListRequest) GetRegionId() string {
    return r.RegionId
}

type QueryPermissionListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryPermissionListResult `json:"result"`
}

type QueryPermissionListResult struct {
    Pagination settlement.Pagination `json:"pagination"`
    Result []settlement.PermissionVo `json:"result"`
}