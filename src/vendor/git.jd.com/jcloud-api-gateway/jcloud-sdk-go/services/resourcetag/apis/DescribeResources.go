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
    resourcetag "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/resourcetag/models"
)

type DescribeResourcesRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 资源参数  */
    ResourceVo *resourcetag.ResourceReqVo `json:"resourceVo"`
}

/*
 * param regionId: Region ID (Required)
 * param resourceVo: 资源参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeResourcesRequest(
    regionId string,
    resourceVo *resourcetag.ResourceReqVo,
) *DescribeResourcesRequest {

	return &DescribeResourcesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/tags:describeResources",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ResourceVo: resourceVo,
	}
}

/*
 * param regionId: Region ID (Required)
 * param resourceVo: 资源参数 (Required)
 */
func NewDescribeResourcesRequestWithAllParams(
    regionId string,
    resourceVo *resourcetag.ResourceReqVo,
) *DescribeResourcesRequest {

    return &DescribeResourcesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/tags:describeResources",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ResourceVo: resourceVo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeResourcesRequestWithoutParam() *DescribeResourcesRequest {

    return &DescribeResourcesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/tags:describeResources",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeResourcesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param resourceVo: 资源参数(Required) */
func (r *DescribeResourcesRequest) SetResourceVo(resourceVo *resourcetag.ResourceReqVo) {
    r.ResourceVo = resourceVo
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeResourcesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeResourcesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeResourcesResult `json:"result"`
}

type DescribeResourcesResult struct {
    Data resourcetag.ResourceResVo `json:"data"`
}