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
    mps "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/mps/models"
)

type ListThumbnailTaskRequest struct {

    core.JDCloudRequest

    /* region id  */
    RegionId string `json:"regionId"`

    /* task 状态 (PENDING, RUNNING, SUCCESS, FAILED) (Optional) */
    Status *string `json:"status"`

    /* 开始时间 时间格式(GMT): yyyy-MM-dd'T'HH:mm:ss.SSS'Z' (Optional) */
    Begin *string `json:"begin"`

    /* 结束时间 时间格式(GMT): yyyy-MM-dd'T'HH:mm:ss.SSS'Z' (Optional) */
    End *string `json:"end"`

    /* 查询标记 (Optional) */
    Marker *string `json:"marker"`

    /* 查询记录数 [1, 1000] (Optional) */
    Limit *int `json:"limit"`
}

/*
 * param regionId: region id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewListThumbnailTaskRequest(
    regionId string,
) *ListThumbnailTaskRequest {

	return &ListThumbnailTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/thumbnail",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: region id (Required)
 * param status: task 状态 (PENDING, RUNNING, SUCCESS, FAILED) (Optional)
 * param begin: 开始时间 时间格式(GMT): yyyy-MM-dd'T'HH:mm:ss.SSS'Z' (Optional)
 * param end: 结束时间 时间格式(GMT): yyyy-MM-dd'T'HH:mm:ss.SSS'Z' (Optional)
 * param marker: 查询标记 (Optional)
 * param limit: 查询记录数 [1, 1000] (Optional)
 */
func NewListThumbnailTaskRequestWithAllParams(
    regionId string,
    status *string,
    begin *string,
    end *string,
    marker *string,
    limit *int,
) *ListThumbnailTaskRequest {

    return &ListThumbnailTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/thumbnail",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Status: status,
        Begin: begin,
        End: end,
        Marker: marker,
        Limit: limit,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewListThumbnailTaskRequestWithoutParam() *ListThumbnailTaskRequest {

    return &ListThumbnailTaskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/thumbnail",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: region id(Required) */
func (r *ListThumbnailTaskRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param status: task 状态 (PENDING, RUNNING, SUCCESS, FAILED)(Optional) */
func (r *ListThumbnailTaskRequest) SetStatus(status string) {
    r.Status = &status
}

/* param begin: 开始时间 时间格式(GMT): yyyy-MM-dd'T'HH:mm:ss.SSS'Z'(Optional) */
func (r *ListThumbnailTaskRequest) SetBegin(begin string) {
    r.Begin = &begin
}

/* param end: 结束时间 时间格式(GMT): yyyy-MM-dd'T'HH:mm:ss.SSS'Z'(Optional) */
func (r *ListThumbnailTaskRequest) SetEnd(end string) {
    r.End = &end
}

/* param marker: 查询标记(Optional) */
func (r *ListThumbnailTaskRequest) SetMarker(marker string) {
    r.Marker = &marker
}

/* param limit: 查询记录数 [1, 1000](Optional) */
func (r *ListThumbnailTaskRequest) SetLimit(limit int) {
    r.Limit = &limit
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ListThumbnailTaskRequest) GetRegionId() string {
    return r.RegionId
}

type ListThumbnailTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ListThumbnailTaskResult `json:"result"`
}

type ListThumbnailTaskResult struct {
    Status string `json:"status"`
    Begin string `json:"begin"`
    End string `json:"end"`
    Marker string `json:"marker"`
    Limit int `json:"limit"`
    NextMarker string `json:"nextMarker"`
    Truncated bool `json:"truncated"`
    TaskList []mps.ThumbnailTask `json:"taskList"`
}