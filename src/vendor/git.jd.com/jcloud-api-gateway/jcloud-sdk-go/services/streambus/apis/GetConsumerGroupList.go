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
    streambus "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/streambus/models"
)

type GetConsumerGroupListRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 主题id  */
    TopicId int `json:"topicId"`
}

/*
 * param regionId: Region ID (Required)
 * param topicId: 主题id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetConsumerGroupListRequest(
    regionId string,
    topicId int,
) *GetConsumerGroupListRequest {

	return &GetConsumerGroupListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/consumerGroupList",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        TopicId: topicId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param topicId: 主题id (Required)
 */
func NewGetConsumerGroupListRequestWithAllParams(
    regionId string,
    topicId int,
) *GetConsumerGroupListRequest {

    return &GetConsumerGroupListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/consumerGroupList",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        TopicId: topicId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetConsumerGroupListRequestWithoutParam() *GetConsumerGroupListRequest {

    return &GetConsumerGroupListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/consumerGroupList",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *GetConsumerGroupListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param topicId: 主题id(Required) */
func (r *GetConsumerGroupListRequest) SetTopicId(topicId int) {
    r.TopicId = topicId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetConsumerGroupListRequest) GetRegionId() string {
    return r.RegionId
}

type GetConsumerGroupListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetConsumerGroupListResult `json:"result"`
}

type GetConsumerGroupListResult struct {
    ConsumerGroup []streambus.ConsumerGroup `json:"consumerGroup"`
}