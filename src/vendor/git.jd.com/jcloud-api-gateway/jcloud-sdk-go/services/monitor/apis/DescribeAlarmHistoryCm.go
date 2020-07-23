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

type DescribeAlarmHistoryCmRequest struct {

    core.JDCloudRequest

    /* region  */
    RegionId string `json:"regionId"`

    /* 当前所在页，默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 页面大小，默认为20；取值范围[1, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 报警规则的Id (Optional) */
    Id *int `json:"id"`

    /* obj (Optional) */
    Obi *string `json:"obi"`

    /* namespace (Optional) */
    Namespace *string `json:"namespace"`

    /* 产品名称 (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /* 资源Id (Optional) */
    ResourceId *string `json:"resourceId"`

    /* 查询数据开始时间，默认24小时前，可以输入long型时间，也可以输入yyyy-MM-dd'T'HH:mm:ssZ类型时间  */
    StartTime string `json:"startTime"`

    /* 查询数据结束时间，默认当前时间，可以输入long型时间，也可以输入yyyy-MM-dd'T'HH:mm:ssZ类型时间  */
    EndTime string `json:"endTime"`
}

/*
 * param regionId: region (Required)
 * param startTime: 查询数据开始时间，默认24小时前，可以输入long型时间，也可以输入yyyy-MM-dd'T'HH:mm:ssZ类型时间 (Required)
 * param endTime: 查询数据结束时间，默认当前时间，可以输入long型时间，也可以输入yyyy-MM-dd'T'HH:mm:ssZ类型时间 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAlarmHistoryCmRequest(
    regionId string,
    startTime string,
    endTime string,
) *DescribeAlarmHistoryCmRequest {

	return &DescribeAlarmHistoryCmRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cmAlarmHistory",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        StartTime: startTime,
        EndTime: endTime,
	}
}

/*
 * param regionId: region (Required)
 * param pageNumber: 当前所在页，默认为1 (Optional)
 * param pageSize: 页面大小，默认为20；取值范围[1, 100] (Optional)
 * param id: 报警规则的Id (Optional)
 * param obi: obj (Optional)
 * param namespace: namespace (Optional)
 * param serviceCode: 产品名称 (Optional)
 * param resourceId: 资源Id (Optional)
 * param startTime: 查询数据开始时间，默认24小时前，可以输入long型时间，也可以输入yyyy-MM-dd'T'HH:mm:ssZ类型时间 (Required)
 * param endTime: 查询数据结束时间，默认当前时间，可以输入long型时间，也可以输入yyyy-MM-dd'T'HH:mm:ssZ类型时间 (Required)
 */
func NewDescribeAlarmHistoryCmRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    id *int,
    obi *string,
    namespace *string,
    serviceCode *string,
    resourceId *string,
    startTime string,
    endTime string,
) *DescribeAlarmHistoryCmRequest {

    return &DescribeAlarmHistoryCmRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cmAlarmHistory",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Id: id,
        Obi: obi,
        Namespace: namespace,
        ServiceCode: serviceCode,
        ResourceId: resourceId,
        StartTime: startTime,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAlarmHistoryCmRequestWithoutParam() *DescribeAlarmHistoryCmRequest {

    return &DescribeAlarmHistoryCmRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cmAlarmHistory",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: region(Required) */
func (r *DescribeAlarmHistoryCmRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 当前所在页，默认为1(Optional) */
func (r *DescribeAlarmHistoryCmRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 页面大小，默认为20；取值范围[1, 100](Optional) */
func (r *DescribeAlarmHistoryCmRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param id: 报警规则的Id(Optional) */
func (r *DescribeAlarmHistoryCmRequest) SetId(id int) {
    r.Id = &id
}

/* param obi: obj(Optional) */
func (r *DescribeAlarmHistoryCmRequest) SetObi(obi string) {
    r.Obi = &obi
}

/* param namespace: namespace(Optional) */
func (r *DescribeAlarmHistoryCmRequest) SetNamespace(namespace string) {
    r.Namespace = &namespace
}

/* param serviceCode: 产品名称(Optional) */
func (r *DescribeAlarmHistoryCmRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

/* param resourceId: 资源Id(Optional) */
func (r *DescribeAlarmHistoryCmRequest) SetResourceId(resourceId string) {
    r.ResourceId = &resourceId
}

/* param startTime: 查询数据开始时间，默认24小时前，可以输入long型时间，也可以输入yyyy-MM-dd'T'HH:mm:ssZ类型时间(Required) */
func (r *DescribeAlarmHistoryCmRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 查询数据结束时间，默认当前时间，可以输入long型时间，也可以输入yyyy-MM-dd'T'HH:mm:ssZ类型时间(Required) */
func (r *DescribeAlarmHistoryCmRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAlarmHistoryCmRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeAlarmHistoryCmResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAlarmHistoryCmResult `json:"result"`
}

type DescribeAlarmHistoryCmResult struct {
    AlarmHistoryList []monitor.CmAlarmHistory `json:"alarmHistoryList"`
    NumberPages int64 `json:"numberPages"`
    NumberRecords int64 `json:"numberRecords"`
    PageNumber int64 `json:"pageNumber"`
    PageSize int64 `json:"pageSize"`
}