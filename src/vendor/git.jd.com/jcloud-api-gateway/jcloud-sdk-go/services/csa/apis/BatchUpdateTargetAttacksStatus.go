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

type BatchUpdateTargetAttacksStatusRequest struct {

    core.JDCloudRequest

    /* 事件IDs  */
    TargetAttackIds []string `json:"targetAttackIds"`

    /* 事件状态(1:忽略;2:误报;3:确认)  */
    Status int `json:"status"`
}

/*
 * param targetAttackIds: 事件IDs (Required)
 * param status: 事件状态(1:忽略;2:误报;3:确认) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewBatchUpdateTargetAttacksStatusRequest(
    targetAttackIds []string,
    status int,
) *BatchUpdateTargetAttacksStatusRequest {

	return &BatchUpdateTargetAttacksStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/targetAttacks:batchUpdateStatus",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        TargetAttackIds: targetAttackIds,
        Status: status,
	}
}

/*
 * param targetAttackIds: 事件IDs (Required)
 * param status: 事件状态(1:忽略;2:误报;3:确认) (Required)
 */
func NewBatchUpdateTargetAttacksStatusRequestWithAllParams(
    targetAttackIds []string,
    status int,
) *BatchUpdateTargetAttacksStatusRequest {

    return &BatchUpdateTargetAttacksStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/targetAttacks:batchUpdateStatus",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        TargetAttackIds: targetAttackIds,
        Status: status,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewBatchUpdateTargetAttacksStatusRequestWithoutParam() *BatchUpdateTargetAttacksStatusRequest {

    return &BatchUpdateTargetAttacksStatusRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/targetAttacks:batchUpdateStatus",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param targetAttackIds: 事件IDs(Required) */
func (r *BatchUpdateTargetAttacksStatusRequest) SetTargetAttackIds(targetAttackIds []string) {
    r.TargetAttackIds = targetAttackIds
}

/* param status: 事件状态(1:忽略;2:误报;3:确认)(Required) */
func (r *BatchUpdateTargetAttacksStatusRequest) SetStatus(status int) {
    r.Status = status
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r BatchUpdateTargetAttacksStatusRequest) GetRegionId() string {
    return ""
}

type BatchUpdateTargetAttacksStatusResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result BatchUpdateTargetAttacksStatusResult `json:"result"`
}

type BatchUpdateTargetAttacksStatusResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}