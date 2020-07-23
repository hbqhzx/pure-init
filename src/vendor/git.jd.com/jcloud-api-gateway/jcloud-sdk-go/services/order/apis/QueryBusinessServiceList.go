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

type QueryBusinessServiceListRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /*  (Optional) */
    BusinessCode *int `json:"businessCode"`

    /*  (Optional) */
    SiteType *int `json:"siteType"`

    /*  (Optional) */
    Data *int `json:"data"`

    /*  (Optional) */
    AppCode *string `json:"appCode"`

    /*  (Optional) */
    ServiceCode *string `json:"serviceCode"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryBusinessServiceListRequest(
    regionId string,
) *QueryBusinessServiceListRequest {

	return &QueryBusinessServiceListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/order:queryBusinessServiceList",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param businessCode:  (Optional)
 * param siteType:  (Optional)
 * param data:  (Optional)
 * param appCode:  (Optional)
 * param serviceCode:  (Optional)
 */
func NewQueryBusinessServiceListRequestWithAllParams(
    regionId string,
    businessCode *int,
    siteType *int,
    data *int,
    appCode *string,
    serviceCode *string,
) *QueryBusinessServiceListRequest {

    return &QueryBusinessServiceListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/order:queryBusinessServiceList",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        BusinessCode: businessCode,
        SiteType: siteType,
        Data: data,
        AppCode: appCode,
        ServiceCode: serviceCode,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryBusinessServiceListRequestWithoutParam() *QueryBusinessServiceListRequest {

    return &QueryBusinessServiceListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/order:queryBusinessServiceList",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *QueryBusinessServiceListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param businessCode: (Optional) */
func (r *QueryBusinessServiceListRequest) SetBusinessCode(businessCode int) {
    r.BusinessCode = &businessCode
}

/* param siteType: (Optional) */
func (r *QueryBusinessServiceListRequest) SetSiteType(siteType int) {
    r.SiteType = &siteType
}

/* param data: (Optional) */
func (r *QueryBusinessServiceListRequest) SetData(data int) {
    r.Data = &data
}

/* param appCode: (Optional) */
func (r *QueryBusinessServiceListRequest) SetAppCode(appCode string) {
    r.AppCode = &appCode
}

/* param serviceCode: (Optional) */
func (r *QueryBusinessServiceListRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryBusinessServiceListRequest) GetRegionId() string {
    return r.RegionId
}

type QueryBusinessServiceListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryBusinessServiceListResult `json:"result"`
}

type QueryBusinessServiceListResult struct {
    Data []order.BusinessServiceVo `json:"data"`
}