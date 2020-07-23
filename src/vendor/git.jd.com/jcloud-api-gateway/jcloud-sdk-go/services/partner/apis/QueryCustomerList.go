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

type QueryCustomerListRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* 客户pin（客户账户） (Optional) */
    CustomerPin *string `json:"customerPin"`

    /* 客户昵称 (Optional) */
    AliasName *string `json:"aliasName"`

    /* 渠道商级次(0一级、1 二级) (Optional) */
    DistributorLevel *int `json:"distributorLevel"`

    /* 渠道商名称 (Optional) */
    DistributorName *string `json:"distributorName"`

    /* 一级渠道商名称 (Optional) */
    ParentDistributorName *string `json:"parentDistributorName"`

    /* 关联开始时间（格式：yyyy-MM-dd HH:mm:ss） (Optional) */
    StartRelTime *string `json:"startRelTime"`

    /* 关联结束时间（格式：yyyy-MM-dd HH:mm:ss） (Optional) */
    EndRelTime *string `json:"endRelTime"`

    /* 部门 (Optional) */
    Dept *int `json:"dept"`

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
func NewQueryCustomerListRequest(
    regionId string,
) *QueryCustomerListRequest {

	return &QueryCustomerListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/customerManage/queryCustomerList",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId:  (Required)
 * param customerPin: 客户pin（客户账户） (Optional)
 * param aliasName: 客户昵称 (Optional)
 * param distributorLevel: 渠道商级次(0一级、1 二级) (Optional)
 * param distributorName: 渠道商名称 (Optional)
 * param parentDistributorName: 一级渠道商名称 (Optional)
 * param startRelTime: 关联开始时间（格式：yyyy-MM-dd HH:mm:ss） (Optional)
 * param endRelTime: 关联结束时间（格式：yyyy-MM-dd HH:mm:ss） (Optional)
 * param dept: 部门 (Optional)
 * param pageIndex: 当前页序号 (Optional)
 * param pageSize: 当前条数 (Optional)
 */
func NewQueryCustomerListRequestWithAllParams(
    regionId string,
    customerPin *string,
    aliasName *string,
    distributorLevel *int,
    distributorName *string,
    parentDistributorName *string,
    startRelTime *string,
    endRelTime *string,
    dept *int,
    pageIndex *int,
    pageSize *int,
) *QueryCustomerListRequest {

    return &QueryCustomerListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/customerManage/queryCustomerList",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CustomerPin: customerPin,
        AliasName: aliasName,
        DistributorLevel: distributorLevel,
        DistributorName: distributorName,
        ParentDistributorName: parentDistributorName,
        StartRelTime: startRelTime,
        EndRelTime: endRelTime,
        Dept: dept,
        PageIndex: pageIndex,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryCustomerListRequestWithoutParam() *QueryCustomerListRequest {

    return &QueryCustomerListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/customerManage/queryCustomerList",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *QueryCustomerListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param customerPin: 客户pin（客户账户）(Optional) */
func (r *QueryCustomerListRequest) SetCustomerPin(customerPin string) {
    r.CustomerPin = &customerPin
}

/* param aliasName: 客户昵称(Optional) */
func (r *QueryCustomerListRequest) SetAliasName(aliasName string) {
    r.AliasName = &aliasName
}

/* param distributorLevel: 渠道商级次(0一级、1 二级)(Optional) */
func (r *QueryCustomerListRequest) SetDistributorLevel(distributorLevel int) {
    r.DistributorLevel = &distributorLevel
}

/* param distributorName: 渠道商名称(Optional) */
func (r *QueryCustomerListRequest) SetDistributorName(distributorName string) {
    r.DistributorName = &distributorName
}

/* param parentDistributorName: 一级渠道商名称(Optional) */
func (r *QueryCustomerListRequest) SetParentDistributorName(parentDistributorName string) {
    r.ParentDistributorName = &parentDistributorName
}

/* param startRelTime: 关联开始时间（格式：yyyy-MM-dd HH:mm:ss）(Optional) */
func (r *QueryCustomerListRequest) SetStartRelTime(startRelTime string) {
    r.StartRelTime = &startRelTime
}

/* param endRelTime: 关联结束时间（格式：yyyy-MM-dd HH:mm:ss）(Optional) */
func (r *QueryCustomerListRequest) SetEndRelTime(endRelTime string) {
    r.EndRelTime = &endRelTime
}

/* param dept: 部门(Optional) */
func (r *QueryCustomerListRequest) SetDept(dept int) {
    r.Dept = &dept
}

/* param pageIndex: 当前页序号(Optional) */
func (r *QueryCustomerListRequest) SetPageIndex(pageIndex int) {
    r.PageIndex = &pageIndex
}

/* param pageSize: 当前条数(Optional) */
func (r *QueryCustomerListRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryCustomerListRequest) GetRegionId() string {
    return r.RegionId
}

type QueryCustomerListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryCustomerListResult `json:"result"`
}

type QueryCustomerListResult struct {
    Pagination partner.Pagination `json:"pagination"`
    Result []partner.CustomerVo `json:"result"`
}