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

type QueryAlarmEventRequest struct {

    core.JDCloudRequest

    /* alarmEvent事件ID  */
    AlarmEventId string `json:"alarmEventId"`
}

/*
 * param alarmEventId: alarmEvent事件ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryAlarmEventRequest(
    alarmEventId string,
) *QueryAlarmEventRequest {

	return &QueryAlarmEventRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/alarmEvents/{alarmEventId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        AlarmEventId: alarmEventId,
	}
}

/*
 * param alarmEventId: alarmEvent事件ID (Required)
 */
func NewQueryAlarmEventRequestWithAllParams(
    alarmEventId string,
) *QueryAlarmEventRequest {

    return &QueryAlarmEventRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/alarmEvents/{alarmEventId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        AlarmEventId: alarmEventId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryAlarmEventRequestWithoutParam() *QueryAlarmEventRequest {

    return &QueryAlarmEventRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/alarmEvents/{alarmEventId}",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param alarmEventId: alarmEvent事件ID(Required) */
func (r *QueryAlarmEventRequest) SetAlarmEventId(alarmEventId string) {
    r.AlarmEventId = alarmEventId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryAlarmEventRequest) GetRegionId() string {
    return ""
}

type QueryAlarmEventResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryAlarmEventResult `json:"result"`
}

type QueryAlarmEventResult struct {
    AlarmEventDetail csa.AlarmEventDetail `json:"alarmEventDetail"`
}