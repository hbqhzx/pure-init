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

type DeleteAppRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /*   */
    ClientId string `json:"clientId"`
}

/*
 * param regionId:  (Required)
 * param clientId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteAppRequest(
    regionId string,
    clientId string,
) *DeleteAppRequest {

	return &DeleteAppRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/app/{clientId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ClientId: clientId,
	}
}

/*
 * param regionId:  (Required)
 * param clientId:  (Required)
 */
func NewDeleteAppRequestWithAllParams(
    regionId string,
    clientId string,
) *DeleteAppRequest {

    return &DeleteAppRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/app/{clientId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ClientId: clientId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteAppRequestWithoutParam() *DeleteAppRequest {

    return &DeleteAppRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/app/{clientId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *DeleteAppRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param clientId: (Required) */
func (r *DeleteAppRequest) SetClientId(clientId string) {
    r.ClientId = clientId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteAppRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteAppResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteAppResult `json:"result"`
}

type DeleteAppResult struct {
    Count int `json:"count"`
}