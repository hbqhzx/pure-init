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

type AdminQueryBillStatisticsInfoRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* 查询类别   1：资源账单   2：消费记录 (Optional) */
    QueryType *int `json:"queryType"`

    /* 用户pin (Optional) */
    Pin *string `json:"pin"`

    /* 应用码 (Optional) */
    AppCode *string `json:"appCode"`

    /* 服务码 (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /* 付费类型 (Optional) */
    BillingType *int `json:"billingType"`

    /* 付费类型 (Optional) */
    PayType *int `json:"payType"`

    /* 付费状态 (Optional) */
    PayState *int `json:"payState"`

    /* 1: 按账期 2：按消费时间 (Optional) */
    TimeType *int `json:"timeType"`

    /* 开始时间 (Optional) */
    StartTime *string `json:"startTime"`

    /* 结束时间 (Optional) */
    EndTime *string `json:"endTime"`

    /*  (Optional) */
    IgnoreZero *int `json:"ignoreZero"`

    /* 资源ID (Optional) */
    ResourceId *string `json:"resourceId"`

    /* 站点 (Optional) */
    Site *int `json:"site"`

    /* 角色 (Optional) */
    Role *int `json:"role"`

    /* 区域 (Optional) */
    Region *string `json:"region"`
}

/*
 * param regionId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAdminQueryBillStatisticsInfoRequest(
    regionId string,
) *AdminQueryBillStatisticsInfoRequest {

	return &AdminQueryBillStatisticsInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/accounting/adminQueryBillStatisticsInfo",
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
 * param pin: 用户pin (Optional)
 * param appCode: 应用码 (Optional)
 * param serviceCode: 服务码 (Optional)
 * param billingType: 付费类型 (Optional)
 * param payType: 付费类型 (Optional)
 * param payState: 付费状态 (Optional)
 * param timeType: 1: 按账期 2：按消费时间 (Optional)
 * param startTime: 开始时间 (Optional)
 * param endTime: 结束时间 (Optional)
 * param ignoreZero:  (Optional)
 * param resourceId: 资源ID (Optional)
 * param site: 站点 (Optional)
 * param role: 角色 (Optional)
 * param region: 区域 (Optional)
 */
func NewAdminQueryBillStatisticsInfoRequestWithAllParams(
    regionId string,
    queryType *int,
    pin *string,
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
) *AdminQueryBillStatisticsInfoRequest {

    return &AdminQueryBillStatisticsInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/accounting/adminQueryBillStatisticsInfo",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        QueryType: queryType,
        Pin: pin,
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
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAdminQueryBillStatisticsInfoRequestWithoutParam() *AdminQueryBillStatisticsInfoRequest {

    return &AdminQueryBillStatisticsInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/accounting/adminQueryBillStatisticsInfo",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *AdminQueryBillStatisticsInfoRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param queryType: 查询类别   1：资源账单   2：消费记录(Optional) */
func (r *AdminQueryBillStatisticsInfoRequest) SetQueryType(queryType int) {
    r.QueryType = &queryType
}

/* param pin: 用户pin(Optional) */
func (r *AdminQueryBillStatisticsInfoRequest) SetPin(pin string) {
    r.Pin = &pin
}

/* param appCode: 应用码(Optional) */
func (r *AdminQueryBillStatisticsInfoRequest) SetAppCode(appCode string) {
    r.AppCode = &appCode
}

/* param serviceCode: 服务码(Optional) */
func (r *AdminQueryBillStatisticsInfoRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

/* param billingType: 付费类型(Optional) */
func (r *AdminQueryBillStatisticsInfoRequest) SetBillingType(billingType int) {
    r.BillingType = &billingType
}

/* param payType: 付费类型(Optional) */
func (r *AdminQueryBillStatisticsInfoRequest) SetPayType(payType int) {
    r.PayType = &payType
}

/* param payState: 付费状态(Optional) */
func (r *AdminQueryBillStatisticsInfoRequest) SetPayState(payState int) {
    r.PayState = &payState
}

/* param timeType: 1: 按账期 2：按消费时间(Optional) */
func (r *AdminQueryBillStatisticsInfoRequest) SetTimeType(timeType int) {
    r.TimeType = &timeType
}

/* param startTime: 开始时间(Optional) */
func (r *AdminQueryBillStatisticsInfoRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 结束时间(Optional) */
func (r *AdminQueryBillStatisticsInfoRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param ignoreZero: (Optional) */
func (r *AdminQueryBillStatisticsInfoRequest) SetIgnoreZero(ignoreZero int) {
    r.IgnoreZero = &ignoreZero
}

/* param resourceId: 资源ID(Optional) */
func (r *AdminQueryBillStatisticsInfoRequest) SetResourceId(resourceId string) {
    r.ResourceId = &resourceId
}

/* param site: 站点(Optional) */
func (r *AdminQueryBillStatisticsInfoRequest) SetSite(site int) {
    r.Site = &site
}

/* param role: 角色(Optional) */
func (r *AdminQueryBillStatisticsInfoRequest) SetRole(role int) {
    r.Role = &role
}

/* param region: 区域(Optional) */
func (r *AdminQueryBillStatisticsInfoRequest) SetRegion(region string) {
    r.Region = &region
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AdminQueryBillStatisticsInfoRequest) GetRegionId() string {
    return r.RegionId
}

type AdminQueryBillStatisticsInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AdminQueryBillStatisticsInfoResult `json:"result"`
}

type AdminQueryBillStatisticsInfoResult struct {
    BillStatisticsInfoVo billing.BillStatisticsInfoVo `json:"billStatisticsInfoVo"`
}