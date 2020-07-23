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

type BatchDeleteLogtopicRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 日志集 UID  */
    LogsetUID string `json:"logsetUID"`

    /* uid列表，以 | 分割  */
    Ids string `json:"ids"`
}

/*
 * param regionId: 地域 Id (Required)
 * param logsetUID: 日志集 UID (Required)
 * param ids: uid列表，以 | 分割 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewBatchDeleteLogtopicRequest(
    regionId string,
    logsetUID string,
    ids string,
) *BatchDeleteLogtopicRequest {

	return &BatchDeleteLogtopicRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/logsets/{logsetUID}/logtopics",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        LogsetUID: logsetUID,
        Ids: ids,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param logsetUID: 日志集 UID (Required)
 * param ids: uid列表，以 | 分割 (Required)
 */
func NewBatchDeleteLogtopicRequestWithAllParams(
    regionId string,
    logsetUID string,
    ids string,
) *BatchDeleteLogtopicRequest {

    return &BatchDeleteLogtopicRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logsets/{logsetUID}/logtopics",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        LogsetUID: logsetUID,
        Ids: ids,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewBatchDeleteLogtopicRequestWithoutParam() *BatchDeleteLogtopicRequest {

    return &BatchDeleteLogtopicRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/logsets/{logsetUID}/logtopics",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *BatchDeleteLogtopicRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param logsetUID: 日志集 UID(Required) */
func (r *BatchDeleteLogtopicRequest) SetLogsetUID(logsetUID string) {
    r.LogsetUID = logsetUID
}

/* param ids: uid列表，以 | 分割(Required) */
func (r *BatchDeleteLogtopicRequest) SetIds(ids string) {
    r.Ids = ids
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r BatchDeleteLogtopicRequest) GetRegionId() string {
    return r.RegionId
}

type BatchDeleteLogtopicResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result BatchDeleteLogtopicResult `json:"result"`
}

type BatchDeleteLogtopicResult struct {
}