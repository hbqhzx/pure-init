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

type DescribeWafTrendRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 请求的时间段,取值范围(1~30天)或(1-23小时) (Optional) */
    Timespan *int `json:"timespan"`

    /* 请求的时间单位,1表示小时,2表示天,默认单位是天 (Optional) */
    Timeunit *int `json:"timeunit"`
}

/*
 * param regionId: 地域 Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeWafTrendRequest(
    regionId string,
) *DescribeWafTrendRequest {

	return &DescribeWafTrendRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/statistics:wafTrend",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param timespan: 请求的时间段,取值范围(1~30天)或(1-23小时) (Optional)
 * param timeunit: 请求的时间单位,1表示小时,2表示天,默认单位是天 (Optional)
 */
func NewDescribeWafTrendRequestWithAllParams(
    regionId string,
    timespan *int,
    timeunit *int,
) *DescribeWafTrendRequest {

    return &DescribeWafTrendRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/statistics:wafTrend",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Timespan: timespan,
        Timeunit: timeunit,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeWafTrendRequestWithoutParam() *DescribeWafTrendRequest {

    return &DescribeWafTrendRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/statistics:wafTrend",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DescribeWafTrendRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param timespan: 请求的时间段,取值范围(1~30天)或(1-23小时)(Optional) */
func (r *DescribeWafTrendRequest) SetTimespan(timespan int) {
    r.Timespan = &timespan
}

/* param timeunit: 请求的时间单位,1表示小时,2表示天,默认单位是天(Optional) */
func (r *DescribeWafTrendRequest) SetTimeunit(timeunit int) {
    r.Timeunit = &timeunit
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeWafTrendRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeWafTrendResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeWafTrendResult `json:"result"`
}

type DescribeWafTrendResult struct {
    Data []interface{} `json:"data"`
}