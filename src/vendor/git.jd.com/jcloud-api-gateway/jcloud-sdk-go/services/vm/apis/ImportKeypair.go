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

type ImportKeypairRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 密钥对名称，需要全局唯一。只允许数字、大小写字母、下划线“_”及中划线“-”，不超过32个字符。
  */
    KeyName string `json:"keyName"`

    /* 密钥对的公钥部分  */
    PublicKey string `json:"publicKey"`
}

/*
 * param regionId: 地域ID (Required)
 * param keyName: 密钥对名称，需要全局唯一。只允许数字、大小写字母、下划线“_”及中划线“-”，不超过32个字符。
 (Required)
 * param publicKey: 密钥对的公钥部分 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewImportKeypairRequest(
    regionId string,
    keyName string,
    publicKey string,
) *ImportKeypairRequest {

	return &ImportKeypairRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/keypairs:import",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        KeyName: keyName,
        PublicKey: publicKey,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param keyName: 密钥对名称，需要全局唯一。只允许数字、大小写字母、下划线“_”及中划线“-”，不超过32个字符。
 (Required)
 * param publicKey: 密钥对的公钥部分 (Required)
 */
func NewImportKeypairRequestWithAllParams(
    regionId string,
    keyName string,
    publicKey string,
) *ImportKeypairRequest {

    return &ImportKeypairRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/keypairs:import",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        KeyName: keyName,
        PublicKey: publicKey,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewImportKeypairRequestWithoutParam() *ImportKeypairRequest {

    return &ImportKeypairRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/keypairs:import",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *ImportKeypairRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param keyName: 密钥对名称，需要全局唯一。只允许数字、大小写字母、下划线“_”及中划线“-”，不超过32个字符。
(Required) */
func (r *ImportKeypairRequest) SetKeyName(keyName string) {
    r.KeyName = keyName
}

/* param publicKey: 密钥对的公钥部分(Required) */
func (r *ImportKeypairRequest) SetPublicKey(publicKey string) {
    r.PublicKey = publicKey
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ImportKeypairRequest) GetRegionId() string {
    return r.RegionId
}

type ImportKeypairResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ImportKeypairResult `json:"result"`
}

type ImportKeypairResult struct {
    KeyName string `json:"keyName"`
    KeyFingerprint string `json:"keyFingerprint"`
}