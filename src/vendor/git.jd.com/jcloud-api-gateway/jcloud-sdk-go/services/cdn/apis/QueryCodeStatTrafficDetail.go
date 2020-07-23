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

type QueryCodeStatTrafficDetailRequest struct {

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

    /* 时间粒度，可选值:[oneMin,fiveMin,followTime],followTime只会返回一个汇总后的数据 (Optional) */
    Period *string `json:"period"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryCodeStatTrafficDetailRequest(
) *QueryCodeStatTrafficDetailRequest {

	return &QueryCodeStatTrafficDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/console:codestatDetail",
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
 * param period: 时间粒度，可选值:[oneMin,fiveMin,followTime],followTime只会返回一个汇总后的数据 (Optional)
 */
func NewQueryCodeStatTrafficDetailRequestWithAllParams(
    startTime *string,
    endTime *string,
    domain *string,
    area *string,
    isp *string,
    period *string,
) *QueryCodeStatTrafficDetailRequest {

    return &QueryCodeStatTrafficDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/console:codestatDetail",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        StartTime: startTime,
        EndTime: endTime,
        Domain: domain,
        Area: area,
        Isp: isp,
        Period: period,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryCodeStatTrafficDetailRequestWithoutParam() *QueryCodeStatTrafficDetailRequest {

    return &QueryCodeStatTrafficDetailRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/console:codestatDetail",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryCodeStatTrafficDetailRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryCodeStatTrafficDetailRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param domain: 需要查询的域名, 必须为用户pin下有权限的域名(Optional) */
func (r *QueryCodeStatTrafficDetailRequest) SetDomain(domain string) {
    r.Domain = &domain
}

/* param area: (Optional) */
func (r *QueryCodeStatTrafficDetailRequest) SetArea(area string) {
    r.Area = &area
}

/* param isp: (Optional) */
func (r *QueryCodeStatTrafficDetailRequest) SetIsp(isp string) {
    r.Isp = &isp
}

/* param period: 时间粒度，可选值:[oneMin,fiveMin,followTime],followTime只会返回一个汇总后的数据(Optional) */
func (r *QueryCodeStatTrafficDetailRequest) SetPeriod(period string) {
    r.Period = &period
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryCodeStatTrafficDetailRequest) GetRegionId() string {
    return ""
}

type QueryCodeStatTrafficDetailResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryCodeStatTrafficDetailResult `json:"result"`
}

type QueryCodeStatTrafficDetailResult struct {
    Details []cdn.CodeDetailItem `json:"details"`
}