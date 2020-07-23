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
    common "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/common/models"
)

type DescribeWafRunningStatusRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 查询条件,支持 wafIds(Waf实例id,数组) (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域 Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeWafRunningStatusRequest(
    regionId string,
) *DescribeWafRunningStatusRequest {

	return &DescribeWafRunningStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/statistics:wafRunningStatus",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param filters: 查询条件,支持 wafIds(Waf实例id,数组) (Optional)
 */
func NewDescribeWafRunningStatusRequestWithAllParams(
    regionId string,
    filters []common.Filter,
) *DescribeWafRunningStatusRequest {

    return &DescribeWafRunningStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/statistics:wafRunningStatus",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeWafRunningStatusRequestWithoutParam() *DescribeWafRunningStatusRequest {

    return &DescribeWafRunningStatusRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/statistics:wafRunningStatus",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DescribeWafRunningStatusRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param filters: 查询条件,支持 wafIds(Waf实例id,数组)(Optional) */
func (r *DescribeWafRunningStatusRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeWafRunningStatusRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeWafRunningStatusResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeWafRunningStatusResult `json:"result"`
}

type DescribeWafRunningStatusResult struct {
    Running int `json:"running"`
    Total int `json:"total"`
}