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
)

type HandleAutoPayRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /*  (Optional) */
    AppCode *string `json:"appCode"`

    /*  (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /*  (Optional) */
    DataCenter *string `json:"dataCenter"`

    /*  (Optional) */
    TimeSpan *int `json:"timeSpan"`

    /*  (Optional) */
    TimeUnit *int `json:"timeUnit"`

    /*  (Optional) */
    AllAutoPay *string `json:"allAutoPay"`

    /*  (Optional) */
    ResourceIdList []string `json:"resourceIdList"`
}

/*
 * param regionId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewHandleAutoPayRequest(
    regionId string,
) *HandleAutoPayRequest {

	return &HandleAutoPayRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/autoPay",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId:  (Required)
 * param appCode:  (Optional)
 * param serviceCode:  (Optional)
 * param dataCenter:  (Optional)
 * param timeSpan:  (Optional)
 * param timeUnit:  (Optional)
 * param allAutoPay:  (Optional)
 * param resourceIdList:  (Optional)
 */
func NewHandleAutoPayRequestWithAllParams(
    regionId string,
    appCode *string,
    serviceCode *string,
    dataCenter *string,
    timeSpan *int,
    timeUnit *int,
    allAutoPay *string,
    resourceIdList []string,
) *HandleAutoPayRequest {

    return &HandleAutoPayRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/autoPay",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AppCode: appCode,
        ServiceCode: serviceCode,
        DataCenter: dataCenter,
        TimeSpan: timeSpan,
        TimeUnit: timeUnit,
        AllAutoPay: allAutoPay,
        ResourceIdList: resourceIdList,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewHandleAutoPayRequestWithoutParam() *HandleAutoPayRequest {

    return &HandleAutoPayRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/autoPay",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *HandleAutoPayRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param appCode: (Optional) */
func (r *HandleAutoPayRequest) SetAppCode(appCode string) {
    r.AppCode = &appCode
}

/* param serviceCode: (Optional) */
func (r *HandleAutoPayRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

/* param dataCenter: (Optional) */
func (r *HandleAutoPayRequest) SetDataCenter(dataCenter string) {
    r.DataCenter = &dataCenter
}

/* param timeSpan: (Optional) */
func (r *HandleAutoPayRequest) SetTimeSpan(timeSpan int) {
    r.TimeSpan = &timeSpan
}

/* param timeUnit: (Optional) */
func (r *HandleAutoPayRequest) SetTimeUnit(timeUnit int) {
    r.TimeUnit = &timeUnit
}

/* param allAutoPay: (Optional) */
func (r *HandleAutoPayRequest) SetAllAutoPay(allAutoPay string) {
    r.AllAutoPay = &allAutoPay
}

/* param resourceIdList: (Optional) */
func (r *HandleAutoPayRequest) SetResourceIdList(resourceIdList []string) {
    r.ResourceIdList = resourceIdList
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r HandleAutoPayRequest) GetRegionId() string {
    return r.RegionId
}

type HandleAutoPayResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result HandleAutoPayResult `json:"result"`
}

type HandleAutoPayResult struct {
    StringResult string `json:"stringResult"`
}