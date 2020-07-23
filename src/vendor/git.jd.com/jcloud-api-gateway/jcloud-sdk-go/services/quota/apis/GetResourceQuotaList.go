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
    quota "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/quota/models"
)

type GetResourceQuotaListRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 业务线 (Optional) */
    AppCode *string `json:"appCode"`

    /* 可用量 (Optional) */
    AvailableAmount *int `json:"availableAmount"`

    /*  (Optional) */
    CountSql *bool `json:"countSql"`

    /*  (Optional) */
    Id *int `json:"id"`

    /*  (Optional) */
    OrderBy *string `json:"orderBy"`

    /*  (Optional) */
    PageNum *int `json:"pageNum"`

    /*  (Optional) */
    PageSize *int `json:"pageSize"`

    /*  (Optional) */
    PageSizeZero *bool `json:"pageSizeZero"`

    /* 预占量 (Optional) */
    PreOccupyAmount *int `json:"preOccupyAmount"`

    /*  (Optional) */
    Reasonable *bool `json:"reasonable"`

    /* 地域 (Optional) */
    Region *string `json:"region"`

    /* 资源默认配额 (Optional) */
    ResourceQuotaDefault *int `json:"resourceQuotaDefault"`

    /* 可售卖额 (Optional) */
    SaleableAmount *int `json:"saleableAmount"`

    /* 资源类型 (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /* 资源类型名称 (Optional) */
    ServiceName *string `json:"serviceName"`

    /* 库存总额 (Optional) */
    TotalInventory *int `json:"totalInventory"`

    /* 已用量 (Optional) */
    UsedAmount *int `json:"usedAmount"`

    /* 用户默认配额 (Optional) */
    UserQuotaDefault *int `json:"userQuotaDefault"`

    /* 预警量 (Optional) */
    WarningAmount *string `json:"warningAmount"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetResourceQuotaListRequest(
    regionId string,
) *GetResourceQuotaListRequest {

	return &GetResourceQuotaListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/resourceQuota/list",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param appCode: 业务线 (Optional)
 * param availableAmount: 可用量 (Optional)
 * param countSql:  (Optional)
 * param id:  (Optional)
 * param orderBy:  (Optional)
 * param pageNum:  (Optional)
 * param pageSize:  (Optional)
 * param pageSizeZero:  (Optional)
 * param preOccupyAmount: 预占量 (Optional)
 * param reasonable:  (Optional)
 * param region: 地域 (Optional)
 * param resourceQuotaDefault: 资源默认配额 (Optional)
 * param saleableAmount: 可售卖额 (Optional)
 * param serviceCode: 资源类型 (Optional)
 * param serviceName: 资源类型名称 (Optional)
 * param totalInventory: 库存总额 (Optional)
 * param usedAmount: 已用量 (Optional)
 * param userQuotaDefault: 用户默认配额 (Optional)
 * param warningAmount: 预警量 (Optional)
 */
func NewGetResourceQuotaListRequestWithAllParams(
    regionId string,
    appCode *string,
    availableAmount *int,
    countSql *bool,
    id *int,
    orderBy *string,
    pageNum *int,
    pageSize *int,
    pageSizeZero *bool,
    preOccupyAmount *int,
    reasonable *bool,
    region *string,
    resourceQuotaDefault *int,
    saleableAmount *int,
    serviceCode *string,
    serviceName *string,
    totalInventory *int,
    usedAmount *int,
    userQuotaDefault *int,
    warningAmount *string,
) *GetResourceQuotaListRequest {

    return &GetResourceQuotaListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/resourceQuota/list",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AppCode: appCode,
        AvailableAmount: availableAmount,
        CountSql: countSql,
        Id: id,
        OrderBy: orderBy,
        PageNum: pageNum,
        PageSize: pageSize,
        PageSizeZero: pageSizeZero,
        PreOccupyAmount: preOccupyAmount,
        Reasonable: reasonable,
        Region: region,
        ResourceQuotaDefault: resourceQuotaDefault,
        SaleableAmount: saleableAmount,
        ServiceCode: serviceCode,
        ServiceName: serviceName,
        TotalInventory: totalInventory,
        UsedAmount: usedAmount,
        UserQuotaDefault: userQuotaDefault,
        WarningAmount: warningAmount,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetResourceQuotaListRequestWithoutParam() *GetResourceQuotaListRequest {

    return &GetResourceQuotaListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/resourceQuota/list",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *GetResourceQuotaListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param appCode: 业务线(Optional) */
func (r *GetResourceQuotaListRequest) SetAppCode(appCode string) {
    r.AppCode = &appCode
}

/* param availableAmount: 可用量(Optional) */
func (r *GetResourceQuotaListRequest) SetAvailableAmount(availableAmount int) {
    r.AvailableAmount = &availableAmount
}

/* param countSql: (Optional) */
func (r *GetResourceQuotaListRequest) SetCountSql(countSql bool) {
    r.CountSql = &countSql
}

/* param id: (Optional) */
func (r *GetResourceQuotaListRequest) SetId(id int) {
    r.Id = &id
}

/* param orderBy: (Optional) */
func (r *GetResourceQuotaListRequest) SetOrderBy(orderBy string) {
    r.OrderBy = &orderBy
}

/* param pageNum: (Optional) */
func (r *GetResourceQuotaListRequest) SetPageNum(pageNum int) {
    r.PageNum = &pageNum
}

/* param pageSize: (Optional) */
func (r *GetResourceQuotaListRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param pageSizeZero: (Optional) */
func (r *GetResourceQuotaListRequest) SetPageSizeZero(pageSizeZero bool) {
    r.PageSizeZero = &pageSizeZero
}

/* param preOccupyAmount: 预占量(Optional) */
func (r *GetResourceQuotaListRequest) SetPreOccupyAmount(preOccupyAmount int) {
    r.PreOccupyAmount = &preOccupyAmount
}

/* param reasonable: (Optional) */
func (r *GetResourceQuotaListRequest) SetReasonable(reasonable bool) {
    r.Reasonable = &reasonable
}

/* param region: 地域(Optional) */
func (r *GetResourceQuotaListRequest) SetRegion(region string) {
    r.Region = &region
}

/* param resourceQuotaDefault: 资源默认配额(Optional) */
func (r *GetResourceQuotaListRequest) SetResourceQuotaDefault(resourceQuotaDefault int) {
    r.ResourceQuotaDefault = &resourceQuotaDefault
}

/* param saleableAmount: 可售卖额(Optional) */
func (r *GetResourceQuotaListRequest) SetSaleableAmount(saleableAmount int) {
    r.SaleableAmount = &saleableAmount
}

/* param serviceCode: 资源类型(Optional) */
func (r *GetResourceQuotaListRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

/* param serviceName: 资源类型名称(Optional) */
func (r *GetResourceQuotaListRequest) SetServiceName(serviceName string) {
    r.ServiceName = &serviceName
}

/* param totalInventory: 库存总额(Optional) */
func (r *GetResourceQuotaListRequest) SetTotalInventory(totalInventory int) {
    r.TotalInventory = &totalInventory
}

/* param usedAmount: 已用量(Optional) */
func (r *GetResourceQuotaListRequest) SetUsedAmount(usedAmount int) {
    r.UsedAmount = &usedAmount
}

/* param userQuotaDefault: 用户默认配额(Optional) */
func (r *GetResourceQuotaListRequest) SetUserQuotaDefault(userQuotaDefault int) {
    r.UserQuotaDefault = &userQuotaDefault
}

/* param warningAmount: 预警量(Optional) */
func (r *GetResourceQuotaListRequest) SetWarningAmount(warningAmount string) {
    r.WarningAmount = &warningAmount
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetResourceQuotaListRequest) GetRegionId() string {
    return r.RegionId
}

type GetResourceQuotaListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetResourceQuotaListResult `json:"result"`
}

type GetResourceQuotaListResult struct {
    ResultList []quota.ResourceQuotaVo `json:"resultList"`
    TotalCount int `json:"totalCount"`
    TotalPage int `json:"totalPage"`
}