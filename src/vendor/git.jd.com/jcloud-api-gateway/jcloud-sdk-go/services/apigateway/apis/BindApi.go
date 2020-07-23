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

type BindApiRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 接口ID  */
    ApiId string `json:"apiId"`

    /* 函数ID  */
    FunctionId string `json:"functionId"`

    /* 函数名称  */
    FunctionName string `json:"functionName"`

    /* 函数版本名称  */
    VersionName string `json:"versionName"`
}

/*
 * param regionId: 地域ID (Required)
 * param apiId: 接口ID (Required)
 * param functionId: 函数ID (Required)
 * param functionName: 函数名称 (Required)
 * param versionName: 函数版本名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewBindApiRequest(
    regionId string,
    apiId string,
    functionId string,
    functionName string,
    versionName string,
) *BindApiRequest {

	return &BindApiRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/function/{functionId}/apis/{apiId}",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ApiId: apiId,
        FunctionId: functionId,
        FunctionName: functionName,
        VersionName: versionName,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param apiId: 接口ID (Required)
 * param functionId: 函数ID (Required)
 * param functionName: 函数名称 (Required)
 * param versionName: 函数版本名称 (Required)
 */
func NewBindApiRequestWithAllParams(
    regionId string,
    apiId string,
    functionId string,
    functionName string,
    versionName string,
) *BindApiRequest {

    return &BindApiRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/function/{functionId}/apis/{apiId}",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ApiId: apiId,
        FunctionId: functionId,
        FunctionName: functionName,
        VersionName: versionName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewBindApiRequestWithoutParam() *BindApiRequest {

    return &BindApiRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/function/{functionId}/apis/{apiId}",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *BindApiRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param apiId: 接口ID(Required) */
func (r *BindApiRequest) SetApiId(apiId string) {
    r.ApiId = apiId
}

/* param functionId: 函数ID(Required) */
func (r *BindApiRequest) SetFunctionId(functionId string) {
    r.FunctionId = functionId
}

/* param functionName: 函数名称(Required) */
func (r *BindApiRequest) SetFunctionName(functionName string) {
    r.FunctionName = functionName
}

/* param versionName: 函数版本名称(Required) */
func (r *BindApiRequest) SetVersionName(versionName string) {
    r.VersionName = versionName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r BindApiRequest) GetRegionId() string {
    return r.RegionId
}

type BindApiResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result BindApiResult `json:"result"`
}

type BindApiResult struct {
}