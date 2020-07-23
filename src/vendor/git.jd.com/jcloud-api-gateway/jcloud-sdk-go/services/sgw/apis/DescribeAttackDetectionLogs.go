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
    sgw "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/sgw/models"
)

type DescribeAttackDetectionLogsRequest struct {

    core.JDCloudRequest

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为10；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 查询开始时间 (Optional) */
    QueryStartTime *int `json:"queryStartTime"`

    /* 查询结束时间 (Optional) */
    QueryEndTime *int `json:"queryEndTime"`

    /* 数据来源，vpc-waf:应用安全网关，cloud-waf:云WAF，sa:态势感知 (Optional) */
    Source *string `json:"source"`

    /* 模型类型，0:默认 (Optional) */
    ModelType *int `json:"modelType"`

    /* 处理状态，0：未处理，1：标记，2：忽略 (Optional) */
    ProcessedStatus *int `json:"processedStatus"`

    /* 域名，支持模糊查询 (Optional) */
    Host *string `json:"host"`

    /* 用户pin (Optional) */
    Pin *string `json:"pin"`

    /* 规则id (Optional) */
    RuleId *string `json:"ruleId"`

    /* 实例id (Optional) */
    WafId *string `json:"wafId"`

    /* 0:误报，1：漏报 (Optional) */
    LogType *int `json:"logType"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAttackDetectionLogsRequest(
) *DescribeAttackDetectionLogsRequest {

	return &DescribeAttackDetectionLogsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/attackDetectionLogs",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为10；取值范围[10, 100] (Optional)
 * param queryStartTime: 查询开始时间 (Optional)
 * param queryEndTime: 查询结束时间 (Optional)
 * param source: 数据来源，vpc-waf:应用安全网关，cloud-waf:云WAF，sa:态势感知 (Optional)
 * param modelType: 模型类型，0:默认 (Optional)
 * param processedStatus: 处理状态，0：未处理，1：标记，2：忽略 (Optional)
 * param host: 域名，支持模糊查询 (Optional)
 * param pin: 用户pin (Optional)
 * param ruleId: 规则id (Optional)
 * param wafId: 实例id (Optional)
 * param logType: 0:误报，1：漏报 (Optional)
 */
func NewDescribeAttackDetectionLogsRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    queryStartTime *int,
    queryEndTime *int,
    source *string,
    modelType *int,
    processedStatus *int,
    host *string,
    pin *string,
    ruleId *string,
    wafId *string,
    logType *int,
) *DescribeAttackDetectionLogsRequest {

    return &DescribeAttackDetectionLogsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/attackDetectionLogs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
        QueryStartTime: queryStartTime,
        QueryEndTime: queryEndTime,
        Source: source,
        ModelType: modelType,
        ProcessedStatus: processedStatus,
        Host: host,
        Pin: pin,
        RuleId: ruleId,
        WafId: wafId,
        LogType: logType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAttackDetectionLogsRequestWithoutParam() *DescribeAttackDetectionLogsRequest {

    return &DescribeAttackDetectionLogsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/attackDetectionLogs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeAttackDetectionLogsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为10；取值范围[10, 100](Optional) */
func (r *DescribeAttackDetectionLogsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param queryStartTime: 查询开始时间(Optional) */
func (r *DescribeAttackDetectionLogsRequest) SetQueryStartTime(queryStartTime int) {
    r.QueryStartTime = &queryStartTime
}

/* param queryEndTime: 查询结束时间(Optional) */
func (r *DescribeAttackDetectionLogsRequest) SetQueryEndTime(queryEndTime int) {
    r.QueryEndTime = &queryEndTime
}

/* param source: 数据来源，vpc-waf:应用安全网关，cloud-waf:云WAF，sa:态势感知(Optional) */
func (r *DescribeAttackDetectionLogsRequest) SetSource(source string) {
    r.Source = &source
}

/* param modelType: 模型类型，0:默认(Optional) */
func (r *DescribeAttackDetectionLogsRequest) SetModelType(modelType int) {
    r.ModelType = &modelType
}

/* param processedStatus: 处理状态，0：未处理，1：标记，2：忽略(Optional) */
func (r *DescribeAttackDetectionLogsRequest) SetProcessedStatus(processedStatus int) {
    r.ProcessedStatus = &processedStatus
}

/* param host: 域名，支持模糊查询(Optional) */
func (r *DescribeAttackDetectionLogsRequest) SetHost(host string) {
    r.Host = &host
}

/* param pin: 用户pin(Optional) */
func (r *DescribeAttackDetectionLogsRequest) SetPin(pin string) {
    r.Pin = &pin
}

/* param ruleId: 规则id(Optional) */
func (r *DescribeAttackDetectionLogsRequest) SetRuleId(ruleId string) {
    r.RuleId = &ruleId
}

/* param wafId: 实例id(Optional) */
func (r *DescribeAttackDetectionLogsRequest) SetWafId(wafId string) {
    r.WafId = &wafId
}

/* param logType: 0:误报，1：漏报(Optional) */
func (r *DescribeAttackDetectionLogsRequest) SetLogType(logType int) {
    r.LogType = &logType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAttackDetectionLogsRequest) GetRegionId() string {
    return ""
}

type DescribeAttackDetectionLogsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAttackDetectionLogsResult `json:"result"`
}

type DescribeAttackDetectionLogsResult struct {
    Logs []sgw.AttackDetectionLog `json:"logs"`
    Total int `json:"total"`
}