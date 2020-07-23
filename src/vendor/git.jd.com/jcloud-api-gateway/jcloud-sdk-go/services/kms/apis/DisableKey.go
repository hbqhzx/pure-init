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

type DisableKeyRequest struct {

    core.JDCloudRequest

    /* 密钥ID  */
    KeyId string `json:"keyId"`
}

/*
 * param keyId: 密钥ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDisableKeyRequest(
    keyId string,
) *DisableKeyRequest {

	return &DisableKeyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/key/{keyId}:disable",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        KeyId: keyId,
	}
}

/*
 * param keyId: 密钥ID (Required)
 */
func NewDisableKeyRequestWithAllParams(
    keyId string,
) *DisableKeyRequest {

    return &DisableKeyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/key/{keyId}:disable",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        KeyId: keyId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDisableKeyRequestWithoutParam() *DisableKeyRequest {

    return &DisableKeyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/key/{keyId}:disable",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param keyId: 密钥ID(Required) */
func (r *DisableKeyRequest) SetKeyId(keyId string) {
    r.KeyId = keyId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DisableKeyRequest) GetRegionId() string {
    return ""
}

type DisableKeyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DisableKeyResult `json:"result"`
}

type DisableKeyResult struct {
}