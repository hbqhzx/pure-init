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
    ark "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/ark/models"
)

type DescribeServicesRequest struct {

    core.JDCloudRequest

    /* 区域ID (Optional) */
    RegionId *string `json:"regionId"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeServicesRequest(
) *DescribeServicesRequest {

	return &DescribeServicesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/services",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param regionId: 区域ID (Optional)
 */
func NewDescribeServicesRequestWithAllParams(
    regionId *string,
) *DescribeServicesRequest {

    return &DescribeServicesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/services",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeServicesRequestWithoutParam() *DescribeServicesRequest {

    return &DescribeServicesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/services",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域ID(Optional) */
func (r *DescribeServicesRequest) SetRegionId(regionId string) {
    r.RegionId = &regionId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeServicesRequest) GetRegionId() string {
    return ""
}

type DescribeServicesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeServicesResult `json:"result"`
}

type DescribeServicesResult struct {
    Services []ark.Service `json:"services"`
}