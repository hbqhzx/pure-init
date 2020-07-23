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
    clouddnsservice "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/clouddnsservice/models"
)

type GetMonitorRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 域名ID，请使用getDomains接口获取。  */
    DomainId string `json:"domainId"`

    /* 当前页数，起始值为1，默认为1 (Optional) */
    PageIndex *int `json:"pageIndex"`

    /* 分页查询时设置的每页行数 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 查询的值 (Optional) */
    SearchValue *string `json:"searchValue"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用getDomains接口获取。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetMonitorRequest(
    regionId string,
    domainId string,
) *GetMonitorRequest {

	return &GetMonitorRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/domain/{domainId}/monitor",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DomainId: domainId,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用getDomains接口获取。 (Required)
 * param pageIndex: 当前页数，起始值为1，默认为1 (Optional)
 * param pageSize: 分页查询时设置的每页行数 (Optional)
 * param searchValue: 查询的值 (Optional)
 */
func NewGetMonitorRequestWithAllParams(
    regionId string,
    domainId string,
    pageIndex *int,
    pageSize *int,
    searchValue *string,
) *GetMonitorRequest {

    return &GetMonitorRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/monitor",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DomainId: domainId,
        PageIndex: pageIndex,
        PageSize: pageSize,
        SearchValue: searchValue,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetMonitorRequestWithoutParam() *GetMonitorRequest {

    return &GetMonitorRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/monitor",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *GetMonitorRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param domainId: 域名ID，请使用getDomains接口获取。(Required) */
func (r *GetMonitorRequest) SetDomainId(domainId string) {
    r.DomainId = domainId
}

/* param pageIndex: 当前页数，起始值为1，默认为1(Optional) */
func (r *GetMonitorRequest) SetPageIndex(pageIndex int) {
    r.PageIndex = &pageIndex
}

/* param pageSize: 分页查询时设置的每页行数(Optional) */
func (r *GetMonitorRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param searchValue: 查询的值(Optional) */
func (r *GetMonitorRequest) SetSearchValue(searchValue string) {
    r.SearchValue = &searchValue
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetMonitorRequest) GetRegionId() string {
    return r.RegionId
}

type GetMonitorResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetMonitorResult `json:"result"`
}

type GetMonitorResult struct {
    CurrentCount int `json:"currentCount"`
    TotalCount int `json:"totalCount"`
    TotalPage int `json:"totalPage"`
    DataList []clouddnsservice.Monitor `json:"dataList"`
}