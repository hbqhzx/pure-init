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

type StopSlbRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* LB ID  */
    Id string `json:"id"`
}

/*
 * param regionId: 地域ID (Required)
 * param id: LB ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewStopSlbRequest(
    regionId string,
    id string,
) *StopSlbRequest {

	return &StopSlbRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/vpc_slbs/{id}:stop",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Id: id,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param id: LB ID (Required)
 */
func NewStopSlbRequestWithAllParams(
    regionId string,
    id string,
) *StopSlbRequest {

    return &StopSlbRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpc_slbs/{id}:stop",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Id: id,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewStopSlbRequestWithoutParam() *StopSlbRequest {

    return &StopSlbRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpc_slbs/{id}:stop",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *StopSlbRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param id: LB ID(Required) */
func (r *StopSlbRequest) SetId(id string) {
    r.Id = id
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r StopSlbRequest) GetRegionId() string {
    return r.RegionId
}

type StopSlbResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result StopSlbResult `json:"result"`
}

type StopSlbResult struct {
}