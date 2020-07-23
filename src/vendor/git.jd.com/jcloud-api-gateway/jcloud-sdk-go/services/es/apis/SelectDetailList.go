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
    es "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/es/models"
)

type SelectDetailListRequest struct {

    core.JDCloudRequest

    /* regionId  */
    RegionId string `json:"regionId"`

    /* 资源id列表（多个以 , 分隔）  */
    ResourceList string `json:"resourceList"`
}

/*
 * param regionId: regionId (Required)
 * param resourceList: 资源id列表（多个以 , 分隔） (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSelectDetailListRequest(
    regionId string,
    resourceList string,
) *SelectDetailListRequest {

	return &SelectDetailListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instanceNames",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ResourceList: resourceList,
	}
}

/*
 * param regionId: regionId (Required)
 * param resourceList: 资源id列表（多个以 , 分隔） (Required)
 */
func NewSelectDetailListRequestWithAllParams(
    regionId string,
    resourceList string,
) *SelectDetailListRequest {

    return &SelectDetailListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instanceNames",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ResourceList: resourceList,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSelectDetailListRequestWithoutParam() *SelectDetailListRequest {

    return &SelectDetailListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instanceNames",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: regionId(Required) */
func (r *SelectDetailListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param resourceList: 资源id列表（多个以 , 分隔）(Required) */
func (r *SelectDetailListRequest) SetResourceList(resourceList string) {
    r.ResourceList = resourceList
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SelectDetailListRequest) GetRegionId() string {
    return r.RegionId
}

type SelectDetailListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SelectDetailListResult `json:"result"`
}

type SelectDetailListResult struct {
    Data []es.InstanceName `json:"data"`
}