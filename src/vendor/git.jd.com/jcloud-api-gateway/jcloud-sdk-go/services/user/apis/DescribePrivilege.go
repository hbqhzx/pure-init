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

type DescribePrivilegeRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 用户pin  */
    Pin string `json:"pin"`

    /* 产品线  */
    AppCode string `json:"appCode"`

    /* 具体产品  */
    ServiceCode string `json:"serviceCode"`
}

/*
 * param regionId: Region ID (Required)
 * param pin: 用户pin (Required)
 * param appCode: 产品线 (Required)
 * param serviceCode: 具体产品 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribePrivilegeRequest(
    regionId string,
    pin string,
    appCode string,
    serviceCode string,
) *DescribePrivilegeRequest {

	return &DescribePrivilegeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/user:describePrivilege",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Pin: pin,
        AppCode: appCode,
        ServiceCode: serviceCode,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pin: 用户pin (Required)
 * param appCode: 产品线 (Required)
 * param serviceCode: 具体产品 (Required)
 */
func NewDescribePrivilegeRequestWithAllParams(
    regionId string,
    pin string,
    appCode string,
    serviceCode string,
) *DescribePrivilegeRequest {

    return &DescribePrivilegeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/user:describePrivilege",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Pin: pin,
        AppCode: appCode,
        ServiceCode: serviceCode,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribePrivilegeRequestWithoutParam() *DescribePrivilegeRequest {

    return &DescribePrivilegeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/user:describePrivilege",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribePrivilegeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pin: 用户pin(Required) */
func (r *DescribePrivilegeRequest) SetPin(pin string) {
    r.Pin = pin
}

/* param appCode: 产品线(Required) */
func (r *DescribePrivilegeRequest) SetAppCode(appCode string) {
    r.AppCode = appCode
}

/* param serviceCode: 具体产品(Required) */
func (r *DescribePrivilegeRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = serviceCode
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribePrivilegeRequest) GetRegionId() string {
    return r.RegionId
}

type DescribePrivilegeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribePrivilegeResult `json:"result"`
}

type DescribePrivilegeResult struct {
    Data bool `json:"data"`
}