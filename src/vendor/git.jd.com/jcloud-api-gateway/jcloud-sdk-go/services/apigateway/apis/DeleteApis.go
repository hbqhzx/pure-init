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

type DeleteApisRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 分组ID  */
    ApiGroupId string `json:"apiGroupId"`

    /* 版本号  */
    Revision string `json:"revision"`

    /* 要删除API的ID集合，以,分隔 (Optional) */
    ApiIds *string `json:"apiIds"`
}

/*
 * param regionId: 地域ID (Required)
 * param apiGroupId: 分组ID (Required)
 * param revision: 版本号 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteApisRequest(
    regionId string,
    apiGroupId string,
    revision string,
) *DeleteApisRequest {

	return &DeleteApisRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/revision/{revision}/apis",
			Method:  "DELETE",
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
 * param apiIds: 要删除API的ID集合，以,分隔 (Optional)
 */
func NewDeleteApisRequestWithAllParams(
    regionId string,
    apiGroupId string,
    revision string,
    apiIds *string,
) *DeleteApisRequest {

    return &DeleteApisRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/revision/{revision}/apis",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ApiGroupId: apiGroupId,
        Revision: revision,
        ApiIds: apiIds,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteApisRequestWithoutParam() *DeleteApisRequest {

    return &DeleteApisRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/revision/{revision}/apis",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DeleteApisRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param apiGroupId: 分组ID(Required) */
func (r *DeleteApisRequest) SetApiGroupId(apiGroupId string) {
    r.ApiGroupId = apiGroupId
}

/* param revision: 版本号(Required) */
func (r *DeleteApisRequest) SetRevision(revision string) {
    r.Revision = revision
}

/* param apiIds: 要删除API的ID集合，以,分隔(Optional) */
func (r *DeleteApisRequest) SetApiIds(apiIds string) {
    r.ApiIds = &apiIds
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteApisRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteApisResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteApisResult `json:"result"`
}

type DeleteApisResult struct {
}