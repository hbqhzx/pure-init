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
    monitor "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/monitor/models"
)

type DescribeMetricDataAmRequest struct {

    core.JDCloudRequest

    /* 资源的uuid  */
    ResourceId string `json:"resourceId"`

    /* 查询时间范围的开始时间， UTC时间，格式：yyyy-MM-dd'T'HH:mm:ssZ（默认为当前时间，早于30d时，将被重置为30d） (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询时间范围的结束时间， UTC时间，格式：2016-12- yyyy-MM-dd'T'HH:mm:ssZ（为空时，将由startTime与timeInterval计算得出） (Optional) */
    EndTime *string `json:"endTime"`

    /* 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval 与 endTime 至少填一项 (Optional) */
    TimeInterval *string `json:"timeInterval"`
}

/*
 * param resourceId: 资源的uuid (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeMetricDataAmRequest(
    resourceId string,
) *DescribeMetricDataAmRequest {

	return &DescribeMetricDataAmRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/amMetricData",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        ResourceId: resourceId,
	}
}

/*
 * param resourceId: 资源的uuid (Required)
 * param startTime: 查询时间范围的开始时间， UTC时间，格式：yyyy-MM-dd'T'HH:mm:ssZ（默认为当前时间，早于30d时，将被重置为30d） (Optional)
 * param endTime: 查询时间范围的结束时间， UTC时间，格式：2016-12- yyyy-MM-dd'T'HH:mm:ssZ（为空时，将由startTime与timeInterval计算得出） (Optional)
 * param timeInterval: 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval 与 endTime 至少填一项 (Optional)
 */
func NewDescribeMetricDataAmRequestWithAllParams(
    resourceId string,
    startTime *string,
    endTime *string,
    timeInterval *string,
) *DescribeMetricDataAmRequest {

    return &DescribeMetricDataAmRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/amMetricData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        ResourceId: resourceId,
        StartTime: startTime,
        EndTime: endTime,
        TimeInterval: timeInterval,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeMetricDataAmRequestWithoutParam() *DescribeMetricDataAmRequest {

    return &DescribeMetricDataAmRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/amMetricData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param resourceId: 资源的uuid(Required) */
func (r *DescribeMetricDataAmRequest) SetResourceId(resourceId string) {
    r.ResourceId = resourceId
}

/* param startTime: 查询时间范围的开始时间， UTC时间，格式：yyyy-MM-dd'T'HH:mm:ssZ（默认为当前时间，早于30d时，将被重置为30d）(Optional) */
func (r *DescribeMetricDataAmRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 查询时间范围的结束时间， UTC时间，格式：2016-12- yyyy-MM-dd'T'HH:mm:ssZ（为空时，将由startTime与timeInterval计算得出）(Optional) */
func (r *DescribeMetricDataAmRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param timeInterval: 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval 与 endTime 至少填一项(Optional) */
func (r *DescribeMetricDataAmRequest) SetTimeInterval(timeInterval string) {
    r.TimeInterval = &timeInterval
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeMetricDataAmRequest) GetRegionId() string {
    return ""
}

type DescribeMetricDataAmResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeMetricDataAmResult `json:"result"`
}

type DescribeMetricDataAmResult struct {
    MetricDatas []monitor.MetricData `json:"metricDatas"`
}