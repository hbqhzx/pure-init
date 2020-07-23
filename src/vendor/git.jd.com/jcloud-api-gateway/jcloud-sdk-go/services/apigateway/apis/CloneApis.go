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
    apigateway "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/apigateway/models"
)

type CloneApisRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 分组ID  */
    ApiGroupId string `json:"apiGroupId"`

    /* 版本号  */
    Revision string `json:"revision"`

    /* 待复制Api (Optional) */
    ApiSource *apigateway.CreateApi `json:"apiSource"`
}

/*
 * param regionId: 地域ID (Required)
 * param apiGroupId: 分组ID (Required)
 * param revision: 版本号 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCloneApisRequest(
    regionId string,
    apiGroupId string,
    revision string,
) *CloneApisRequest {

	return &CloneApisRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/revision/{revision}/apis:clone",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ApiGroupId: apiGroupId,
        Revision: revision,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param apiGroupId: 分组ID (Required)
 * param revision: 版本号 (Required)
 * param apiSource: 待复制Api (Optional)
 */
func NewCloneApisRequestWithAllParams(
    regionId string,
    apiGroupId string,
    revision string,
    apiSource *apigateway.CreateApi,
) *CloneApisRequest {

    return &CloneApisRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/revision/{revision}/apis:clone",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ApiGroupId: apiGroupId,
        Revision: revision,
        ApiSource: apiSource,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCloneApisRequestWithoutParam() *CloneApisRequest {

    return &CloneApisRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/revision/{revision}/apis:clone",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CloneApisRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param apiGroupId: 分组ID(Required) */
func (r *CloneApisRequest) SetApiGroupId(apiGroupId string) {
    r.ApiGroupId = apiGroupId
}

/* param revision: 版本号(Required) */
func (r *CloneApisRequest) SetRevision(revision string) {
    r.Revision = revision
}

/* param apiSource: 待复制Api(Optional) */
func (r *CloneApisRequest) SetApiSource(apiSource *apigateway.CreateApi) {
    r.ApiSource = apiSource
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CloneApisRequest) GetRegionId() string {
    return r.RegionId
}

type CloneApisResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CloneApisResult `json:"result"`
}

type CloneApisResult struct {
    ApiId string `json:"apiId"`
}