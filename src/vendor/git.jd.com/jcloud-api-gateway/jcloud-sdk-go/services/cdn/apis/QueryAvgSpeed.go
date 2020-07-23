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

type QueryAvgSpeedRequest struct {

    core.JDCloudRequest

    /* 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    EndTime *string `json:"endTime"`

    /* 需要查询的域名, 必须为用户pin下有权限的域名 (Optional) */
    Domain *string `json:"domain"`

    /*  (Optional) */
    Area *string `json:"area"`

    /*  (Optional) */
    Isp *string `json:"isp"`

    /* 排序字段默认avgbandwidth (Optional) */
    SortField *string `json:"sortField"`

    /* 排序规则,默认desc (Optional) */
    SortRule *string `json:"sortRule"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryAvgSpeedRequest(
) *QueryAvgSpeedRequest {

	return &QueryAvgSpeedRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/console:avgspeed",
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
 * param area:  (Optional)
 * param isp:  (Optional)
 * param sortField: 排序字段默认avgbandwidth (Optional)
 * param sortRule: 排序规则,默认desc (Optional)
 */
func NewQueryAvgSpeedRequestWithAllParams(
    startTime *string,
    endTime *string,
    domain *string,
    area *string,
    isp *string,
    sortField *string,
    sortRule *string,
) *QueryAvgSpeedRequest {

    return &QueryAvgSpeedRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/console:avgspeed",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        StartTime: startTime,
        EndTime: endTime,
        Domain: domain,
        Area: area,
        Isp: isp,
        SortField: sortField,
        SortRule: sortRule,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryAvgSpeedRequestWithoutParam() *QueryAvgSpeedRequest {

    return &QueryAvgSpeedRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/console:avgspeed",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryAvgSpeedRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryAvgSpeedRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param domain: 需要查询的域名, 必须为用户pin下有权限的域名(Optional) */
func (r *QueryAvgSpeedRequest) SetDomain(domain string) {
    r.Domain = &domain
}

/* param area: (Optional) */
func (r *QueryAvgSpeedRequest) SetArea(area string) {
    r.Area = &area
}

/* param isp: (Optional) */
func (r *QueryAvgSpeedRequest) SetIsp(isp string) {
    r.Isp = &isp
}

/* param sortField: 排序字段默认avgbandwidth(Optional) */
func (r *QueryAvgSpeedRequest) SetSortField(sortField string) {
    r.SortField = &sortField
}

/* param sortRule: 排序规则,默认desc(Optional) */
func (r *QueryAvgSpeedRequest) SetSortRule(sortRule string) {
    r.SortRule = &sortRule
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryAvgSpeedRequest) GetRegionId() string {
    return ""
}

type QueryAvgSpeedResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryAvgSpeedResult `json:"result"`
}

type QueryAvgSpeedResult struct {
    Total int `json:"total"`
    Data []cdn.AvgspeedItem `json:"data"`
}