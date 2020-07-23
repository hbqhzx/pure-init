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

type QueryLiveStatisticsDataRequest struct {

    core.JDCloudRequest

    /* 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z (Optional) */
    EndTime *string `json:"endTime"`

    /* 需要查询的域名, 必须为用户pin下有权限的域名 (Optional) */
    Domain *string `json:"domain"`

    /* app名 (Optional) */
    Appname *string `json:"appname"`

    /* 流名 (Optional) */
    Streamname *string `json:"streamname"`

    /* 子域名 (Optional) */
    SubDomain *string `json:"subDomain"`

    /* 需要查询的字段 (Optional) */
    Fields *string `json:"fields"`

    /*  (Optional) */
    Area *string `json:"area"`

    /*  (Optional) */
    Isp *string `json:"isp"`

    /*  (Optional) */
    ReqMethod *string `json:"reqMethod"`

    /* 查询的流协议类型 (Optional) */
    Scheme *string `json:"scheme"`

    /* 时间粒度，可选值:[oneMin,fiveMin,followTime],followTime只会返回一个汇总后的数据 (Optional) */
    Period *string `json:"period"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryLiveStatisticsDataRequest(
) *QueryLiveStatisticsDataRequest {

	return &QueryLiveStatisticsDataRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/liveStatistics",
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
 * param appname: app名 (Optional)
 * param streamname: 流名 (Optional)
 * param subDomain: 子域名 (Optional)
 * param fields: 需要查询的字段 (Optional)
 * param area:  (Optional)
 * param isp:  (Optional)
 * param reqMethod:  (Optional)
 * param scheme: 查询的流协议类型 (Optional)
 * param period: 时间粒度，可选值:[oneMin,fiveMin,followTime],followTime只会返回一个汇总后的数据 (Optional)
 */
func NewQueryLiveStatisticsDataRequestWithAllParams(
    startTime *string,
    endTime *string,
    domain *string,
    appname *string,
    streamname *string,
    subDomain *string,
    fields *string,
    area *string,
    isp *string,
    reqMethod *string,
    scheme *string,
    period *string,
) *QueryLiveStatisticsDataRequest {

    return &QueryLiveStatisticsDataRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveStatistics",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        StartTime: startTime,
        EndTime: endTime,
        Domain: domain,
        Appname: appname,
        Streamname: streamname,
        SubDomain: subDomain,
        Fields: fields,
        Area: area,
        Isp: isp,
        ReqMethod: reqMethod,
        Scheme: scheme,
        Period: period,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryLiveStatisticsDataRequestWithoutParam() *QueryLiveStatisticsDataRequest {

    return &QueryLiveStatisticsDataRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveStatistics",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param startTime: 查询起始时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryLiveStatisticsDataRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 查询截止时间,UTC时间，格式为:yyyy-MM-dd'T'HH:mm:ss'Z'，示例:2018-10-21T10:00:00Z(Optional) */
func (r *QueryLiveStatisticsDataRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param domain: 需要查询的域名, 必须为用户pin下有权限的域名(Optional) */
func (r *QueryLiveStatisticsDataRequest) SetDomain(domain string) {
    r.Domain = &domain
}

/* param appname: app名(Optional) */
func (r *QueryLiveStatisticsDataRequest) SetAppname(appname string) {
    r.Appname = &appname
}

/* param streamname: 流名(Optional) */
func (r *QueryLiveStatisticsDataRequest) SetStreamname(streamname string) {
    r.Streamname = &streamname
}

/* param subDomain: 子域名(Optional) */
func (r *QueryLiveStatisticsDataRequest) SetSubDomain(subDomain string) {
    r.SubDomain = &subDomain
}

/* param fields: 需要查询的字段(Optional) */
func (r *QueryLiveStatisticsDataRequest) SetFields(fields string) {
    r.Fields = &fields
}

/* param area: (Optional) */
func (r *QueryLiveStatisticsDataRequest) SetArea(area string) {
    r.Area = &area
}

/* param isp: (Optional) */
func (r *QueryLiveStatisticsDataRequest) SetIsp(isp string) {
    r.Isp = &isp
}

/* param reqMethod: (Optional) */
func (r *QueryLiveStatisticsDataRequest) SetReqMethod(reqMethod string) {
    r.ReqMethod = &reqMethod
}

/* param scheme: 查询的流协议类型(Optional) */
func (r *QueryLiveStatisticsDataRequest) SetScheme(scheme string) {
    r.Scheme = &scheme
}

/* param period: 时间粒度，可选值:[oneMin,fiveMin,followTime],followTime只会返回一个汇总后的数据(Optional) */
func (r *QueryLiveStatisticsDataRequest) SetPeriod(period string) {
    r.Period = &period
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryLiveStatisticsDataRequest) GetRegionId() string {
    return ""
}

type QueryLiveStatisticsDataResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryLiveStatisticsDataResult `json:"result"`
}

type QueryLiveStatisticsDataResult struct {
    StartTime string `json:"startTime"`
    EndTime string `json:"endTime"`
    Domain string `json:"domain"`
    Statistics []cdn.StatisticsDataItem `json:"statistics"`
}