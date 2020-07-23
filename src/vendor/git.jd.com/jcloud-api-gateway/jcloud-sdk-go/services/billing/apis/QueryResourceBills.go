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

type QueryResourceBillsRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* 查询类别   1：资源账单   2：消费记录 (Optional) */
    QueryType *int `json:"queryType"`

    /* appCode (Optional) */
    AppCode *string `json:"appCode"`

    /* serviceCode (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /* billingType (Optional) */
    BillingType *int `json:"billingType"`

    /* payType (Optional) */
    PayType *int `json:"payType"`

    /* payState (Optional) */
    PayState *int `json:"payState"`

    /* timeType (Optional) */
    TimeType *int `json:"timeType"`

    /* startTime (Optional) */
    StartTime *string `json:"startTime"`

    /* endTime (Optional) */
    EndTime *string `json:"endTime"`

    /* ignoreZero (Optional) */
    IgnoreZero *int `json:"ignoreZero"`

    /* resourceId (Optional) */
    ResourceId *string `json:"resourceId"`

    /* site (Optional) */
    Site *int `json:"site"`

    /* role (Optional) */
    Role *int `json:"role"`

    /* region (Optional) */
    Region *string `json:"region"`

    /* pageIndex (Optional) */
    PageIndex *int `json:"pageIndex"`

    /* pageSize (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryResourceBillsRequest(
    regionId string,
) *QueryResourceBillsRequest {

	return &QueryResourceBillsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/resourceBills",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId:  (Required)
 * param queryType: 查询类别   1：资源账单   2：消费记录 (Optional)
 * param appCode: appCode (Optional)
 * param serviceCode: serviceCode (Optional)
 * param billingType: billingType (Optional)
 * param payType: payType (Optional)
 * param payState: payState (Optional)
 * param timeType: timeType (Optional)
 * param startTime: startTime (Optional)
 * param endTime: endTime (Optional)
 * param ignoreZero: ignoreZero (Optional)
 * param resourceId: resourceId (Optional)
 * param site: site (Optional)
 * param role: role (Optional)
 * param region: region (Optional)
 * param pageIndex: pageIndex (Optional)
 * param pageSize: pageSize (Optional)
 */
func NewQueryResourceBillsRequestWithAllParams(
    regionId string,
    queryType *int,
    appCode *string,
    serviceCode *string,
    billingType *int,
    payType *int,
    payState *int,
    timeType *int,
    startTime *string,
    endTime *string,
    ignoreZero *int,
    resourceId *string,
    site *int,
    role *int,
    region *string,
    pageIndex *int,
    pageSize *int,
) *QueryResourceBillsRequest {

    return &QueryResourceBillsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/resourceBills",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        QueryType: queryType,
        AppCode: appCode,
        ServiceCode: serviceCode,
        BillingType: billingType,
        PayType: payType,
        PayState: payState,
        TimeType: timeType,
        StartTime: startTime,
        EndTime: endTime,
        IgnoreZero: ignoreZero,
        ResourceId: resourceId,
        Site: site,
        Role: role,
        Region: region,
        PageIndex: pageIndex,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryResourceBillsRequestWithoutParam() *QueryResourceBillsRequest {

    return &QueryResourceBillsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/resourceBills",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *QueryResourceBillsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param queryType: 查询类别   1：资源账单   2：消费记录(Optional) */
func (r *QueryResourceBillsRequest) SetQueryType(queryType int) {
    r.QueryType = &queryType
}

/* param appCode: appCode(Optional) */
func (r *QueryResourceBillsRequest) SetAppCode(appCode string) {
    r.AppCode = &appCode
}

/* param serviceCode: serviceCode(Optional) */
func (r *QueryResourceBillsRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

/* param billingType: billingType(Optional) */
func (r *QueryResourceBillsRequest) SetBillingType(billingType int) {
    r.BillingType = &billingType
}

/* param payType: payType(Optional) */
func (r *QueryResourceBillsRequest) SetPayType(payType int) {
    r.PayType = &payType
}

/* param payState: payState(Optional) */
func (r *QueryResourceBillsRequest) SetPayState(payState int) {
    r.PayState = &payState
}

/* param timeType: timeType(Optional) */
func (r *QueryResourceBillsRequest) SetTimeType(timeType int) {
    r.TimeType = &timeType
}

/* param startTime: startTime(Optional) */
func (r *QueryResourceBillsRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: endTime(Optional) */
func (r *QueryResourceBillsRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param ignoreZero: ignoreZero(Optional) */
func (r *QueryResourceBillsRequest) SetIgnoreZero(ignoreZero int) {
    r.IgnoreZero = &ignoreZero
}

/* param resourceId: resourceId(Optional) */
func (r *QueryResourceBillsRequest) SetResourceId(resourceId string) {
    r.ResourceId = &resourceId
}

/* param site: site(Optional) */
func (r *QueryResourceBillsRequest) SetSite(site int) {
    r.Site = &site
}

/* param role: role(Optional) */
func (r *QueryResourceBillsRequest) SetRole(role int) {
    r.Role = &role
}

/* param region: region(Optional) */
func (r *QueryResourceBillsRequest) SetRegion(region string) {
    r.Region = &region
}

/* param pageIndex: pageIndex(Optional) */
func (r *QueryResourceBillsRequest) SetPageIndex(pageIndex int) {
    r.PageIndex = &pageIndex
}

/* param pageSize: pageSize(Optional) */
func (r *QueryResourceBillsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryResourceBillsRequest) GetRegionId() string {
    return r.RegionId
}

type QueryResourceBillsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryResourceBillsResult `json:"result"`
}

type QueryResourceBillsResult struct {
    Pagination billing.Pagination `json:"pagination"`
    Result []billing.ResourceBillQueryResultItem `json:"result"`
}