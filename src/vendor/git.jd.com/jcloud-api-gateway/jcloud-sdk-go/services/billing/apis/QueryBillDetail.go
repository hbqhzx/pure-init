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
    billing "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/billing/models"
)

type QueryBillDetailRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 计费开始时间  */
    StartTime string `json:"startTime"`

    /* 计费结束时间  */
    EndTime string `json:"endTime"`

    /* 产品线代码 (Optional) */
    AppCode *string `json:"appCode"`

    /* 产品代码 (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /* 计费类型 1、按配置 2、按用量 3、包年包月 4、按次 (Optional) */
    BillingType *int `json:"billingType"`

    /* 资源单id列表 (Optional) */
    ResourceIds []string `json:"resourceIds"`

    /* pageIndex (Optional) */
    PageIndex *int `json:"pageIndex"`

    /* pageSize (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: Region ID (Required)
 * param startTime: 计费开始时间 (Required)
 * param endTime: 计费结束时间 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryBillDetailRequest(
    regionId string,
    startTime string,
    endTime string,
) *QueryBillDetailRequest {

	return &QueryBillDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/billDetail:list",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        StartTime: startTime,
        EndTime: endTime,
	}
}

/*
 * param regionId: Region ID (Required)
 * param startTime: 计费开始时间 (Required)
 * param endTime: 计费结束时间 (Required)
 * param appCode: 产品线代码 (Optional)
 * param serviceCode: 产品代码 (Optional)
 * param billingType: 计费类型 1、按配置 2、按用量 3、包年包月 4、按次 (Optional)
 * param resourceIds: 资源单id列表 (Optional)
 * param pageIndex: pageIndex (Optional)
 * param pageSize: pageSize (Optional)
 */
func NewQueryBillDetailRequestWithAllParams(
    regionId string,
    startTime string,
    endTime string,
    appCode *string,
    serviceCode *string,
    billingType *int,
    resourceIds []string,
    pageIndex *int,
    pageSize *int,
) *QueryBillDetailRequest {

    return &QueryBillDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/billDetail:list",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        StartTime: startTime,
        EndTime: endTime,
        AppCode: appCode,
        ServiceCode: serviceCode,
        BillingType: billingType,
        ResourceIds: resourceIds,
        PageIndex: pageIndex,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryBillDetailRequestWithoutParam() *QueryBillDetailRequest {

    return &QueryBillDetailRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/billDetail:list",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *QueryBillDetailRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param startTime: 计费开始时间(Required) */
func (r *QueryBillDetailRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 计费结束时间(Required) */
func (r *QueryBillDetailRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

/* param appCode: 产品线代码(Optional) */
func (r *QueryBillDetailRequest) SetAppCode(appCode string) {
    r.AppCode = &appCode
}

/* param serviceCode: 产品代码(Optional) */
func (r *QueryBillDetailRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

/* param billingType: 计费类型 1、按配置 2、按用量 3、包年包月 4、按次(Optional) */
func (r *QueryBillDetailRequest) SetBillingType(billingType int) {
    r.BillingType = &billingType
}

/* param resourceIds: 资源单id列表(Optional) */
func (r *QueryBillDetailRequest) SetResourceIds(resourceIds []string) {
    r.ResourceIds = resourceIds
}

/* param pageIndex: pageIndex(Optional) */
func (r *QueryBillDetailRequest) SetPageIndex(pageIndex int) {
    r.PageIndex = &pageIndex
}

/* param pageSize: pageSize(Optional) */
func (r *QueryBillDetailRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryBillDetailRequest) GetRegionId() string {
    return r.RegionId
}

type QueryBillDetailResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryBillDetailResult `json:"result"`
}

type QueryBillDetailResult struct {
    Pagination billing.Pagination `json:"pagination"`
    Result []billing.BillSummary `json:"result"`
}