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

type GetUnpaidOrderCountRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /*   */
    Role string `json:"role"`
}

/*
 * param regionId: Region ID (Required)
 * param role:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetUnpaidOrderCountRequest(
    regionId string,
    role string,
) *GetUnpaidOrderCountRequest {

	return &GetUnpaidOrderCountRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/unpaidorder",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Role: role,
	}
}

/*
 * param regionId: Region ID (Required)
 * param role:  (Required)
 */
func NewGetUnpaidOrderCountRequestWithAllParams(
    regionId string,
    role string,
) *GetUnpaidOrderCountRequest {

    return &GetUnpaidOrderCountRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/unpaidorder",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Role: role,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetUnpaidOrderCountRequestWithoutParam() *GetUnpaidOrderCountRequest {

    return &GetUnpaidOrderCountRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/unpaidorder",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *GetUnpaidOrderCountRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param role: (Required) */
func (r *GetUnpaidOrderCountRequest) SetRole(role string) {
    r.Role = role
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetUnpaidOrderCountRequest) GetRegionId() string {
    return r.RegionId
}

type GetUnpaidOrderCountResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetUnpaidOrderCountResult `json:"result"`
}

type GetUnpaidOrderCountResult struct {
    Count int `json:"count"`
}