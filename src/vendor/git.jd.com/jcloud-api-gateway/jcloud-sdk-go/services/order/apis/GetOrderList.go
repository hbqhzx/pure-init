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
    order "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/order/models"
)

type GetOrderListRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 下单开始时间 (Optional) */
    StartTime *string `json:"startTime"`

    /* 下单结束时间 (Optional) */
    EndTime *string `json:"endTime"`

    /* 产品线 (Optional) */
    AppCode *string `json:"appCode"`

    /* 产品类型 (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /* 流程类型 (Optional) */
    ProcessType *int `json:"processType"`

    /* 站点类型 (Optional) */
    SiteType *int `json:"siteType"`

    /* 自营类型 (Optional) */
    SelfSupportType *int `json:"selfSupportType"`

    /* 订单状态 (Optional) */
    Status *int `json:"status"`

    /* 订单类型 (Optional) */
    OrderType *int `json:"orderType"`

    /* 计费类型 (Optional) */
    ChargeMode *int `json:"chargeMode"`

    /*  (Optional) */
    Start *int `json:"start"`

    /*  (Optional) */
    PageNum *int `json:"pageNum"`

    /*  (Optional) */
    PageSize *int `json:"pageSize"`

    /*  (Optional) */
    TotalCount *int `json:"totalCount"`

    /* 付费类型 (Optional) */
    PayType *int `json:"payType"`

    /* 订单编号 (Optional) */
    OrderNumber *string `json:"orderNumber"`

    /*  (Optional) */
    PaymentNumber *string `json:"paymentNumber"`

    /*  (Optional) */
    ResourceId *string `json:"resourceId"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetOrderListRequest(
    regionId string,
) *GetOrderListRequest {

	return &GetOrderListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/orders",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param startTime: 下单开始时间 (Optional)
 * param endTime: 下单结束时间 (Optional)
 * param appCode: 产品线 (Optional)
 * param serviceCode: 产品类型 (Optional)
 * param processType: 流程类型 (Optional)
 * param siteType: 站点类型 (Optional)
 * param selfSupportType: 自营类型 (Optional)
 * param status: 订单状态 (Optional)
 * param orderType: 订单类型 (Optional)
 * param chargeMode: 计费类型 (Optional)
 * param start:  (Optional)
 * param pageNum:  (Optional)
 * param pageSize:  (Optional)
 * param totalCount:  (Optional)
 * param payType: 付费类型 (Optional)
 * param orderNumber: 订单编号 (Optional)
 * param paymentNumber:  (Optional)
 * param resourceId:  (Optional)
 */
func NewGetOrderListRequestWithAllParams(
    regionId string,
    startTime *string,
    endTime *string,
    appCode *string,
    serviceCode *string,
    processType *int,
    siteType *int,
    selfSupportType *int,
    status *int,
    orderType *int,
    chargeMode *int,
    start *int,
    pageNum *int,
    pageSize *int,
    totalCount *int,
    payType *int,
    orderNumber *string,
    paymentNumber *string,
    resourceId *string,
) *GetOrderListRequest {

    return &GetOrderListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/orders",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        StartTime: startTime,
        EndTime: endTime,
        AppCode: appCode,
        ServiceCode: serviceCode,
        ProcessType: processType,
        SiteType: siteType,
        SelfSupportType: selfSupportType,
        Status: status,
        OrderType: orderType,
        ChargeMode: chargeMode,
        Start: start,
        PageNum: pageNum,
        PageSize: pageSize,
        TotalCount: totalCount,
        PayType: payType,
        OrderNumber: orderNumber,
        PaymentNumber: paymentNumber,
        ResourceId: resourceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetOrderListRequestWithoutParam() *GetOrderListRequest {

    return &GetOrderListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/orders",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *GetOrderListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param startTime: 下单开始时间(Optional) */
func (r *GetOrderListRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 下单结束时间(Optional) */
func (r *GetOrderListRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param appCode: 产品线(Optional) */
func (r *GetOrderListRequest) SetAppCode(appCode string) {
    r.AppCode = &appCode
}

/* param serviceCode: 产品类型(Optional) */
func (r *GetOrderListRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

/* param processType: 流程类型(Optional) */
func (r *GetOrderListRequest) SetProcessType(processType int) {
    r.ProcessType = &processType
}

/* param siteType: 站点类型(Optional) */
func (r *GetOrderListRequest) SetSiteType(siteType int) {
    r.SiteType = &siteType
}

/* param selfSupportType: 自营类型(Optional) */
func (r *GetOrderListRequest) SetSelfSupportType(selfSupportType int) {
    r.SelfSupportType = &selfSupportType
}

/* param status: 订单状态(Optional) */
func (r *GetOrderListRequest) SetStatus(status int) {
    r.Status = &status
}

/* param orderType: 订单类型(Optional) */
func (r *GetOrderListRequest) SetOrderType(orderType int) {
    r.OrderType = &orderType
}

/* param chargeMode: 计费类型(Optional) */
func (r *GetOrderListRequest) SetChargeMode(chargeMode int) {
    r.ChargeMode = &chargeMode
}

/* param start: (Optional) */
func (r *GetOrderListRequest) SetStart(start int) {
    r.Start = &start
}

/* param pageNum: (Optional) */
func (r *GetOrderListRequest) SetPageNum(pageNum int) {
    r.PageNum = &pageNum
}

/* param pageSize: (Optional) */
func (r *GetOrderListRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param totalCount: (Optional) */
func (r *GetOrderListRequest) SetTotalCount(totalCount int) {
    r.TotalCount = &totalCount
}

/* param payType: 付费类型(Optional) */
func (r *GetOrderListRequest) SetPayType(payType int) {
    r.PayType = &payType
}

/* param orderNumber: 订单编号(Optional) */
func (r *GetOrderListRequest) SetOrderNumber(orderNumber string) {
    r.OrderNumber = &orderNumber
}

/* param paymentNumber: (Optional) */
func (r *GetOrderListRequest) SetPaymentNumber(paymentNumber string) {
    r.PaymentNumber = &paymentNumber
}

/* param resourceId: (Optional) */
func (r *GetOrderListRequest) SetResourceId(resourceId string) {
    r.ResourceId = &resourceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetOrderListRequest) GetRegionId() string {
    return r.RegionId
}

type GetOrderListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetOrderListResult `json:"result"`
}

type GetOrderListResult struct {
    OrdersDetail order.OrdersDetail `json:"ordersDetail"`
}