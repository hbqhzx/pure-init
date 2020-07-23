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
    bgw "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/bgw/models"
)

type DescribeConnectionRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Connection ID  */
    ConnectionId string `json:"connectionId"`
}

/*
 * param regionId: Region ID (Required)
 * param connectionId: Connection ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeConnectionRequest(
    regionId string,
    connectionId string,
) *DescribeConnectionRequest {

	return &DescribeConnectionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/connections/{connectionId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ConnectionId: connectionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param connectionId: Connection ID (Required)
 */
func NewDescribeConnectionRequestWithAllParams(
    regionId string,
    connectionId string,
) *DescribeConnectionRequest {

    return &DescribeConnectionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/connections/{connectionId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ConnectionId: connectionId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeConnectionRequestWithoutParam() *DescribeConnectionRequest {

    return &DescribeConnectionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/connections/{connectionId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeConnectionRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param connectionId: Connection ID(Required) */
func (r *DescribeConnectionRequest) SetConnectionId(connectionId string) {
    r.ConnectionId = connectionId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeConnectionRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeConnectionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeConnectionResult `json:"result"`
}

type DescribeConnectionResult struct {
    Connection bgw.Connection `json:"connection"`
}