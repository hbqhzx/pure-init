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

type DescribeBaseInfoRequest struct {

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
func NewDescribeBaseInfoRequest(
    regionId string,
    pin string,
) *DescribeBaseInfoRequest {

	return &DescribeBaseInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/user:describeBaseInfo",
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
func NewDescribeBaseInfoRequestWithAllParams(
    regionId string,
    pin string,
) *DescribeBaseInfoRequest {

    return &DescribeBaseInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/user:describeBaseInfo",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Pin: pin,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeBaseInfoRequestWithoutParam() *DescribeBaseInfoRequest {

    return &DescribeBaseInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/user:describeBaseInfo",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeBaseInfoRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pin: 用户pin(Required) */
func (r *DescribeBaseInfoRequest) SetPin(pin string) {
    r.Pin = pin
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeBaseInfoRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeBaseInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeBaseInfoResult `json:"result"`
}

type DescribeBaseInfoResult struct {
    Pin string `json:"pin"`
    Enabled int `json:"enabled"`
    UserType int `json:"userType"`
    PersonAuthStatus int `json:"personAuthStatus"`
    CompanyaAuthStatus int `json:"companyaAuthStatus"`
    AccountBalance int `json:"accountBalance"`
    AccountId int64 `json:"accountId"`
    Meta string `json:"meta"`
    ActiveStatus int `json:"activeStatus"`
    CscPhone string `json:"cscPhone"`
    CscEmail string `json:"cscEmail"`
    Able string `json:"able"`
    NickName string `json:"nickName"`
    Vip string `json:"vip"`
    Name string `json:"name"`
    UserGroup int `json:"userGroup"`
    UserGroupName string `json:"userGroupName"`
    CompanyName string `json:"companyName"`
}