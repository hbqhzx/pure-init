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

type DescribeLivePlayAuthKeyRequest struct {

    core.JDCloudRequest

    /* 您的播放加速域名  */
    PlayDomain string `json:"playDomain"`
}

/*
 * param playDomain: 您的播放加速域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLivePlayAuthKeyRequest(
    playDomain string,
) *DescribeLivePlayAuthKeyRequest {

	return &DescribeLivePlayAuthKeyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/livePlayAuthKey",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        PlayDomain: playDomain,
	}
}

/*
 * param playDomain: 您的播放加速域名 (Required)
 */
func NewDescribeLivePlayAuthKeyRequestWithAllParams(
    playDomain string,
) *DescribeLivePlayAuthKeyRequest {

    return &DescribeLivePlayAuthKeyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/livePlayAuthKey",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PlayDomain: playDomain,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLivePlayAuthKeyRequestWithoutParam() *DescribeLivePlayAuthKeyRequest {

    return &DescribeLivePlayAuthKeyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/livePlayAuthKey",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param playDomain: 您的播放加速域名(Required) */
func (r *DescribeLivePlayAuthKeyRequest) SetPlayDomain(playDomain string) {
    r.PlayDomain = playDomain
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLivePlayAuthKeyRequest) GetRegionId() string {
    return ""
}

type DescribeLivePlayAuthKeyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLivePlayAuthKeyResult `json:"result"`
}

type DescribeLivePlayAuthKeyResult struct {
    PlayDomain string `json:"playDomain"`
    AuthStatus string `json:"authStatus"`
    AuthKey string `json:"authKey"`
}