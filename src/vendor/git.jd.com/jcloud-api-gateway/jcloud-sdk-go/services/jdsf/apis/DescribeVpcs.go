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
    jdsf "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/jdsf/models"
)

type DescribeVpcsRequest struct {

    core.JDCloudRequest

    /* 可用区id  */
    RegionId string `json:"regionId"`
}

/*
 * param regionId: 可用区id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeVpcsRequest(
    regionId string,
) *DescribeVpcsRequest {

	return &DescribeVpcsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/vpcs",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 可用区id (Required)
 */
func NewDescribeVpcsRequestWithAllParams(
    regionId string,
) *DescribeVpcsRequest {

    return &DescribeVpcsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpcs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeVpcsRequestWithoutParam() *DescribeVpcsRequest {

    return &DescribeVpcsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpcs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 可用区id(Required) */
func (r *DescribeVpcsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeVpcsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeVpcsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeVpcsResult `json:"result"`
}

type DescribeVpcsResult struct {
    Vpcs []jdsf.VpcInfo `json:"vpcs"`
}