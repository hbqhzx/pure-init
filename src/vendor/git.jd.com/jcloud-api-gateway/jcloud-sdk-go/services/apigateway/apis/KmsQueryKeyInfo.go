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
    apigateway "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/apigateway/models"
)

type KmsQueryKeyInfoRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* keyId  */
    KeyId string `json:"keyId"`
}

/*
 * param regionId: 地域ID (Required)
 * param keyId: keyId (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewKmsQueryKeyInfoRequest(
    regionId string,
    keyId string,
) *KmsQueryKeyInfoRequest {

	return &KmsQueryKeyInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/kms/keyinfo",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        KeyId: keyId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param keyId: keyId (Required)
 */
func NewKmsQueryKeyInfoRequestWithAllParams(
    regionId string,
    keyId string,
) *KmsQueryKeyInfoRequest {

    return &KmsQueryKeyInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/kms/keyinfo",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        KeyId: keyId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewKmsQueryKeyInfoRequestWithoutParam() *KmsQueryKeyInfoRequest {

    return &KmsQueryKeyInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/kms/keyinfo",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *KmsQueryKeyInfoRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param keyId: keyId(Required) */
func (r *KmsQueryKeyInfoRequest) SetKeyId(keyId string) {
    r.KeyId = keyId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r KmsQueryKeyInfoRequest) GetRegionId() string {
    return r.RegionId
}

type KmsQueryKeyInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result KmsQueryKeyInfoResult `json:"result"`
}

type KmsQueryKeyInfoResult struct {
    KeyInfo apigateway.KeyInfo `json:"keyInfo"`
}