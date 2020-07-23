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

type DescribeStatisticsRequest struct {

    core.JDCloudRequest

    /* 资源的类型，取值vm, lb, ip, database 等  */
    ServiceCode string `json:"serviceCode"`

    /* 资源的uuid (Optional) */
    ResourceId *string `json:"resourceId"`

    /* 聚合方式，默认：sum、可选值：sum、avg、max、min (Optional) */
    Aggregate *string `json:"aggregate"`

    /* 要查询的metric  */
    Metric string `json:"metric"`

    /* 自定义标签 (Optional) */
    Tags []monitor.TagFilter `json:"tags"`

    /* 查询时间范围的开始时间， UTC时间，格式：yyyy-MM-dd'T'HH:mm:ssZ（默认为当前时间，早于30d时，将被重置为30d） (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询时间范围的结束时间， UTC时间，格式：2016-12- yyyy-MM-dd'T'HH:mm:ssZ（为空时，将由startTime与timeInterval计算得出） (Optional) */
    EndTime *string `json:"endTime"`

    /* 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval 与 endTime 至少填一项 (Optional) */
    TimeInterval *string `json:"timeInterval"`
}

/*
 * param serviceCode: 资源的类型，取值vm, lb, ip, database 等 (Required)
 * param metric: 要查询的metric (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeStatisticsRequest(
    serviceCode string,
    metric string,
) *DescribeStatisticsRequest {

	return &DescribeStatisticsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/statistics",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        ServiceCode: serviceCode,
        Metric: metric,
	}
}

/*
 * param serviceCode: 资源的类型，取值vm, lb, ip, database 等 (Required)
 * param resourceId: 资源的uuid (Optional)
 * param aggregate: 聚合方式，默认：sum、可选值：sum、avg、max、min (Optional)
 * param metric: 要查询的metric (Required)
 * param tags: 自定义标签 (Optional)
 * param startTime: 查询时间范围的开始时间， UTC时间，格式：yyyy-MM-dd'T'HH:mm:ssZ（默认为当前时间，早于30d时，将被重置为30d） (Optional)
 * param endTime: 查询时间范围的结束时间， UTC时间，格式：2016-12- yyyy-MM-dd'T'HH:mm:ssZ（为空时，将由startTime与timeInterval计算得出） (Optional)
 * param timeInterval: 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval 与 endTime 至少填一项 (Optional)
 */
func NewDescribeStatisticsRequestWithAllParams(
    serviceCode string,
    resourceId *string,
    aggregate *string,
    metric string,
    tags []monitor.TagFilter,
    startTime *string,
    endTime *string,
    timeInterval *string,
) *DescribeStatisticsRequest {

    return &DescribeStatisticsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/statistics",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        ServiceCode: serviceCode,
        ResourceId: resourceId,
        Aggregate: aggregate,
        Metric: metric,
        Tags: tags,
        StartTime: startTime,
        EndTime: endTime,
        TimeInterval: timeInterval,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeStatisticsRequestWithoutParam() *DescribeStatisticsRequest {

    return &DescribeStatisticsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/statistics",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param serviceCode: 资源的类型，取值vm, lb, ip, database 等(Required) */
func (r *DescribeStatisticsRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = serviceCode
}

/* param resourceId: 资源的uuid(Optional) */
func (r *DescribeStatisticsRequest) SetResourceId(resourceId string) {
    r.ResourceId = &resourceId
}

/* param aggregate: 聚合方式，默认：sum、可选值：sum、avg、max、min(Optional) */
func (r *DescribeStatisticsRequest) SetAggregate(aggregate string) {
    r.Aggregate = &aggregate
}

/* param metric: 要查询的metric(Required) */
func (r *DescribeStatisticsRequest) SetMetric(metric string) {
    r.Metric = metric
}

/* param tags: 自定义标签(Optional) */
func (r *DescribeStatisticsRequest) SetTags(tags []monitor.TagFilter) {
    r.Tags = tags
}

/* param startTime: 查询时间范围的开始时间， UTC时间，格式：yyyy-MM-dd'T'HH:mm:ssZ（默认为当前时间，早于30d时，将被重置为30d）(Optional) */
func (r *DescribeStatisticsRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 查询时间范围的结束时间， UTC时间，格式：2016-12- yyyy-MM-dd'T'HH:mm:ssZ（为空时，将由startTime与timeInterval计算得出）(Optional) */
func (r *DescribeStatisticsRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param timeInterval: 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval 与 endTime 至少填一项(Optional) */
func (r *DescribeStatisticsRequest) SetTimeInterval(timeInterval string) {
    r.TimeInterval = &timeInterval
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeStatisticsRequest) GetRegionId() string {
    return ""
}

type DescribeStatisticsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeStatisticsResult `json:"result"`
}

type DescribeStatisticsResult struct {
    List []monitor.StatsItem `json:"list"`
}