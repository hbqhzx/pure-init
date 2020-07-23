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

type CreateConsumerGroupRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 消费组对象 (Optional) */
    ConsumerGroupStr *streambus.ConsumerGroup `json:"consumerGroupStr"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateConsumerGroupRequest(
    regionId string,
) *CreateConsumerGroupRequest {

	return &CreateConsumerGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/consumerGroup",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param consumerGroupStr: 消费组对象 (Optional)
 */
func NewCreateConsumerGroupRequestWithAllParams(
    regionId string,
    consumerGroupStr *streambus.ConsumerGroup,
) *CreateConsumerGroupRequest {

    return &CreateConsumerGroupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/consumerGroup",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ConsumerGroupStr: consumerGroupStr,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateConsumerGroupRequestWithoutParam() *CreateConsumerGroupRequest {

    return &CreateConsumerGroupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/consumerGroup",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateConsumerGroupRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param consumerGroupStr: 消费组对象(Optional) */
func (r *CreateConsumerGroupRequest) SetConsumerGroupStr(consumerGroupStr *streambus.ConsumerGroup) {
    r.ConsumerGroupStr = consumerGroupStr
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateConsumerGroupRequest) GetRegionId() string {
    return r.RegionId
}

type CreateConsumerGroupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateConsumerGroupResult `json:"result"`
}

type CreateConsumerGroupResult struct {
    Message string `json:"message"`
    Status bool `json:"status"`
}