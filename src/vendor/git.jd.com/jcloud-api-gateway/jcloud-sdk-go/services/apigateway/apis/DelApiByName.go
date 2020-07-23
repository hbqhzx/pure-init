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

type DelApiByNameRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 分组ID  */
    ApiGroupId string `json:"apiGroupId"`

    /* 版本号  */
    Revision string `json:"revision"`

    /*   */
    ApiName string `json:"apiName"`
}

/*
 * param regionId: 地域ID (Required)
 * param apiGroupId: 分组ID (Required)
 * param revision: 版本号 (Required)
 * param apiName:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDelApiByNameRequest(
    regionId string,
    apiGroupId string,
    revision string,
    apiName string,
) *DelApiByNameRequest {

	return &DelApiByNameRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/revision/{revision}/apisName/{apiName}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ApiGroupId: apiGroupId,
        Revision: revision,
        ApiName: apiName,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param apiGroupId: 分组ID (Required)
 * param revision: 版本号 (Required)
 * param apiName:  (Required)
 */
func NewDelApiByNameRequestWithAllParams(
    regionId string,
    apiGroupId string,
    revision string,
    apiName string,
) *DelApiByNameRequest {

    return &DelApiByNameRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/revision/{revision}/apisName/{apiName}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ApiGroupId: apiGroupId,
        Revision: revision,
        ApiName: apiName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDelApiByNameRequestWithoutParam() *DelApiByNameRequest {

    return &DelApiByNameRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/revision/{revision}/apisName/{apiName}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DelApiByNameRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param apiGroupId: 分组ID(Required) */
func (r *DelApiByNameRequest) SetApiGroupId(apiGroupId string) {
    r.ApiGroupId = apiGroupId
}

/* param revision: 版本号(Required) */
func (r *DelApiByNameRequest) SetRevision(revision string) {
    r.Revision = revision
}

/* param apiName: (Required) */
func (r *DelApiByNameRequest) SetApiName(apiName string) {
    r.ApiName = apiName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DelApiByNameRequest) GetRegionId() string {
    return r.RegionId
}

type DelApiByNameResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DelApiByNameResult `json:"result"`
}

type DelApiByNameResult struct {
}