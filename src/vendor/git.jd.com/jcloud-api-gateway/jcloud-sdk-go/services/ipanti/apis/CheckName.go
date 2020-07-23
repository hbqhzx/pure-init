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

type CheckNameRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 待检测实例名称  */
    Name string `json:"name"`
}

/*
 * param regionId: Region ID (Required)
 * param name: 待检测实例名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCheckNameRequest(
    regionId string,
    name string,
) *CheckNameRequest {

	return &CheckNameRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/checkName",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Name: name,
	}
}

/*
 * param regionId: Region ID (Required)
 * param name: 待检测实例名称 (Required)
 */
func NewCheckNameRequestWithAllParams(
    regionId string,
    name string,
) *CheckNameRequest {

    return &CheckNameRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/checkName",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Name: name,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCheckNameRequestWithoutParam() *CheckNameRequest {

    return &CheckNameRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/checkName",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CheckNameRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param name: 待检测实例名称(Required) */
func (r *CheckNameRequest) SetName(name string) {
    r.Name = name
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CheckNameRequest) GetRegionId() string {
    return r.RegionId
}

type CheckNameResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CheckNameResult `json:"result"`
}

type CheckNameResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}