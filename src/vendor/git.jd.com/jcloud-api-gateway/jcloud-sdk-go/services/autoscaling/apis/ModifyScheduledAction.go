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

type ModifyScheduledActionRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 组uuid  */
    AsGroupUuid string `json:"asGroupUuid"`

    /* 定时任务uuid  */
    ScheduledActionUuid string `json:"scheduledActionUuid"`

    /* 实例数量  */
    Amount int64 `json:"amount"`

    /* 按天执行时的时间间隔 (Optional) */
    DayInterval *int64 `json:"dayInterval"`

    /* 定时任务描述 (Optional) */
    Description *string `json:"description"`

    /* 执行时间段-终止，当为重复任务是，结束时间必填，格式为 yyyy-MM-dd'T'HH:mm:ssZ ，例如：2018-06-01T00:00:00+0800 (Optional) */
    EndTime *string `json:"endTime"`

    /* d-天, w-周, m-月 (Optional) */
    Every *string `json:"every"`

    /* 按月执行时，执行的结束日期，如每月10号结束执行 (Optional) */
    MonthEnd *int64 `json:"monthEnd"`

    /* 按月执行时，执行的开始日期，如每月5号开始执行 (Optional) */
    MonthStart *int64 `json:"monthStart"`

    /* 定时任务名称  */
    Name string `json:"name"`

    /* 0-仅一次, 1-重复  */
    Recurrence int64 `json:"recurrence"`

    /* 伸缩类型：0-scale out弹出,1-scale in收缩  */
    ScaleType int64 `json:"scaleType"`

    /* 起始时间段-起始，格式为 yyyy-MM-dd'T'HH:mm:ssZ ，例如：2018-06-01T00:00:00+0800  */
    StartTime string `json:"startTime"`

    /* 按周执行时，具体执行的星期几 (Optional) */
    WeekDays []int64 `json:"weekDays"`
}

/*
 * param regionId: 地域 Id (Required)
 * param asGroupUuid: 组uuid (Required)
 * param scheduledActionUuid: 定时任务uuid (Required)
 * param amount: 实例数量 (Required)
 * param name: 定时任务名称 (Required)
 * param recurrence: 0-仅一次, 1-重复 (Required)
 * param scaleType: 伸缩类型：0-scale out弹出,1-scale in收缩 (Required)
 * param startTime: 起始时间段-起始，格式为 yyyy-MM-dd'T'HH:mm:ssZ ，例如：2018-06-01T00:00:00+0800 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyScheduledActionRequest(
    regionId string,
    asGroupUuid string,
    scheduledActionUuid string,
    amount int64,
    name string,
    recurrence int64,
    scaleType int64,
    startTime string,
) *ModifyScheduledActionRequest {

	return &ModifyScheduledActionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/autoscalingGroups/{asGroupUuid}/scheduledActions/{scheduledActionUuid}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AsGroupUuid: asGroupUuid,
        ScheduledActionUuid: scheduledActionUuid,
        Amount: amount,
        Name: name,
        Recurrence: recurrence,
        ScaleType: scaleType,
        StartTime: startTime,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param asGroupUuid: 组uuid (Required)
 * param scheduledActionUuid: 定时任务uuid (Required)
 * param amount: 实例数量 (Required)
 * param dayInterval: 按天执行时的时间间隔 (Optional)
 * param description: 定时任务描述 (Optional)
 * param endTime: 执行时间段-终止，当为重复任务是，结束时间必填，格式为 yyyy-MM-dd'T'HH:mm:ssZ ，例如：2018-06-01T00:00:00+0800 (Optional)
 * param every: d-天, w-周, m-月 (Optional)
 * param monthEnd: 按月执行时，执行的结束日期，如每月10号结束执行 (Optional)
 * param monthStart: 按月执行时，执行的开始日期，如每月5号开始执行 (Optional)
 * param name: 定时任务名称 (Required)
 * param recurrence: 0-仅一次, 1-重复 (Required)
 * param scaleType: 伸缩类型：0-scale out弹出,1-scale in收缩 (Required)
 * param startTime: 起始时间段-起始，格式为 yyyy-MM-dd'T'HH:mm:ssZ ，例如：2018-06-01T00:00:00+0800 (Required)
 * param weekDays: 按周执行时，具体执行的星期几 (Optional)
 */
func NewModifyScheduledActionRequestWithAllParams(
    regionId string,
    asGroupUuid string,
    scheduledActionUuid string,
    amount int64,
    dayInterval *int64,
    description *string,
    endTime *string,
    every *string,
    monthEnd *int64,
    monthStart *int64,
    name string,
    recurrence int64,
    scaleType int64,
    startTime string,
    weekDays []int64,
) *ModifyScheduledActionRequest {

    return &ModifyScheduledActionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/autoscalingGroups/{asGroupUuid}/scheduledActions/{scheduledActionUuid}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AsGroupUuid: asGroupUuid,
        ScheduledActionUuid: scheduledActionUuid,
        Amount: amount,
        DayInterval: dayInterval,
        Description: description,
        EndTime: endTime,
        Every: every,
        MonthEnd: monthEnd,
        MonthStart: monthStart,
        Name: name,
        Recurrence: recurrence,
        ScaleType: scaleType,
        StartTime: startTime,
        WeekDays: weekDays,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyScheduledActionRequestWithoutParam() *ModifyScheduledActionRequest {

    return &ModifyScheduledActionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/autoscalingGroups/{asGroupUuid}/scheduledActions/{scheduledActionUuid}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *ModifyScheduledActionRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param asGroupUuid: 组uuid(Required) */
func (r *ModifyScheduledActionRequest) SetAsGroupUuid(asGroupUuid string) {
    r.AsGroupUuid = asGroupUuid
}

/* param scheduledActionUuid: 定时任务uuid(Required) */
func (r *ModifyScheduledActionRequest) SetScheduledActionUuid(scheduledActionUuid string) {
    r.ScheduledActionUuid = scheduledActionUuid
}

/* param amount: 实例数量(Required) */
func (r *ModifyScheduledActionRequest) SetAmount(amount int64) {
    r.Amount = amount
}

/* param dayInterval: 按天执行时的时间间隔(Optional) */
func (r *ModifyScheduledActionRequest) SetDayInterval(dayInterval int64) {
    r.DayInterval = &dayInterval
}

/* param description: 定时任务描述(Optional) */
func (r *ModifyScheduledActionRequest) SetDescription(description string) {
    r.Description = &description
}

/* param endTime: 执行时间段-终止，当为重复任务是，结束时间必填，格式为 yyyy-MM-dd'T'HH:mm:ssZ ，例如：2018-06-01T00:00:00+0800(Optional) */
func (r *ModifyScheduledActionRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param every: d-天, w-周, m-月(Optional) */
func (r *ModifyScheduledActionRequest) SetEvery(every string) {
    r.Every = &every
}

/* param monthEnd: 按月执行时，执行的结束日期，如每月10号结束执行(Optional) */
func (r *ModifyScheduledActionRequest) SetMonthEnd(monthEnd int64) {
    r.MonthEnd = &monthEnd
}

/* param monthStart: 按月执行时，执行的开始日期，如每月5号开始执行(Optional) */
func (r *ModifyScheduledActionRequest) SetMonthStart(monthStart int64) {
    r.MonthStart = &monthStart
}

/* param name: 定时任务名称(Required) */
func (r *ModifyScheduledActionRequest) SetName(name string) {
    r.Name = name
}

/* param recurrence: 0-仅一次, 1-重复(Required) */
func (r *ModifyScheduledActionRequest) SetRecurrence(recurrence int64) {
    r.Recurrence = recurrence
}

/* param scaleType: 伸缩类型：0-scale out弹出,1-scale in收缩(Required) */
func (r *ModifyScheduledActionRequest) SetScaleType(scaleType int64) {
    r.ScaleType = scaleType
}

/* param startTime: 起始时间段-起始，格式为 yyyy-MM-dd'T'HH:mm:ssZ ，例如：2018-06-01T00:00:00+0800(Required) */
func (r *ModifyScheduledActionRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param weekDays: 按周执行时，具体执行的星期几(Optional) */
func (r *ModifyScheduledActionRequest) SetWeekDays(weekDays []int64) {
    r.WeekDays = weekDays
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyScheduledActionRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyScheduledActionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyScheduledActionResult `json:"result"`
}

type ModifyScheduledActionResult struct {
    Uuid string `json:"uuid"`
}