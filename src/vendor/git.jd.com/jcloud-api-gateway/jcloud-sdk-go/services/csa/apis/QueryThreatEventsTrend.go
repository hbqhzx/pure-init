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

type QueryThreatEventsTrendRequest struct {

    core.JDCloudRequest

    /* 起始时间, 毫秒时间戳。如果不传，则需要传timeSpan (Optional) */
    TimeBegin *string `json:"timeBegin"`

    /* 结束时间, 毫秒时间戳, 不传为当前时间 (Optional) */
    TimeEnd *string `json:"timeEnd"`

    /* 天数，如果timeBegin，则该参数无效 (Optional) */
    TimeSpan *int `json:"timeSpan"`

    /* 事件等级,支持多个，逗号分隔 (Optional) */
    Severities *string `json:"severities"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryThreatEventsTrendRequest(
) *QueryThreatEventsTrendRequest {

	return &QueryThreatEventsTrendRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/threatEvents:queryTrend",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param timeBegin: 起始时间, 毫秒时间戳。如果不传，则需要传timeSpan (Optional)
 * param timeEnd: 结束时间, 毫秒时间戳, 不传为当前时间 (Optional)
 * param timeSpan: 天数，如果timeBegin，则该参数无效 (Optional)
 * param severities: 事件等级,支持多个，逗号分隔 (Optional)
 */
func NewQueryThreatEventsTrendRequestWithAllParams(
    timeBegin *string,
    timeEnd *string,
    timeSpan *int,
    severities *string,
) *QueryThreatEventsTrendRequest {

    return &QueryThreatEventsTrendRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/threatEvents:queryTrend",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        TimeBegin: timeBegin,
        TimeEnd: timeEnd,
        TimeSpan: timeSpan,
        Severities: severities,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryThreatEventsTrendRequestWithoutParam() *QueryThreatEventsTrendRequest {

    return &QueryThreatEventsTrendRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/threatEvents:queryTrend",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param timeBegin: 起始时间, 毫秒时间戳。如果不传，则需要传timeSpan(Optional) */
func (r *QueryThreatEventsTrendRequest) SetTimeBegin(timeBegin string) {
    r.TimeBegin = &timeBegin
}

/* param timeEnd: 结束时间, 毫秒时间戳, 不传为当前时间(Optional) */
func (r *QueryThreatEventsTrendRequest) SetTimeEnd(timeEnd string) {
    r.TimeEnd = &timeEnd
}

/* param timeSpan: 天数，如果timeBegin，则该参数无效(Optional) */
func (r *QueryThreatEventsTrendRequest) SetTimeSpan(timeSpan int) {
    r.TimeSpan = &timeSpan
}

/* param severities: 事件等级,支持多个，逗号分隔(Optional) */
func (r *QueryThreatEventsTrendRequest) SetSeverities(severities string) {
    r.Severities = &severities
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryThreatEventsTrendRequest) GetRegionId() string {
    return ""
}

type QueryThreatEventsTrendResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryThreatEventsTrendResult `json:"result"`
}

type QueryThreatEventsTrendResult struct {
    Trend string `json:"trend"`
    Percentage int `json:"percentage"`
    Message string `json:"message"`
}