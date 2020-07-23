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

type UpdateApiByNameRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 分组ID  */
    ApiGroupId string `json:"apiGroupId"`

    /* 版本号  */
    Revision string `json:"revision"`

    /*   */
    ApiName string `json:"apiName"`

    /* api (Optional) */
    Api *apigateway.CreateApi `json:"api"`
}

/*
 * param regionId: 地域ID (Required)
 * param apiGroupId: 分组ID (Required)
 * param revision: 版本号 (Required)
 * param apiName:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateApiByNameRequest(
    regionId string,
    apiGroupId string,
    revision string,
    apiName string,
) *UpdateApiByNameRequest {

	return &UpdateApiByNameRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/revision/{revision}/apisName/{apiName}",
			Method:  "PUT",
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
 * param api: api (Optional)
 */
func NewUpdateApiByNameRequestWithAllParams(
    regionId string,
    apiGroupId string,
    revision string,
    apiName string,
    api *apigateway.CreateApi,
) *UpdateApiByNameRequest {

    return &UpdateApiByNameRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/revision/{revision}/apisName/{apiName}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ApiGroupId: apiGroupId,
        Revision: revision,
        ApiName: apiName,
        Api: api,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateApiByNameRequestWithoutParam() *UpdateApiByNameRequest {

    return &UpdateApiByNameRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/revision/{revision}/apisName/{apiName}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *UpdateApiByNameRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param apiGroupId: 分组ID(Required) */
func (r *UpdateApiByNameRequest) SetApiGroupId(apiGroupId string) {
    r.ApiGroupId = apiGroupId
}

/* param revision: 版本号(Required) */
func (r *UpdateApiByNameRequest) SetRevision(revision string) {
    r.Revision = revision
}

/* param apiName: (Required) */
func (r *UpdateApiByNameRequest) SetApiName(apiName string) {
    r.ApiName = apiName
}

/* param api: api(Optional) */
func (r *UpdateApiByNameRequest) SetApi(api *apigateway.CreateApi) {
    r.Api = api
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateApiByNameRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateApiByNameResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateApiByNameResult `json:"result"`
}

type UpdateApiByNameResult struct {
}