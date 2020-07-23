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

type DeleteContactGroupCheckRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 组id (Optional) */
    GroupId *int `json:"groupId"`

    /* 用户pin (Optional) */
    Pin *string `json:"pin"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteContactGroupCheckRequest(
    regionId string,
) *DeleteContactGroupCheckRequest {

	return &DeleteContactGroupCheckRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/contact/group:delCheck",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param groupId: 组id (Optional)
 * param pin: 用户pin (Optional)
 */
func NewDeleteContactGroupCheckRequestWithAllParams(
    regionId string,
    groupId *int,
    pin *string,
) *DeleteContactGroupCheckRequest {

    return &DeleteContactGroupCheckRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/contact/group:delCheck",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        GroupId: groupId,
        Pin: pin,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteContactGroupCheckRequestWithoutParam() *DeleteContactGroupCheckRequest {

    return &DeleteContactGroupCheckRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/contact/group:delCheck",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DeleteContactGroupCheckRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param groupId: 组id(Optional) */
func (r *DeleteContactGroupCheckRequest) SetGroupId(groupId int) {
    r.GroupId = &groupId
}

/* param pin: 用户pin(Optional) */
func (r *DeleteContactGroupCheckRequest) SetPin(pin string) {
    r.Pin = &pin
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteContactGroupCheckRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteContactGroupCheckResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteContactGroupCheckResult `json:"result"`
}

type DeleteContactGroupCheckResult struct {
}