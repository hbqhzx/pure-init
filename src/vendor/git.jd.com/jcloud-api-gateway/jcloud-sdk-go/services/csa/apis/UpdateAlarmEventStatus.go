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

type UpdateAlarmEventStatusRequest struct {

    core.JDCloudRequest

    /* alarmEvent事件ID  */
    AlarmEventId string `json:"alarmEventId"`

    /* 事件状态(1:忽略;2:误报;3:确认)  */
    Status int `json:"status"`
}

/*
 * param alarmEventId: alarmEvent事件ID (Required)
 * param status: 事件状态(1:忽略;2:误报;3:确认) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateAlarmEventStatusRequest(
    alarmEventId string,
    status int,
) *UpdateAlarmEventStatusRequest {

	return &UpdateAlarmEventStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/alarmEvents/{alarmEventId}:updateStatus",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        AlarmEventId: alarmEventId,
        Status: status,
	}
}

/*
 * param alarmEventId: alarmEvent事件ID (Required)
 * param status: 事件状态(1:忽略;2:误报;3:确认) (Required)
 */
func NewUpdateAlarmEventStatusRequestWithAllParams(
    alarmEventId string,
    status int,
) *UpdateAlarmEventStatusRequest {

    return &UpdateAlarmEventStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/alarmEvents/{alarmEventId}:updateStatus",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        AlarmEventId: alarmEventId,
        Status: status,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateAlarmEventStatusRequestWithoutParam() *UpdateAlarmEventStatusRequest {

    return &UpdateAlarmEventStatusRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/alarmEvents/{alarmEventId}:updateStatus",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param alarmEventId: alarmEvent事件ID(Required) */
func (r *UpdateAlarmEventStatusRequest) SetAlarmEventId(alarmEventId string) {
    r.AlarmEventId = alarmEventId
}

/* param status: 事件状态(1:忽略;2:误报;3:确认)(Required) */
func (r *UpdateAlarmEventStatusRequest) SetStatus(status int) {
    r.Status = status
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateAlarmEventStatusRequest) GetRegionId() string {
    return ""
}

type UpdateAlarmEventStatusResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateAlarmEventStatusResult `json:"result"`
}

type UpdateAlarmEventStatusResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}