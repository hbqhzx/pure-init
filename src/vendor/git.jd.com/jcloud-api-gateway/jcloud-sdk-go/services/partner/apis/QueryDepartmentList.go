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
    partner "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/partner/models"
)

type QueryDepartmentListRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* 部门编号 (Optional) */
    DepId *string `json:"depId"`

    /* 部门名称 (Optional) */
    DepName *string `json:"depName"`

    /* 当前页序号 (Optional) */
    PageIndex *int `json:"pageIndex"`

    /* 当前条数 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryDepartmentListRequest(
    regionId string,
) *QueryDepartmentListRequest {

	return &QueryDepartmentListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/department/queryDepartmentList",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId:  (Required)
 * param depId: 部门编号 (Optional)
 * param depName: 部门名称 (Optional)
 * param pageIndex: 当前页序号 (Optional)
 * param pageSize: 当前条数 (Optional)
 */
func NewQueryDepartmentListRequestWithAllParams(
    regionId string,
    depId *string,
    depName *string,
    pageIndex *int,
    pageSize *int,
) *QueryDepartmentListRequest {

    return &QueryDepartmentListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/department/queryDepartmentList",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DepId: depId,
        DepName: depName,
        PageIndex: pageIndex,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryDepartmentListRequestWithoutParam() *QueryDepartmentListRequest {

    return &QueryDepartmentListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/department/queryDepartmentList",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *QueryDepartmentListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param depId: 部门编号(Optional) */
func (r *QueryDepartmentListRequest) SetDepId(depId string) {
    r.DepId = &depId
}

/* param depName: 部门名称(Optional) */
func (r *QueryDepartmentListRequest) SetDepName(depName string) {
    r.DepName = &depName
}

/* param pageIndex: 当前页序号(Optional) */
func (r *QueryDepartmentListRequest) SetPageIndex(pageIndex int) {
    r.PageIndex = &pageIndex
}

/* param pageSize: 当前条数(Optional) */
func (r *QueryDepartmentListRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryDepartmentListRequest) GetRegionId() string {
    return r.RegionId
}

type QueryDepartmentListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryDepartmentListResult `json:"result"`
}

type QueryDepartmentListResult struct {
    Pagination partner.Pagination `json:"pagination"`
    Result []partner.Department `json:"result"`
}