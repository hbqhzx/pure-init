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

type QueryAlarmEventsFileSandboxAnalysisRequest struct {

    core.JDCloudRequest

    /* 事件ID  */
    EventId string `json:"eventId"`
}

/*
 * param eventId: 事件ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryAlarmEventsFileSandboxAnalysisRequest(
    eventId string,
) *QueryAlarmEventsFileSandboxAnalysisRequest {

	return &QueryAlarmEventsFileSandboxAnalysisRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/alarmEvents:queryFileSandboxAnalysis",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        EventId: eventId,
	}
}

/*
 * param eventId: 事件ID (Required)
 */
func NewQueryAlarmEventsFileSandboxAnalysisRequestWithAllParams(
    eventId string,
) *QueryAlarmEventsFileSandboxAnalysisRequest {

    return &QueryAlarmEventsFileSandboxAnalysisRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/alarmEvents:queryFileSandboxAnalysis",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        EventId: eventId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryAlarmEventsFileSandboxAnalysisRequestWithoutParam() *QueryAlarmEventsFileSandboxAnalysisRequest {

    return &QueryAlarmEventsFileSandboxAnalysisRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/alarmEvents:queryFileSandboxAnalysis",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param eventId: 事件ID(Required) */
func (r *QueryAlarmEventsFileSandboxAnalysisRequest) SetEventId(eventId string) {
    r.EventId = eventId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryAlarmEventsFileSandboxAnalysisRequest) GetRegionId() string {
    return ""
}

type QueryAlarmEventsFileSandboxAnalysisResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryAlarmEventsFileSandboxAnalysisResult `json:"result"`
}

type QueryAlarmEventsFileSandboxAnalysisResult struct {
    FileSandboxAnalysis csa.FileSandboxAnalysis `json:"fileSandboxAnalysis"`
    Alert csa.SingleAttackAlert `json:"alert"`
}