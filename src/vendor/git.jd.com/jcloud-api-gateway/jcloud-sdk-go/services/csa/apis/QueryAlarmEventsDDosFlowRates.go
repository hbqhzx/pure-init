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

type QueryAlarmEventsDDosFlowRatesRequest struct {

    core.JDCloudRequest

    /* alarmEvent事件ID  */
    AlarmEventId string `json:"alarmEventId"`

    /* IP地址  */
    Ip int `json:"ip"`
}

/*
 * param alarmEventId: alarmEvent事件ID (Required)
 * param ip: IP地址 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryAlarmEventsDDosFlowRatesRequest(
    alarmEventId string,
    ip int,
) *QueryAlarmEventsDDosFlowRatesRequest {

	return &QueryAlarmEventsDDosFlowRatesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/alarmEvents/{alarmEventId}:queryDDosFlowRates",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        AlarmEventId: alarmEventId,
        Ip: ip,
	}
}

/*
 * param alarmEventId: alarmEvent事件ID (Required)
 * param ip: IP地址 (Required)
 */
func NewQueryAlarmEventsDDosFlowRatesRequestWithAllParams(
    alarmEventId string,
    ip int,
) *QueryAlarmEventsDDosFlowRatesRequest {

    return &QueryAlarmEventsDDosFlowRatesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/alarmEvents/{alarmEventId}:queryDDosFlowRates",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        AlarmEventId: alarmEventId,
        Ip: ip,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryAlarmEventsDDosFlowRatesRequestWithoutParam() *QueryAlarmEventsDDosFlowRatesRequest {

    return &QueryAlarmEventsDDosFlowRatesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/alarmEvents/{alarmEventId}:queryDDosFlowRates",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param alarmEventId: alarmEvent事件ID(Required) */
func (r *QueryAlarmEventsDDosFlowRatesRequest) SetAlarmEventId(alarmEventId string) {
    r.AlarmEventId = alarmEventId
}

/* param ip: IP地址(Required) */
func (r *QueryAlarmEventsDDosFlowRatesRequest) SetIp(ip int) {
    r.Ip = ip
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryAlarmEventsDDosFlowRatesRequest) GetRegionId() string {
    return ""
}

type QueryAlarmEventsDDosFlowRatesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryAlarmEventsDDosFlowRatesResult `json:"result"`
}

type QueryAlarmEventsDDosFlowRatesResult struct {
    DdosFlowRates csa.DdosFlowRate `json:"ddosFlowRates"`
    Threshold csa.ThresholdObject `json:"threshold"`
    TimeRanges []csa.TimeRange `json:"timeRanges"`
}