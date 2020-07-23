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

type DescribePrivilegeNewRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 用户pin  */
    Pin string `json:"pin"`
}

/*
 * param regionId: Region ID (Required)
 * param pin: 用户pin (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribePrivilegeNewRequest(
    regionId string,
    pin string,
) *DescribePrivilegeNewRequest {

	return &DescribePrivilegeNewRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/user:describePrivilegeNew",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Pin: pin,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pin: 用户pin (Required)
 */
func NewDescribePrivilegeNewRequestWithAllParams(
    regionId string,
    pin string,
) *DescribePrivilegeNewRequest {

    return &DescribePrivilegeNewRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/user:describePrivilegeNew",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Pin: pin,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribePrivilegeNewRequestWithoutParam() *DescribePrivilegeNewRequest {

    return &DescribePrivilegeNewRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/user:describePrivilegeNew",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribePrivilegeNewRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pin: 用户pin(Required) */
func (r *DescribePrivilegeNewRequest) SetPin(pin string) {
    r.Pin = pin
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribePrivilegeNewRequest) GetRegionId() string {
    return r.RegionId
}

type DescribePrivilegeNewResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribePrivilegeNewResult `json:"result"`
}

type DescribePrivilegeNewResult struct {
    Data int `json:"data"`
}