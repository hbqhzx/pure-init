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

type DescribeAlarmsRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 当前所在页，默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 页面大小，默认为20；取值范围[1, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 产品名称 (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /* 资源ID (Optional) */
    ResourceId *string `json:"resourceId"`

    /* 规则类型, 1表示资源监控，6表示站点监控,7表示可用性监控 (Optional) */
    RuleType *int `json:"ruleType"`

    /* 规则报警状态, 1：正常, 2：报警，4：数据不足 (Optional) */
    Status *int `json:"status"`

    /* 规则状态：1为启用，0为禁用 (Optional) */
    Enabled *int `json:"enabled"`

    /* 是否为正在报警的规则，0为忽略，1为是，与 status 同时只能生效一个,isAlarming 优先生效 (Optional) */
    IsAlarming *int `json:"isAlarming"`

    /* 规则的id (Optional) */
    AlarmId *string `json:"alarmId"`

    /* 服务码或资源Id列表
filter name 为serviceCodes表示查询多个产品线的规则
filter name 为resourceIds表示查询多个资源的规则 (Optional) */
    Filters []monitor.Filter `json:"filters"`
}

/*
 * param regionId: 地域 Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAlarmsRequest(
    regionId string,
) *DescribeAlarmsRequest {

	return &DescribeAlarmsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/alarms",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param pageNumber: 当前所在页，默认为1 (Optional)
 * param pageSize: 页面大小，默认为20；取值范围[1, 100] (Optional)
 * param serviceCode: 产品名称 (Optional)
 * param resourceId: 资源ID (Optional)
 * param ruleType: 规则类型, 1表示资源监控，6表示站点监控,7表示可用性监控 (Optional)
 * param status: 规则报警状态, 1：正常, 2：报警，4：数据不足 (Optional)
 * param enabled: 规则状态：1为启用，0为禁用 (Optional)
 * param isAlarming: 是否为正在报警的规则，0为忽略，1为是，与 status 同时只能生效一个,isAlarming 优先生效 (Optional)
 * param alarmId: 规则的id (Optional)
 * param filters: 服务码或资源Id列表
filter name 为serviceCodes表示查询多个产品线的规则
filter name 为resourceIds表示查询多个资源的规则 (Optional)
 */
func NewDescribeAlarmsRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    serviceCode *string,
    resourceId *string,
    ruleType *int,
    status *int,
    enabled *int,
    isAlarming *int,
    alarmId *string,
    filters []monitor.Filter,
) *DescribeAlarmsRequest {

    return &DescribeAlarmsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/alarms",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        ServiceCode: serviceCode,
        ResourceId: resourceId,
        RuleType: ruleType,
        Status: status,
        Enabled: enabled,
        IsAlarming: isAlarming,
        AlarmId: alarmId,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAlarmsRequestWithoutParam() *DescribeAlarmsRequest {

    return &DescribeAlarmsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/alarms",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DescribeAlarmsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 当前所在页，默认为1(Optional) */
func (r *DescribeAlarmsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 页面大小，默认为20；取值范围[1, 100](Optional) */
func (r *DescribeAlarmsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param serviceCode: 产品名称(Optional) */
func (r *DescribeAlarmsRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

/* param resourceId: 资源ID(Optional) */
func (r *DescribeAlarmsRequest) SetResourceId(resourceId string) {
    r.ResourceId = &resourceId
}

/* param ruleType: 规则类型, 1表示资源监控，6表示站点监控,7表示可用性监控(Optional) */
func (r *DescribeAlarmsRequest) SetRuleType(ruleType int) {
    r.RuleType = &ruleType
}

/* param status: 规则报警状态, 1：正常, 2：报警，4：数据不足(Optional) */
func (r *DescribeAlarmsRequest) SetStatus(status int) {
    r.Status = &status
}

/* param enabled: 规则状态：1为启用，0为禁用(Optional) */
func (r *DescribeAlarmsRequest) SetEnabled(enabled int) {
    r.Enabled = &enabled
}

/* param isAlarming: 是否为正在报警的规则，0为忽略，1为是，与 status 同时只能生效一个,isAlarming 优先生效(Optional) */
func (r *DescribeAlarmsRequest) SetIsAlarming(isAlarming int) {
    r.IsAlarming = &isAlarming
}

/* param alarmId: 规则的id(Optional) */
func (r *DescribeAlarmsRequest) SetAlarmId(alarmId string) {
    r.AlarmId = &alarmId
}

/* param filters: 服务码或资源Id列表
filter name 为serviceCodes表示查询多个产品线的规则
filter name 为resourceIds表示查询多个资源的规则(Optional) */
func (r *DescribeAlarmsRequest) SetFilters(filters []monitor.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAlarmsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeAlarmsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAlarmsResult `json:"result"`
}

type DescribeAlarmsResult struct {
    AlarmList []monitor.DescribedAlarm `json:"alarmList"`
    Total int64 `json:"total"`
}