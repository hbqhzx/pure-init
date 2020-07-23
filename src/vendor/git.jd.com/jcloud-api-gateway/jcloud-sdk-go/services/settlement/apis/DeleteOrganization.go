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

type DeleteOrganizationRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* id  */
    Id int `json:"id"`
}

/*
 * param regionId:  (Required)
 * param id: id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteOrganizationRequest(
    regionId string,
    id int,
) *DeleteOrganizationRequest {

	return &DeleteOrganizationRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/organization/{id}:delete",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Id: id,
	}
}

/*
 * param regionId:  (Required)
 * param id: id (Required)
 */
func NewDeleteOrganizationRequestWithAllParams(
    regionId string,
    id int,
) *DeleteOrganizationRequest {

    return &DeleteOrganizationRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/organization/{id}:delete",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Id: id,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteOrganizationRequestWithoutParam() *DeleteOrganizationRequest {

    return &DeleteOrganizationRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/organization/{id}:delete",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *DeleteOrganizationRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param id: id(Required) */
func (r *DeleteOrganizationRequest) SetId(id int) {
    r.Id = id
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteOrganizationRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteOrganizationResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteOrganizationResult `json:"result"`
}

type DeleteOrganizationResult struct {
    Result bool `json:"result"`
}