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

type QueryCustPostFeeListForErpRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* 按月查询开始时间  */
    BeginTime string `json:"beginTime"`

    /* 按月查询结束时间  */
    EndTime string `json:"endTime"`

    /* 客户pin (Optional) */
    CustomerPin *string `json:"customerPin"`

    /* 一级服务商名称 (Optional) */
    OneLevelDistributorName *string `json:"oneLevelDistributorName"`

    /* 二级服务商名称 (Optional) */
    SecondLevelDistributorName *string `json:"secondLevelDistributorName"`

    /* 每页条数  */
    PageSize int `json:"pageSize"`

    /* 第几页  */
    PageIndex int `json:"pageIndex"`
}

/*
 * param regionId:  (Required)
 * param beginTime: 按月查询开始时间 (Required)
 * param endTime: 按月查询结束时间 (Required)
 * param pageSize: 每页条数 (Required)
 * param pageIndex: 第几页 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryCustPostFeeListForErpRequest(
    regionId string,
    beginTime string,
    endTime string,
    pageSize int,
    pageIndex int,
) *QueryCustPostFeeListForErpRequest {

	return &QueryCustPostFeeListForErpRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/achievementManage/queryCustPostFeeListForErp",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        BeginTime: beginTime,
        EndTime: endTime,
        PageSize: pageSize,
        PageIndex: pageIndex,
	}
}

/*
 * param regionId:  (Required)
 * param beginTime: 按月查询开始时间 (Required)
 * param endTime: 按月查询结束时间 (Required)
 * param customerPin: 客户pin (Optional)
 * param oneLevelDistributorName: 一级服务商名称 (Optional)
 * param secondLevelDistributorName: 二级服务商名称 (Optional)
 * param pageSize: 每页条数 (Required)
 * param pageIndex: 第几页 (Required)
 */
func NewQueryCustPostFeeListForErpRequestWithAllParams(
    regionId string,
    beginTime string,
    endTime string,
    customerPin *string,
    oneLevelDistributorName *string,
    secondLevelDistributorName *string,
    pageSize int,
    pageIndex int,
) *QueryCustPostFeeListForErpRequest {

    return &QueryCustPostFeeListForErpRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/achievementManage/queryCustPostFeeListForErp",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        BeginTime: beginTime,
        EndTime: endTime,
        CustomerPin: customerPin,
        OneLevelDistributorName: oneLevelDistributorName,
        SecondLevelDistributorName: secondLevelDistributorName,
        PageSize: pageSize,
        PageIndex: pageIndex,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryCustPostFeeListForErpRequestWithoutParam() *QueryCustPostFeeListForErpRequest {

    return &QueryCustPostFeeListForErpRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/achievementManage/queryCustPostFeeListForErp",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *QueryCustPostFeeListForErpRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param beginTime: 按月查询开始时间(Required) */
func (r *QueryCustPostFeeListForErpRequest) SetBeginTime(beginTime string) {
    r.BeginTime = beginTime
}

/* param endTime: 按月查询结束时间(Required) */
func (r *QueryCustPostFeeListForErpRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

/* param customerPin: 客户pin(Optional) */
func (r *QueryCustPostFeeListForErpRequest) SetCustomerPin(customerPin string) {
    r.CustomerPin = &customerPin
}

/* param oneLevelDistributorName: 一级服务商名称(Optional) */
func (r *QueryCustPostFeeListForErpRequest) SetOneLevelDistributorName(oneLevelDistributorName string) {
    r.OneLevelDistributorName = &oneLevelDistributorName
}

/* param secondLevelDistributorName: 二级服务商名称(Optional) */
func (r *QueryCustPostFeeListForErpRequest) SetSecondLevelDistributorName(secondLevelDistributorName string) {
    r.SecondLevelDistributorName = &secondLevelDistributorName
}

/* param pageSize: 每页条数(Required) */
func (r *QueryCustPostFeeListForErpRequest) SetPageSize(pageSize int) {
    r.PageSize = pageSize
}

/* param pageIndex: 第几页(Required) */
func (r *QueryCustPostFeeListForErpRequest) SetPageIndex(pageIndex int) {
    r.PageIndex = pageIndex
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryCustPostFeeListForErpRequest) GetRegionId() string {
    return r.RegionId
}

type QueryCustPostFeeListForErpResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryCustPostFeeListForErpResult `json:"result"`
}

type QueryCustPostFeeListForErpResult struct {
    Pagination partner.Pagination `json:"pagination"`
    Result []partner.OperatorAchievement `json:"result"`
}