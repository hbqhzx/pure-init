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

type BatchUpdateThreatEventsStatusRequest struct {

    core.JDCloudRequest

    /* 事件IDs  */
    ThreatEventIds []string `json:"threatEventIds"`

    /* 事件状态(1:忽略;2:误报;3:确认)  */
    Status int `json:"status"`
}

/*
 * param threatEventIds: 事件IDs (Required)
 * param status: 事件状态(1:忽略;2:误报;3:确认) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewBatchUpdateThreatEventsStatusRequest(
    threatEventIds []string,
    status int,
) *BatchUpdateThreatEventsStatusRequest {

	return &BatchUpdateThreatEventsStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/threatEvents:batchUpdateStatus",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        ThreatEventIds: threatEventIds,
        Status: status,
	}
}

/*
 * param threatEventIds: 事件IDs (Required)
 * param status: 事件状态(1:忽略;2:误报;3:确认) (Required)
 */
func NewBatchUpdateThreatEventsStatusRequestWithAllParams(
    threatEventIds []string,
    status int,
) *BatchUpdateThreatEventsStatusRequest {

    return &BatchUpdateThreatEventsStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/threatEvents:batchUpdateStatus",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        ThreatEventIds: threatEventIds,
        Status: status,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewBatchUpdateThreatEventsStatusRequestWithoutParam() *BatchUpdateThreatEventsStatusRequest {

    return &BatchUpdateThreatEventsStatusRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/threatEvents:batchUpdateStatus",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param threatEventIds: 事件IDs(Required) */
func (r *BatchUpdateThreatEventsStatusRequest) SetThreatEventIds(threatEventIds []string) {
    r.ThreatEventIds = threatEventIds
}

/* param status: 事件状态(1:忽略;2:误报;3:确认)(Required) */
func (r *BatchUpdateThreatEventsStatusRequest) SetStatus(status int) {
    r.Status = status
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r BatchUpdateThreatEventsStatusRequest) GetRegionId() string {
    return ""
}

type BatchUpdateThreatEventsStatusResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result BatchUpdateThreatEventsStatusResult `json:"result"`
}

type BatchUpdateThreatEventsStatusResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}