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
    jcq "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/jcq/models"
)

type DescribeDeadLetterNumbersWithTopicOpsRequest struct {

    core.JDCloudRequest

    /* 所在区域的Region ID  */
    RegionId string `json:"regionId"`

    /* 运维人员erp账号  */
    Erp string `json:"erp"`

    /* topic 名称  */
    TopicName string `json:"topicName"`

    /* 用户pin  */
    UserPin string `json:"userPin"`

    /* consumerGroupId为空则为按照Topic名称查找 (Optional) */
    ConsumerGroupId *string `json:"consumerGroupId"`

    /* 页码 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 每页数 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 所在区域的Region ID (Required)
 * param erp: 运维人员erp账号 (Required)
 * param topicName: topic 名称 (Required)
 * param userPin: 用户pin (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDeadLetterNumbersWithTopicOpsRequest(
    regionId string,
    erp string,
    topicName string,
    userPin string,
) *DescribeDeadLetterNumbersWithTopicOpsRequest {

	return &DescribeDeadLetterNumbersWithTopicOpsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/ops/{erp}/topics/{topicName}/deadLetterNumbers",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Erp: erp,
        TopicName: topicName,
        UserPin: userPin,
	}
}

/*
 * param regionId: 所在区域的Region ID (Required)
 * param erp: 运维人员erp账号 (Required)
 * param topicName: topic 名称 (Required)
 * param userPin: 用户pin (Required)
 * param consumerGroupId: consumerGroupId为空则为按照Topic名称查找 (Optional)
 * param pageNumber: 页码 (Optional)
 * param pageSize: 每页数 (Optional)
 */
func NewDescribeDeadLetterNumbersWithTopicOpsRequestWithAllParams(
    regionId string,
    erp string,
    topicName string,
    userPin string,
    consumerGroupId *string,
    pageNumber *int,
    pageSize *int,
) *DescribeDeadLetterNumbersWithTopicOpsRequest {

    return &DescribeDeadLetterNumbersWithTopicOpsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ops/{erp}/topics/{topicName}/deadLetterNumbers",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Erp: erp,
        TopicName: topicName,
        UserPin: userPin,
        ConsumerGroupId: consumerGroupId,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDeadLetterNumbersWithTopicOpsRequestWithoutParam() *DescribeDeadLetterNumbersWithTopicOpsRequest {

    return &DescribeDeadLetterNumbersWithTopicOpsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ops/{erp}/topics/{topicName}/deadLetterNumbers",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 所在区域的Region ID(Required) */
func (r *DescribeDeadLetterNumbersWithTopicOpsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param erp: 运维人员erp账号(Required) */
func (r *DescribeDeadLetterNumbersWithTopicOpsRequest) SetErp(erp string) {
    r.Erp = erp
}

/* param topicName: topic 名称(Required) */
func (r *DescribeDeadLetterNumbersWithTopicOpsRequest) SetTopicName(topicName string) {
    r.TopicName = topicName
}

/* param userPin: 用户pin(Required) */
func (r *DescribeDeadLetterNumbersWithTopicOpsRequest) SetUserPin(userPin string) {
    r.UserPin = userPin
}

/* param consumerGroupId: consumerGroupId为空则为按照Topic名称查找(Optional) */
func (r *DescribeDeadLetterNumbersWithTopicOpsRequest) SetConsumerGroupId(consumerGroupId string) {
    r.ConsumerGroupId = &consumerGroupId
}

/* param pageNumber: 页码(Optional) */
func (r *DescribeDeadLetterNumbersWithTopicOpsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 每页数(Optional) */
func (r *DescribeDeadLetterNumbersWithTopicOpsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDeadLetterNumbersWithTopicOpsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDeadLetterNumbersWithTopicOpsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDeadLetterNumbersWithTopicOpsResult `json:"result"`
}

type DescribeDeadLetterNumbersWithTopicOpsResult struct {
    DeadLetterNumbers []jcq.DeadLetterNumber `json:"deadLetterNumbers"`
    TotalCount int `json:"totalCount"`
}