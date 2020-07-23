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

type GetHooksRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 代码库名称  */
    Name string `json:"name"`

    /* 代码类型,如GITHUB、CODECOMMIT  */
    Type string `json:"type"`
}

/*
 * param regionId: Region ID (Required)
 * param name: 代码库名称 (Required)
 * param type_: 代码类型,如GITHUB、CODECOMMIT (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetHooksRequest(
    regionId string,
    name string,
    type_ string,
) *GetHooksRequest {

	return &GetHooksRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/repos/{name}/hooks",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Name: name,
        Type: type_,
	}
}

/*
 * param regionId: Region ID (Required)
 * param name: 代码库名称 (Required)
 * param type_: 代码类型,如GITHUB、CODECOMMIT (Required)
 */
func NewGetHooksRequestWithAllParams(
    regionId string,
    name string,
    type_ string,
) *GetHooksRequest {

    return &GetHooksRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/repos/{name}/hooks",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Name: name,
        Type: type_,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetHooksRequestWithoutParam() *GetHooksRequest {

    return &GetHooksRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/repos/{name}/hooks",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *GetHooksRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param name: 代码库名称(Required) */
func (r *GetHooksRequest) SetName(name string) {
    r.Name = name
}

/* param type_: 代码类型,如GITHUB、CODECOMMIT(Required) */
func (r *GetHooksRequest) SetType(type_ string) {
    r.Type = type_
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetHooksRequest) GetRegionId() string {
    return r.RegionId
}

type GetHooksResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetHooksResult `json:"result"`
}

type GetHooksResult struct {
    Hooks []interface{} `json:"hooks"`
}