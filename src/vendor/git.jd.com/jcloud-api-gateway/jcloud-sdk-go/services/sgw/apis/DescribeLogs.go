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
    sgw "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/sgw/models"
    common "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/common/models"
)

type DescribeLogsRequest struct {

    core.JDCloudRequest

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为10；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 请求的时间段,取值范围(1~3天)或(1-23小时) (Optional) */
    Timespan *int `json:"timespan"`

    /* 请求的时间单位,1表示小时,2表示天,默认单位是天 (Optional) */
    Timeunit *int `json:"timeunit"`

    /* 查询条件,支持根据 区域(dataCenter)、实例(wafId)、攻击类型(logType)、动作(action)、源IP(clientAddr) 精确匹配过滤 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLogsRequest(
) *DescribeLogsRequest {

	return &DescribeLogsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/logs",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为10；取值范围[10, 100] (Optional)
 * param timespan: 请求的时间段,取值范围(1~3天)或(1-23小时) (Optional)
 * param timeunit: 请求的时间单位,1表示小时,2表示天,默认单位是天 (Optional)
 * param filters: 查询条件,支持根据 区域(dataCenter)、实例(wafId)、攻击类型(logType)、动作(action)、源IP(clientAddr) 精确匹配过滤 (Optional)
 */
func NewDescribeLogsRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    timespan *int,
    timeunit *int,
    filters []common.Filter,
) *DescribeLogsRequest {

    return &DescribeLogsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/logs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
        Timespan: timespan,
        Timeunit: timeunit,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLogsRequestWithoutParam() *DescribeLogsRequest {

    return &DescribeLogsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/logs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeLogsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为10；取值范围[10, 100](Optional) */
func (r *DescribeLogsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param timespan: 请求的时间段,取值范围(1~3天)或(1-23小时)(Optional) */
func (r *DescribeLogsRequest) SetTimespan(timespan int) {
    r.Timespan = &timespan
}

/* param timeunit: 请求的时间单位,1表示小时,2表示天,默认单位是天(Optional) */
func (r *DescribeLogsRequest) SetTimeunit(timeunit int) {
    r.Timeunit = &timeunit
}

/* param filters: 查询条件,支持根据 区域(dataCenter)、实例(wafId)、攻击类型(logType)、动作(action)、源IP(clientAddr) 精确匹配过滤(Optional) */
func (r *DescribeLogsRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLogsRequest) GetRegionId() string {
    return ""
}

type DescribeLogsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLogsResult `json:"result"`
}

type DescribeLogsResult struct {
    Logs []sgw.Log `json:"logs"`
    Total int `json:"total"`
}