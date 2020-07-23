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
    csa "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/csa/models"
)

type QuerySecurityEventTrendRequest struct {

    core.JDCloudRequest

    /* 起始时间, 毫秒时间戳。如果不传，则需要传timeSpan (Optional) */
    TimeBegin *string `json:"timeBegin"`

    /* 结束时间, 毫秒时间戳, 不传为当前时间 (Optional) */
    TimeEnd *string `json:"timeEnd"`

    /* 天数，1:24时，7:7天，30:30天 (Optional) */
    TimeSpan *int `json:"timeSpan"`

    /* 趋势图时间间隔, 1h:一小时，1d:一天, 1m:一分钟，1s:一秒  */
    Interval string `json:"interval"`

    /* 告警事件等级,支持多个，逗号分隔  */
    Alert string `json:"alert"`

    /* 威胁事件等级,支持多个，逗号分隔  */
    Threat string `json:"threat"`

    /* 弱点事件等级,支持多个，逗号分隔  */
    Weakness string `json:"weakness"`
}

/*
 * param interval: 趋势图时间间隔, 1h:一小时，1d:一天, 1m:一分钟，1s:一秒 (Required)
 * param alert: 告警事件等级,支持多个，逗号分隔 (Required)
 * param threat: 威胁事件等级,支持多个，逗号分隔 (Required)
 * param weakness: 弱点事件等级,支持多个，逗号分隔 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQuerySecurityEventTrendRequest(
    interval string,
    alert string,
    threat string,
    weakness string,
) *QuerySecurityEventTrendRequest {

	return &QuerySecurityEventTrendRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/largeScreen:securityEventTrend",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Interval: interval,
        Alert: alert,
        Threat: threat,
        Weakness: weakness,
	}
}

/*
 * param timeBegin: 起始时间, 毫秒时间戳。如果不传，则需要传timeSpan (Optional)
 * param timeEnd: 结束时间, 毫秒时间戳, 不传为当前时间 (Optional)
 * param timeSpan: 天数，1:24时，7:7天，30:30天 (Optional)
 * param interval: 趋势图时间间隔, 1h:一小时，1d:一天, 1m:一分钟，1s:一秒 (Required)
 * param alert: 告警事件等级,支持多个，逗号分隔 (Required)
 * param threat: 威胁事件等级,支持多个，逗号分隔 (Required)
 * param weakness: 弱点事件等级,支持多个，逗号分隔 (Required)
 */
func NewQuerySecurityEventTrendRequestWithAllParams(
    timeBegin *string,
    timeEnd *string,
    timeSpan *int,
    interval string,
    alert string,
    threat string,
    weakness string,
) *QuerySecurityEventTrendRequest {

    return &QuerySecurityEventTrendRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/largeScreen:securityEventTrend",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        TimeBegin: timeBegin,
        TimeEnd: timeEnd,
        TimeSpan: timeSpan,
        Interval: interval,
        Alert: alert,
        Threat: threat,
        Weakness: weakness,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQuerySecurityEventTrendRequestWithoutParam() *QuerySecurityEventTrendRequest {

    return &QuerySecurityEventTrendRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/largeScreen:securityEventTrend",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param timeBegin: 起始时间, 毫秒时间戳。如果不传，则需要传timeSpan(Optional) */
func (r *QuerySecurityEventTrendRequest) SetTimeBegin(timeBegin string) {
    r.TimeBegin = &timeBegin
}

/* param timeEnd: 结束时间, 毫秒时间戳, 不传为当前时间(Optional) */
func (r *QuerySecurityEventTrendRequest) SetTimeEnd(timeEnd string) {
    r.TimeEnd = &timeEnd
}

/* param timeSpan: 天数，1:24时，7:7天，30:30天(Optional) */
func (r *QuerySecurityEventTrendRequest) SetTimeSpan(timeSpan int) {
    r.TimeSpan = &timeSpan
}

/* param interval: 趋势图时间间隔, 1h:一小时，1d:一天, 1m:一分钟，1s:一秒(Required) */
func (r *QuerySecurityEventTrendRequest) SetInterval(interval string) {
    r.Interval = interval
}

/* param alert: 告警事件等级,支持多个，逗号分隔(Required) */
func (r *QuerySecurityEventTrendRequest) SetAlert(alert string) {
    r.Alert = alert
}

/* param threat: 威胁事件等级,支持多个，逗号分隔(Required) */
func (r *QuerySecurityEventTrendRequest) SetThreat(threat string) {
    r.Threat = threat
}

/* param weakness: 弱点事件等级,支持多个，逗号分隔(Required) */
func (r *QuerySecurityEventTrendRequest) SetWeakness(weakness string) {
    r.Weakness = weakness
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QuerySecurityEventTrendRequest) GetRegionId() string {
    return ""
}

type QuerySecurityEventTrendResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QuerySecurityEventTrendResult `json:"result"`
}

type QuerySecurityEventTrendResult struct {
    TimeVector []string `json:"timeVector"`
    Detail []csa.NameCounts `json:"detail"`
    Code int `json:"code"`
    Message string `json:"message"`
}