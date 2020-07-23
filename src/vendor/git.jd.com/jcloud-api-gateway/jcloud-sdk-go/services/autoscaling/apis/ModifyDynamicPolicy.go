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

type ModifyDynamicPolicyRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 组uuid  */
    AsGroupUuid string `json:"asGroupUuid"`

    /* 动态策略uuid  */
    DynamicPolicyUuid string `json:"dynamicPolicyUuid"`

    /* 动作类型 1-incr, 2-set(暂不支持), 3-decr  */
    Action int64 `json:"action"`

    /* 动作值, 实例个数 或 百分比数值  */
    ActionValue int64 `json:"actionValue"`

    /* 动作类型, 实例(1) 或 百分比(2,目前暂时不支持)  */
    ActionValueType int64 `json:"actionValueType"`

    /* 统计方法：平均值=avg、最大值=max、最小值=min,总和：sum  */
    Calculation string `json:"calculation"`

    /* 预热（冷却）时间。秒。若缺，缺省300秒。  */
    CooldownPeriod int64 `json:"cooldownPeriod"`

    /* 动态策略描述 (Optional) */
    Description *string `json:"description"`

    /* 0-不被保护，1-被保护 (Optional) */
    DisableScaleIn *int64 `json:"disableScaleIn"`

    /* 监控项  */
    Metric string `json:"metric"`

    /* 动态策略名字  */
    Name string `json:"name"`

    /* 1：>=、2：>、3：<、4：<=、5：=、6：!= 传数字  */
    Operation string `json:"operation"`

    /* 统计周期（单位：分钟）  */
    Period int64 `json:"period"`

    /* 报警的阈值  */
    Threshold float64 `json:"threshold"`

    /* 连续多少次后报警  */
    Times int64 `json:"times"`
}

/*
 * param regionId: 地域 Id (Required)
 * param asGroupUuid: 组uuid (Required)
 * param dynamicPolicyUuid: 动态策略uuid (Required)
 * param action: 动作类型 1-incr, 2-set(暂不支持), 3-decr (Required)
 * param actionValue: 动作值, 实例个数 或 百分比数值 (Required)
 * param actionValueType: 动作类型, 实例(1) 或 百分比(2,目前暂时不支持) (Required)
 * param calculation: 统计方法：平均值=avg、最大值=max、最小值=min,总和：sum (Required)
 * param cooldownPeriod: 预热（冷却）时间。秒。若缺，缺省300秒。 (Required)
 * param metric: 监控项 (Required)
 * param name: 动态策略名字 (Required)
 * param operation: 1：>=、2：>、3：<、4：<=、5：=、6：!= 传数字 (Required)
 * param period: 统计周期（单位：分钟） (Required)
 * param threshold: 报警的阈值 (Required)
 * param times: 连续多少次后报警 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyDynamicPolicyRequest(
    regionId string,
    asGroupUuid string,
    dynamicPolicyUuid string,
    action int64,
    actionValue int64,
    actionValueType int64,
    calculation string,
    cooldownPeriod int64,
    metric string,
    name string,
    operation string,
    period int64,
    threshold float64,
    times int64,
) *ModifyDynamicPolicyRequest {

	return &ModifyDynamicPolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/autoscalingGroups/{asGroupUuid}/dynamicPolicies/{dynamicPolicyUuid}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AsGroupUuid: asGroupUuid,
        DynamicPolicyUuid: dynamicPolicyUuid,
        Action: action,
        ActionValue: actionValue,
        ActionValueType: actionValueType,
        Calculation: calculation,
        CooldownPeriod: cooldownPeriod,
        Metric: metric,
        Name: name,
        Operation: operation,
        Period: period,
        Threshold: threshold,
        Times: times,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param asGroupUuid: 组uuid (Required)
 * param dynamicPolicyUuid: 动态策略uuid (Required)
 * param action: 动作类型 1-incr, 2-set(暂不支持), 3-decr (Required)
 * param actionValue: 动作值, 实例个数 或 百分比数值 (Required)
 * param actionValueType: 动作类型, 实例(1) 或 百分比(2,目前暂时不支持) (Required)
 * param calculation: 统计方法：平均值=avg、最大值=max、最小值=min,总和：sum (Required)
 * param cooldownPeriod: 预热（冷却）时间。秒。若缺，缺省300秒。 (Required)
 * param description: 动态策略描述 (Optional)
 * param disableScaleIn: 0-不被保护，1-被保护 (Optional)
 * param metric: 监控项 (Required)
 * param name: 动态策略名字 (Required)
 * param operation: 1：>=、2：>、3：<、4：<=、5：=、6：!= 传数字 (Required)
 * param period: 统计周期（单位：分钟） (Required)
 * param threshold: 报警的阈值 (Required)
 * param times: 连续多少次后报警 (Required)
 */
func NewModifyDynamicPolicyRequestWithAllParams(
    regionId string,
    asGroupUuid string,
    dynamicPolicyUuid string,
    action int64,
    actionValue int64,
    actionValueType int64,
    calculation string,
    cooldownPeriod int64,
    description *string,
    disableScaleIn *int64,
    metric string,
    name string,
    operation string,
    period int64,
    threshold float64,
    times int64,
) *ModifyDynamicPolicyRequest {

    return &ModifyDynamicPolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/autoscalingGroups/{asGroupUuid}/dynamicPolicies/{dynamicPolicyUuid}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AsGroupUuid: asGroupUuid,
        DynamicPolicyUuid: dynamicPolicyUuid,
        Action: action,
        ActionValue: actionValue,
        ActionValueType: actionValueType,
        Calculation: calculation,
        CooldownPeriod: cooldownPeriod,
        Description: description,
        DisableScaleIn: disableScaleIn,
        Metric: metric,
        Name: name,
        Operation: operation,
        Period: period,
        Threshold: threshold,
        Times: times,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyDynamicPolicyRequestWithoutParam() *ModifyDynamicPolicyRequest {

    return &ModifyDynamicPolicyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/autoscalingGroups/{asGroupUuid}/dynamicPolicies/{dynamicPolicyUuid}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *ModifyDynamicPolicyRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param asGroupUuid: 组uuid(Required) */
func (r *ModifyDynamicPolicyRequest) SetAsGroupUuid(asGroupUuid string) {
    r.AsGroupUuid = asGroupUuid
}

/* param dynamicPolicyUuid: 动态策略uuid(Required) */
func (r *ModifyDynamicPolicyRequest) SetDynamicPolicyUuid(dynamicPolicyUuid string) {
    r.DynamicPolicyUuid = dynamicPolicyUuid
}

/* param action: 动作类型 1-incr, 2-set(暂不支持), 3-decr(Required) */
func (r *ModifyDynamicPolicyRequest) SetAction(action int64) {
    r.Action = action
}

/* param actionValue: 动作值, 实例个数 或 百分比数值(Required) */
func (r *ModifyDynamicPolicyRequest) SetActionValue(actionValue int64) {
    r.ActionValue = actionValue
}

/* param actionValueType: 动作类型, 实例(1) 或 百分比(2,目前暂时不支持)(Required) */
func (r *ModifyDynamicPolicyRequest) SetActionValueType(actionValueType int64) {
    r.ActionValueType = actionValueType
}

/* param calculation: 统计方法：平均值=avg、最大值=max、最小值=min,总和：sum(Required) */
func (r *ModifyDynamicPolicyRequest) SetCalculation(calculation string) {
    r.Calculation = calculation
}

/* param cooldownPeriod: 预热（冷却）时间。秒。若缺，缺省300秒。(Required) */
func (r *ModifyDynamicPolicyRequest) SetCooldownPeriod(cooldownPeriod int64) {
    r.CooldownPeriod = cooldownPeriod
}

/* param description: 动态策略描述(Optional) */
func (r *ModifyDynamicPolicyRequest) SetDescription(description string) {
    r.Description = &description
}

/* param disableScaleIn: 0-不被保护，1-被保护(Optional) */
func (r *ModifyDynamicPolicyRequest) SetDisableScaleIn(disableScaleIn int64) {
    r.DisableScaleIn = &disableScaleIn
}

/* param metric: 监控项(Required) */
func (r *ModifyDynamicPolicyRequest) SetMetric(metric string) {
    r.Metric = metric
}

/* param name: 动态策略名字(Required) */
func (r *ModifyDynamicPolicyRequest) SetName(name string) {
    r.Name = name
}

/* param operation: 1：>=、2：>、3：<、4：<=、5：=、6：!= 传数字(Required) */
func (r *ModifyDynamicPolicyRequest) SetOperation(operation string) {
    r.Operation = operation
}

/* param period: 统计周期（单位：分钟）(Required) */
func (r *ModifyDynamicPolicyRequest) SetPeriod(period int64) {
    r.Period = period
}

/* param threshold: 报警的阈值(Required) */
func (r *ModifyDynamicPolicyRequest) SetThreshold(threshold float64) {
    r.Threshold = threshold
}

/* param times: 连续多少次后报警(Required) */
func (r *ModifyDynamicPolicyRequest) SetTimes(times int64) {
    r.Times = times
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyDynamicPolicyRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyDynamicPolicyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyDynamicPolicyResult `json:"result"`
}

type ModifyDynamicPolicyResult struct {
    Uuid string `json:"uuid"`
}