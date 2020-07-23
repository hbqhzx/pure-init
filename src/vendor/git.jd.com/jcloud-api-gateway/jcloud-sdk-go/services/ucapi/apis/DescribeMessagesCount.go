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

type DescribeMessagesCountRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 用户pin (Optional) */
    Pin *string `json:"pin"`

    /* 消息类别id (Optional) */
    Id *int `json:"id"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeMessagesCountRequest(
    regionId string,
) *DescribeMessagesCountRequest {

	return &DescribeMessagesCountRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/messages/count",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pin: 用户pin (Optional)
 * param id: 消息类别id (Optional)
 */
func NewDescribeMessagesCountRequestWithAllParams(
    regionId string,
    pin *string,
    id *int,
) *DescribeMessagesCountRequest {

    return &DescribeMessagesCountRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/messages/count",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Pin: pin,
        Id: id,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeMessagesCountRequestWithoutParam() *DescribeMessagesCountRequest {

    return &DescribeMessagesCountRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/messages/count",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeMessagesCountRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pin: 用户pin(Optional) */
func (r *DescribeMessagesCountRequest) SetPin(pin string) {
    r.Pin = &pin
}

/* param id: 消息类别id(Optional) */
func (r *DescribeMessagesCountRequest) SetId(id int) {
    r.Id = &id
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeMessagesCountRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeMessagesCountResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeMessagesCountResult `json:"result"`
}

type DescribeMessagesCountResult struct {
    Total int `json:"total"`
}