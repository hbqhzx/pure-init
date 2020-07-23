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
    renewal "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/renewal/models"
)

type ResourceListRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /*  (Optional) */
    AppCode *string `json:"appCode"`

    /*  (Optional) */
    DataCenter *string `json:"dataCenter"`

    /*  (Optional) */
    ExpireTime *string `json:"expireTime"`

    /*  (Optional) */
    BillingType *string `json:"billingType"`

    /*  (Optional) */
    AutoRenew *int `json:"autoRenew"`

    /*  (Optional) */
    PageNumber *int `json:"pageNumber"`

    /*  (Optional) */
    PageSize *int `json:"pageSize"`

    /*  (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /*  (Optional) */
    ResourceName *string `json:"resourceName"`

    /*  (Optional) */
    ResourceID *string `json:"resourceID"`

    /*  (Optional) */
    IpAddress *string `json:"ipAddress"`
}

/*
 * param regionId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewResourceListRequest(
    regionId string,
) *ResourceListRequest {

	return &ResourceListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/resourceList",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId:  (Required)
 * param appCode:  (Optional)
 * param dataCenter:  (Optional)
 * param expireTime:  (Optional)
 * param billingType:  (Optional)
 * param autoRenew:  (Optional)
 * param pageNumber:  (Optional)
 * param pageSize:  (Optional)
 * param serviceCode:  (Optional)
 * param resourceName:  (Optional)
 * param resourceID:  (Optional)
 * param ipAddress:  (Optional)
 */
func NewResourceListRequestWithAllParams(
    regionId string,
    appCode *string,
    dataCenter *string,
    expireTime *string,
    billingType *string,
    autoRenew *int,
    pageNumber *int,
    pageSize *int,
    serviceCode *string,
    resourceName *string,
    resourceID *string,
    ipAddress *string,
) *ResourceListRequest {

    return &ResourceListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/resourceList",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AppCode: appCode,
        DataCenter: dataCenter,
        ExpireTime: expireTime,
        BillingType: billingType,
        AutoRenew: autoRenew,
        PageNumber: pageNumber,
        PageSize: pageSize,
        ServiceCode: serviceCode,
        ResourceName: resourceName,
        ResourceID: resourceID,
        IpAddress: ipAddress,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewResourceListRequestWithoutParam() *ResourceListRequest {

    return &ResourceListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/resourceList",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *ResourceListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param appCode: (Optional) */
func (r *ResourceListRequest) SetAppCode(appCode string) {
    r.AppCode = &appCode
}

/* param dataCenter: (Optional) */
func (r *ResourceListRequest) SetDataCenter(dataCenter string) {
    r.DataCenter = &dataCenter
}

/* param expireTime: (Optional) */
func (r *ResourceListRequest) SetExpireTime(expireTime string) {
    r.ExpireTime = &expireTime
}

/* param billingType: (Optional) */
func (r *ResourceListRequest) SetBillingType(billingType string) {
    r.BillingType = &billingType
}

/* param autoRenew: (Optional) */
func (r *ResourceListRequest) SetAutoRenew(autoRenew int) {
    r.AutoRenew = &autoRenew
}

/* param pageNumber: (Optional) */
func (r *ResourceListRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: (Optional) */
func (r *ResourceListRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param serviceCode: (Optional) */
func (r *ResourceListRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

/* param resourceName: (Optional) */
func (r *ResourceListRequest) SetResourceName(resourceName string) {
    r.ResourceName = &resourceName
}

/* param resourceID: (Optional) */
func (r *ResourceListRequest) SetResourceID(resourceID string) {
    r.ResourceID = &resourceID
}

/* param ipAddress: (Optional) */
func (r *ResourceListRequest) SetIpAddress(ipAddress string) {
    r.IpAddress = &ipAddress
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ResourceListRequest) GetRegionId() string {
    return r.RegionId
}

type ResourceListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ResourceListResult `json:"result"`
}

type ResourceListResult struct {
    ListQueries []renewal.ListQuery `json:"listQueries"`
    TotalCount int `json:"totalCount"`
}