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

type DeleteContainerRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Container ID  */
    ContainerId string `json:"containerId"`
}

/*
 * param regionId: Region ID (Required)
 * param containerId: Container ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteContainerRequest(
    regionId string,
    containerId string,
) *DeleteContainerRequest {

	return &DeleteContainerRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/containers/{containerId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ContainerId: containerId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param containerId: Container ID (Required)
 */
func NewDeleteContainerRequestWithAllParams(
    regionId string,
    containerId string,
) *DeleteContainerRequest {

    return &DeleteContainerRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/containers/{containerId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ContainerId: containerId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteContainerRequestWithoutParam() *DeleteContainerRequest {

    return &DeleteContainerRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/containers/{containerId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DeleteContainerRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param containerId: Container ID(Required) */
func (r *DeleteContainerRequest) SetContainerId(containerId string) {
    r.ContainerId = containerId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteContainerRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteContainerResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteContainerResult `json:"result"`
}

type DeleteContainerResult struct {
}