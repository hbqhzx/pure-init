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
    jdfusion "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/jdfusion/models"
)

type GetVmInstancesByIdRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 资源实例ID  */
    Id string `json:"id"`
}

/*
 * param regionId: 地域ID (Required)
 * param id: 资源实例ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetVmInstancesByIdRequest(
    regionId string,
    id string,
) *GetVmInstancesByIdRequest {

	return &GetVmInstancesByIdRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/vm_instances/{id}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Id: id,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param id: 资源实例ID (Required)
 */
func NewGetVmInstancesByIdRequestWithAllParams(
    regionId string,
    id string,
) *GetVmInstancesByIdRequest {

    return &GetVmInstancesByIdRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vm_instances/{id}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Id: id,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetVmInstancesByIdRequestWithoutParam() *GetVmInstancesByIdRequest {

    return &GetVmInstancesByIdRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vm_instances/{id}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GetVmInstancesByIdRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param id: 资源实例ID(Required) */
func (r *GetVmInstancesByIdRequest) SetId(id string) {
    r.Id = id
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetVmInstancesByIdRequest) GetRegionId() string {
    return r.RegionId
}

type GetVmInstancesByIdResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetVmInstancesByIdResult `json:"result"`
}

type GetVmInstancesByIdResult struct {
    Vm jdfusion.VmInfoDetail `json:"vm"`
}