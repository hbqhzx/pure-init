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
    jqs "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/jqs/models"
)

type ReceiveMessageRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /*   */
    QueueId string `json:"queueId"`

    /*   */
    Body *jqs.PbReceiveMessageInput `json:"body"`
}

/*
 * param regionId: Region ID (Required)
 * param queueId:  (Required)
 * param body:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewReceiveMessageRequest(
    regionId string,
    queueId string,
    body *jqs.PbReceiveMessageInput,
) *ReceiveMessageRequest {

	return &ReceiveMessageRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/queue/{queueId}/message:receive",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        QueueId: queueId,
        Body: body,
	}
}

/*
 * param regionId: Region ID (Required)
 * param queueId:  (Required)
 * param body:  (Required)
 */
func NewReceiveMessageRequestWithAllParams(
    regionId string,
    queueId string,
    body *jqs.PbReceiveMessageInput,
) *ReceiveMessageRequest {

    return &ReceiveMessageRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/queue/{queueId}/message:receive",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        QueueId: queueId,
        Body: body,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewReceiveMessageRequestWithoutParam() *ReceiveMessageRequest {

    return &ReceiveMessageRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/queue/{queueId}/message:receive",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *ReceiveMessageRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param queueId: (Required) */
func (r *ReceiveMessageRequest) SetQueueId(queueId string) {
    r.QueueId = queueId
}

/* param body: (Required) */
func (r *ReceiveMessageRequest) SetBody(body *jqs.PbReceiveMessageInput) {
    r.Body = body
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ReceiveMessageRequest) GetRegionId() string {
    return r.RegionId
}

type ReceiveMessageResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ReceiveMessageResult `json:"result"`
}

type ReceiveMessageResult struct {
    PbReceiveMessageOutput jqs.PbReceiveMessageOutput `json:"pbReceiveMessageOutput"`
}