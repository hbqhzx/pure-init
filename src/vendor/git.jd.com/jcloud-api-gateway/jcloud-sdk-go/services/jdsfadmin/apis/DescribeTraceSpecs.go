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
    jdsfadmin "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/jdsfadmin/models"
)

type DescribeTraceSpecsRequest struct {

    core.JDCloudRequest

    /* 区域id  */
    RegionId string `json:"regionId"`
}

/*
 * param regionId: 区域id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeTraceSpecsRequest(
    regionId string,
) *DescribeTraceSpecsRequest {

	return &DescribeTraceSpecsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/trace_specs",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 区域id (Required)
 */
func NewDescribeTraceSpecsRequestWithAllParams(
    regionId string,
) *DescribeTraceSpecsRequest {

    return &DescribeTraceSpecsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/trace_specs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeTraceSpecsRequestWithoutParam() *DescribeTraceSpecsRequest {

    return &DescribeTraceSpecsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/trace_specs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域id(Required) */
func (r *DescribeTraceSpecsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeTraceSpecsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeTraceSpecsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeTraceSpecsResult `json:"result"`
}

type DescribeTraceSpecsResult struct {
    InstanceSpecs []jdsfadmin.InstanceSpecification `json:"instanceSpecs"`
}