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
    monitor "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/monitor/models"
)

type DescribeNamespacesRequest struct {

    core.JDCloudRequest

    /* region  */
    RegionId string `json:"regionId"`
}

/*
 * param regionId: region (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeNamespacesRequest(
    regionId string,
) *DescribeNamespacesRequest {

	return &DescribeNamespacesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cmNameSpaces",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: region (Required)
 */
func NewDescribeNamespacesRequestWithAllParams(
    regionId string,
) *DescribeNamespacesRequest {

    return &DescribeNamespacesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cmNameSpaces",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeNamespacesRequestWithoutParam() *DescribeNamespacesRequest {

    return &DescribeNamespacesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cmNameSpaces",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: region(Required) */
func (r *DescribeNamespacesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeNamespacesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeNamespacesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeNamespacesResult `json:"result"`
}

type DescribeNamespacesResult struct {
    NamespaceList []monitor.NsInfo `json:"namespaceList"`
}