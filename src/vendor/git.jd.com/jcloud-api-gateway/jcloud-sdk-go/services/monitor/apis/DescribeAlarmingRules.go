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

type DescribeAlarmingRulesRequest struct {

    core.JDCloudRequest

    /* 当前所在页，默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 页面大小，默认为20；取值范围[1, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 服务码列表
filter name 为serviceCodes表示查询多个产品线的规则 (Optional) */
    Filters []monitor.Filter `json:"filters"`

    /* 要查询的地域，为空则查询所有的 (Optional) */
    Datacenter *string `json:"datacenter"`

    /* ruleType,默认1 (Optional) */
    RuleType *int `json:"ruleType"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAlarmingRulesRequest(
) *DescribeAlarmingRulesRequest {

	return &DescribeAlarmingRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/overviewAlarmingRules",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNumber: 当前所在页，默认为1 (Optional)
 * param pageSize: 页面大小，默认为20；取值范围[1, 100] (Optional)
 * param filters: 服务码列表
filter name 为serviceCodes表示查询多个产品线的规则 (Optional)
 * param datacenter: 要查询的地域，为空则查询所有的 (Optional)
 * param ruleType: ruleType,默认1 (Optional)
 */
func NewDescribeAlarmingRulesRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    filters []monitor.Filter,
    datacenter *string,
    ruleType *int,
) *DescribeAlarmingRulesRequest {

    return &DescribeAlarmingRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/overviewAlarmingRules",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
        Datacenter: datacenter,
        RuleType: ruleType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAlarmingRulesRequestWithoutParam() *DescribeAlarmingRulesRequest {

    return &DescribeAlarmingRulesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/overviewAlarmingRules",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNumber: 当前所在页，默认为1(Optional) */
func (r *DescribeAlarmingRulesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 页面大小，默认为20；取值范围[1, 100](Optional) */
func (r *DescribeAlarmingRulesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: 服务码列表
filter name 为serviceCodes表示查询多个产品线的规则(Optional) */
func (r *DescribeAlarmingRulesRequest) SetFilters(filters []monitor.Filter) {
    r.Filters = filters
}

/* param datacenter: 要查询的地域，为空则查询所有的(Optional) */
func (r *DescribeAlarmingRulesRequest) SetDatacenter(datacenter string) {
    r.Datacenter = &datacenter
}

/* param ruleType: ruleType,默认1(Optional) */
func (r *DescribeAlarmingRulesRequest) SetRuleType(ruleType int) {
    r.RuleType = &ruleType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAlarmingRulesRequest) GetRegionId() string {
    return ""
}

type DescribeAlarmingRulesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAlarmingRulesResult `json:"result"`
}

type DescribeAlarmingRulesResult struct {
    AlarmHistoryList []monitor.AlarmHistoryWithDetail `json:"alarmHistoryList"`
    NumberPages int64 `json:"numberPages"`
    NumberRecords int64 `json:"numberRecords"`
    PageNumber int64 `json:"pageNumber"`
    PageSize int64 `json:"pageSize"`
}