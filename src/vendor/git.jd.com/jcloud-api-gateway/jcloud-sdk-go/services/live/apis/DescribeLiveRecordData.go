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
    live "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/live/models"
)

type DescribeLiveRecordDataRequest struct {

    core.JDCloudRequest

    /* 您的推流加速域名  */
    PublishDomain string `json:"publishDomain"`

    /* 直播流所属应用名称 (Optional) */
    AppName *string `json:"appName"`

    /* 直播流名称 (Optional) */
    StreamName *string `json:"streamName"`

    /* 起始时间:
  - UTC 时间格式 e.g: 2019-03-12T00:00:00Z
  */
    StartTime string `json:"startTime"`

    /* 结束时间:
  - UTC 时间格式 e.g: 2019-03-12T00:00:00Z
 (Optional) */
    EndTime *string `json:"endTime"`
}

/*
 * param publishDomain: 您的推流加速域名 (Required)
 * param startTime: 起始时间:
  - UTC 时间格式 e.g: 2019-03-12T00:00:00Z
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLiveRecordDataRequest(
    publishDomain string,
    startTime string,
) *DescribeLiveRecordDataRequest {

	return &DescribeLiveRecordDataRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/liveRecordData",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        PublishDomain: publishDomain,
        StartTime: startTime,
	}
}

/*
 * param publishDomain: 您的推流加速域名 (Required)
 * param appName: 直播流所属应用名称 (Optional)
 * param streamName: 直播流名称 (Optional)
 * param startTime: 起始时间:
  - UTC 时间格式 e.g: 2019-03-12T00:00:00Z
 (Required)
 * param endTime: 结束时间:
  - UTC 时间格式 e.g: 2019-03-12T00:00:00Z
 (Optional)
 */
func NewDescribeLiveRecordDataRequestWithAllParams(
    publishDomain string,
    appName *string,
    streamName *string,
    startTime string,
    endTime *string,
) *DescribeLiveRecordDataRequest {

    return &DescribeLiveRecordDataRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveRecordData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PublishDomain: publishDomain,
        AppName: appName,
        StreamName: streamName,
        StartTime: startTime,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLiveRecordDataRequestWithoutParam() *DescribeLiveRecordDataRequest {

    return &DescribeLiveRecordDataRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveRecordData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param publishDomain: 您的推流加速域名(Required) */
func (r *DescribeLiveRecordDataRequest) SetPublishDomain(publishDomain string) {
    r.PublishDomain = publishDomain
}

/* param appName: 直播流所属应用名称(Optional) */
func (r *DescribeLiveRecordDataRequest) SetAppName(appName string) {
    r.AppName = &appName
}

/* param streamName: 直播流名称(Optional) */
func (r *DescribeLiveRecordDataRequest) SetStreamName(streamName string) {
    r.StreamName = &streamName
}

/* param startTime: 起始时间:
  - UTC 时间格式 e.g: 2019-03-12T00:00:00Z
(Required) */
func (r *DescribeLiveRecordDataRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 结束时间:
  - UTC 时间格式 e.g: 2019-03-12T00:00:00Z
(Optional) */
func (r *DescribeLiveRecordDataRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLiveRecordDataRequest) GetRegionId() string {
    return ""
}

type DescribeLiveRecordDataResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLiveRecordDataResult `json:"result"`
}

type DescribeLiveRecordDataResult struct {
    RecordData []live.RecordData `json:"recordData"`
}