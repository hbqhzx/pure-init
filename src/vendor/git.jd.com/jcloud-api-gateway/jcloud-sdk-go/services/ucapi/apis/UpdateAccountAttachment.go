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
    ucapi "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/ucapi/models"
)

type UpdateAccountAttachmentRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 用户pin  */
    Pin string `json:"pin"`

    /* 附属信息 (Optional) */
    Attachment *ucapi.UserAttachment `json:"attachment"`
}

/*
 * param regionId: Region ID (Required)
 * param pin: 用户pin (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateAccountAttachmentRequest(
    regionId string,
    pin string,
) *UpdateAccountAttachmentRequest {

	return &UpdateAccountAttachmentRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/user/{pin}:updateAccountAttachment",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Pin: pin,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pin: 用户pin (Required)
 * param attachment: 附属信息 (Optional)
 */
func NewUpdateAccountAttachmentRequestWithAllParams(
    regionId string,
    pin string,
    attachment *ucapi.UserAttachment,
) *UpdateAccountAttachmentRequest {

    return &UpdateAccountAttachmentRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/user/{pin}:updateAccountAttachment",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Pin: pin,
        Attachment: attachment,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateAccountAttachmentRequestWithoutParam() *UpdateAccountAttachmentRequest {

    return &UpdateAccountAttachmentRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/user/{pin}:updateAccountAttachment",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *UpdateAccountAttachmentRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pin: 用户pin(Required) */
func (r *UpdateAccountAttachmentRequest) SetPin(pin string) {
    r.Pin = pin
}

/* param attachment: 附属信息(Optional) */
func (r *UpdateAccountAttachmentRequest) SetAttachment(attachment *ucapi.UserAttachment) {
    r.Attachment = attachment
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateAccountAttachmentRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateAccountAttachmentResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateAccountAttachmentResult `json:"result"`
}

type UpdateAccountAttachmentResult struct {
}