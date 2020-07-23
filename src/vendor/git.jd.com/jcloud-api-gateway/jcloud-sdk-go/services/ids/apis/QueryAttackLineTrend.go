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
    ids "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/ids/models"
)

type QueryAttackLineTrendRequest struct {

    core.JDCloudRequest

    /*  (Optional) */
    Ip *string `json:"ip"`

    /* 大于0的时间戳距离1970年1月1日,单位毫秒  */
    StartTime int `json:"startTime"`

    /* 大于0的时间戳距离1970年1月1日，大于startTime,单位毫秒  */
    EndTime int `json:"endTime"`

    /* 约定 攻击类型码0 是全部类型的攻击日志 (Optional) */
    TypeCode *int `json:"typeCode"`

    /* 时间间隔: 1h/1d/1w  */
    Interval string `json:"interval"`
}

/*
 * param startTime: 大于0的时间戳距离1970年1月1日,单位毫秒 (Required)
 * param endTime: 大于0的时间戳距离1970年1月1日，大于startTime,单位毫秒 (Required)
 * param interval: 时间间隔: 1h/1d/1w (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryAttackLineTrendRequest(
    startTime int,
    endTime int,
    interval string,
) *QueryAttackLineTrendRequest {

	return &QueryAttackLineTrendRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/attacks:queryLineTrendData",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        StartTime: startTime,
        EndTime: endTime,
        Interval: interval,
	}
}

/*
 * param ip:  (Optional)
 * param startTime: 大于0的时间戳距离1970年1月1日,单位毫秒 (Required)
 * param endTime: 大于0的时间戳距离1970年1月1日，大于startTime,单位毫秒 (Required)
 * param typeCode: 约定 攻击类型码0 是全部类型的攻击日志 (Optional)
 * param interval: 时间间隔: 1h/1d/1w (Required)
 */
func NewQueryAttackLineTrendRequestWithAllParams(
    ip *string,
    startTime int,
    endTime int,
    typeCode *int,
    interval string,
) *QueryAttackLineTrendRequest {

    return &QueryAttackLineTrendRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/attacks:queryLineTrendData",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Ip: ip,
        StartTime: startTime,
        EndTime: endTime,
        TypeCode: typeCode,
        Interval: interval,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryAttackLineTrendRequestWithoutParam() *QueryAttackLineTrendRequest {

    return &QueryAttackLineTrendRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/attacks:queryLineTrendData",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param ip: (Optional) */
func (r *QueryAttackLineTrendRequest) SetIp(ip string) {
    r.Ip = &ip
}

/* param startTime: 大于0的时间戳距离1970年1月1日,单位毫秒(Required) */
func (r *QueryAttackLineTrendRequest) SetStartTime(startTime int) {
    r.StartTime = startTime
}

/* param endTime: 大于0的时间戳距离1970年1月1日，大于startTime,单位毫秒(Required) */
func (r *QueryAttackLineTrendRequest) SetEndTime(endTime int) {
    r.EndTime = endTime
}

/* param typeCode: 约定 攻击类型码0 是全部类型的攻击日志(Optional) */
func (r *QueryAttackLineTrendRequest) SetTypeCode(typeCode int) {
    r.TypeCode = &typeCode
}

/* param interval: 时间间隔: 1h/1d/1w(Required) */
func (r *QueryAttackLineTrendRequest) SetInterval(interval string) {
    r.Interval = interval
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryAttackLineTrendRequest) GetRegionId() string {
    return ""
}

type QueryAttackLineTrendResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryAttackLineTrendResult `json:"result"`
}

type QueryAttackLineTrendResult struct {
    LineTrendData []ids.EntryCount `json:"lineTrendData"`
}