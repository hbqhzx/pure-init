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
    cdn "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/cdn/models"
)

type QueryStreambdwTableRequest struct {

    core.JDCloudRequest

    /* 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    EndTime *string `json:"endTime"`

    /* 需要查询的域名, 必须为用户pin下有权限的域名 (Optional) */
    Domain *string `json:"domain"`

    /* app名 (Optional) */
    AppName *string `json:"appName"`

    /* 流名 (Optional) */
    StreamName *string `json:"streamName"`

    /*  (Optional) */
    PageNumber *int `json:"pageNumber"`

    /*  (Optional) */
    PageSize *int `json:"pageSize"`

    /* 排序字段 (Optional) */
    SortField *string `json:"sortField"`

    /* 排序规则 (Optional) */
    SortRule *string `json:"sortRule"`

    /*  (Optional) */
    Area *string `json:"area"`

    /*  (Optional) */
    Isp *string `json:"isp"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryStreambdwTableRequest(
) *QueryStreambdwTableRequest {

	return &QueryStreambdwTableRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/console:streambdwTable",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 * param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional)
 * param domain: 需要查询的域名, 必须为用户pin下有权限的域名 (Optional)
 * param appName: app名 (Optional)
 * param streamName: 流名 (Optional)
 * param pageNumber:  (Optional)
 * param pageSize:  (Optional)
 * param sortField: 排序字段 (Optional)
 * param sortRule: 排序规则 (Optional)
 * param area:  (Optional)
 * param isp:  (Optional)
 */
func NewQueryStreambdwTableRequestWithAllParams(
    startTime *string,
    endTime *string,
    domain *string,
    appName *string,
    streamName *string,
    pageNumber *int,
    pageSize *int,
    sortField *string,
    sortRule *string,
    area *string,
    isp *string,
) *QueryStreambdwTableRequest {

    return &QueryStreambdwTableRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/console:streambdwTable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        StartTime: startTime,
        EndTime: endTime,
        Domain: domain,
        AppName: appName,
        StreamName: streamName,
        PageNumber: pageNumber,
        PageSize: pageSize,
        SortField: sortField,
        SortRule: sortRule,
        Area: area,
        Isp: isp,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryStreambdwTableRequestWithoutParam() *QueryStreambdwTableRequest {

    return &QueryStreambdwTableRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/console:streambdwTable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryStreambdwTableRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryStreambdwTableRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param domain: 需要查询的域名, 必须为用户pin下有权限的域名(Optional) */
func (r *QueryStreambdwTableRequest) SetDomain(domain string) {
    r.Domain = &domain
}

/* param appName: app名(Optional) */
func (r *QueryStreambdwTableRequest) SetAppName(appName string) {
    r.AppName = &appName
}

/* param streamName: 流名(Optional) */
func (r *QueryStreambdwTableRequest) SetStreamName(streamName string) {
    r.StreamName = &streamName
}

/* param pageNumber: (Optional) */
func (r *QueryStreambdwTableRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: (Optional) */
func (r *QueryStreambdwTableRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param sortField: 排序字段(Optional) */
func (r *QueryStreambdwTableRequest) SetSortField(sortField string) {
    r.SortField = &sortField
}

/* param sortRule: 排序规则(Optional) */
func (r *QueryStreambdwTableRequest) SetSortRule(sortRule string) {
    r.SortRule = &sortRule
}

/* param area: (Optional) */
func (r *QueryStreambdwTableRequest) SetArea(area string) {
    r.Area = &area
}

/* param isp: (Optional) */
func (r *QueryStreambdwTableRequest) SetIsp(isp string) {
    r.Isp = &isp
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryStreambdwTableRequest) GetRegionId() string {
    return ""
}

type QueryStreambdwTableResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryStreambdwTableResult `json:"result"`
}

type QueryStreambdwTableResult struct {
    Data []cdn.StreamTableItem `json:"data"`
}