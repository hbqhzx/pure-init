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

type ModifySubscriptionAttributeRequest struct {

    core.JDCloudRequest

    /* 所在区域的Region ID  */
    RegionId string `json:"regionId"`

    /* topic 名称  */
    TopicName string `json:"topicName"`

    /* consumerGroupId  */
    ConsumerGroupId string `json:"consumerGroupId"`

    /* endpoint  */
    Endpoint string `json:"endpoint"`
}

/*
 * param regionId: 所在区域的Region ID (Required)
 * param topicName: topic 名称 (Required)
 * param consumerGroupId: consumerGroupId (Required)
 * param endpoint: endpoint (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifySubscriptionAttributeRequest(
    regionId string,
    topicName string,
    consumerGroupId string,
    endpoint string,
) *ModifySubscriptionAttributeRequest {

	return &ModifySubscriptionAttributeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/topics/{topicName}/subscriptions/{consumerGroupId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        TopicName: topicName,
        ConsumerGroupId: consumerGroupId,
        Endpoint: endpoint,
	}
}

/*
 * param regionId: 所在区域的Region ID (Required)
 * param topicName: topic 名称 (Required)
 * param consumerGroupId: consumerGroupId (Required)
 * param endpoint: endpoint (Required)
 */
func NewModifySubscriptionAttributeRequestWithAllParams(
    regionId string,
    topicName string,
    consumerGroupId string,
    endpoint string,
) *ModifySubscriptionAttributeRequest {

    return &ModifySubscriptionAttributeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/topics/{topicName}/subscriptions/{consumerGroupId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        TopicName: topicName,
        ConsumerGroupId: consumerGroupId,
        Endpoint: endpoint,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifySubscriptionAttributeRequestWithoutParam() *ModifySubscriptionAttributeRequest {

    return &ModifySubscriptionAttributeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/topics/{topicName}/subscriptions/{consumerGroupId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 所在区域的Region ID(Required) */
func (r *ModifySubscriptionAttributeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param topicName: topic 名称(Required) */
func (r *ModifySubscriptionAttributeRequest) SetTopicName(topicName string) {
    r.TopicName = topicName
}

/* param consumerGroupId: consumerGroupId(Required) */
func (r *ModifySubscriptionAttributeRequest) SetConsumerGroupId(consumerGroupId string) {
    r.ConsumerGroupId = consumerGroupId
}

/* param endpoint: endpoint(Required) */
func (r *ModifySubscriptionAttributeRequest) SetEndpoint(endpoint string) {
    r.Endpoint = endpoint
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifySubscriptionAttributeRequest) GetRegionId() string {
    return r.RegionId
}

type ModifySubscriptionAttributeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifySubscriptionAttributeResult `json:"result"`
}

type ModifySubscriptionAttributeResult struct {
}