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

type RecreateRegistryRequest struct {

    core.JDCloudRequest

    /* 可用区id  */
    RegionId string `json:"regionId"`

    /* 集群id  */
    RegistryId string `json:"registryId"`
}

/*
 * param regionId: 可用区id (Required)
 * param registryId: 集群id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRecreateRegistryRequest(
    regionId string,
    registryId string,
) *RecreateRegistryRequest {

	return &RecreateRegistryRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/registries/{registryId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        RegistryId: registryId,
	}
}

/*
 * param regionId: 可用区id (Required)
 * param registryId: 集群id (Required)
 */
func NewRecreateRegistryRequestWithAllParams(
    regionId string,
    registryId string,
) *RecreateRegistryRequest {

    return &RecreateRegistryRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/registries/{registryId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        RegistryId: registryId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewRecreateRegistryRequestWithoutParam() *RecreateRegistryRequest {

    return &RecreateRegistryRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/registries/{registryId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 可用区id(Required) */
func (r *RecreateRegistryRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param registryId: 集群id(Required) */
func (r *RecreateRegistryRequest) SetRegistryId(registryId string) {
    r.RegistryId = registryId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RecreateRegistryRequest) GetRegionId() string {
    return r.RegionId
}

type RecreateRegistryResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RecreateRegistryResult `json:"result"`
}

type RecreateRegistryResult struct {
    RecoverResult string `json:"recoverResult"`
}