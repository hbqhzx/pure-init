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

type SendMessageRequest struct {

    core.JDCloudRequest

    /* 所在区域的Region ID  */
    RegionId string `json:"regionId"`

    /* topic 名称  */
    TopicName string `json:"topicName"`

    /* 消息内容，不大于256KB  */
    MessageBody string `json:"messageBody"`

    /* 消息延迟发送的时间单位秒 (Optional) */
    DelaySeconds *int `json:"delaySeconds"`

    /* 消息tag (Optional) */
    Tags []string `json:"tags"`
}

/*
 * param regionId: 所在区域的Region ID (Required)
 * param topicName: topic 名称 (Required)
 * param messageBody: 消息内容，不大于256KB (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSendMessageRequest(
    regionId string,
    topicName string,
    messageBody string,
) *SendMessageRequest {

	return &SendMessageRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/topics/{topicName}/messages",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        TopicName: topicName,
        MessageBody: messageBody,
	}
}

/*
 * param regionId: 所在区域的Region ID (Required)
 * param topicName: topic 名称 (Required)
 * param messageBody: 消息内容，不大于256KB (Required)
 * param delaySeconds: 消息延迟发送的时间单位秒 (Optional)
 * param tags: 消息tag (Optional)
 */
func NewSendMessageRequestWithAllParams(
    regionId string,
    topicName string,
    messageBody string,
    delaySeconds *int,
    tags []string,
) *SendMessageRequest {

    return &SendMessageRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/topics/{topicName}/messages",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        TopicName: topicName,
        MessageBody: messageBody,
        DelaySeconds: delaySeconds,
        Tags: tags,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSendMessageRequestWithoutParam() *SendMessageRequest {

    return &SendMessageRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/topics/{topicName}/messages",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 所在区域的Region ID(Required) */
func (r *SendMessageRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param topicName: topic 名称(Required) */
func (r *SendMessageRequest) SetTopicName(topicName string) {
    r.TopicName = topicName
}

/* param messageBody: 消息内容，不大于256KB(Required) */
func (r *SendMessageRequest) SetMessageBody(messageBody string) {
    r.MessageBody = messageBody
}

/* param delaySeconds: 消息延迟发送的时间单位秒(Optional) */
func (r *SendMessageRequest) SetDelaySeconds(delaySeconds int) {
    r.DelaySeconds = &delaySeconds
}

/* param tags: 消息tag(Optional) */
func (r *SendMessageRequest) SetTags(tags []string) {
    r.Tags = tags
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SendMessageRequest) GetRegionId() string {
    return r.RegionId
}

type SendMessageResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SendMessageResult `json:"result"`
}

type SendMessageResult struct {
    MessageId string `json:"messageId"`
}