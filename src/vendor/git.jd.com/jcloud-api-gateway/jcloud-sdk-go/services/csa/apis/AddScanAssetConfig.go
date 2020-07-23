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
    csa "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/csa/models"
)

type AddScanAssetConfigRequest struct {

    core.JDCloudRequest

    /* 域名  */
    Domain string `json:"domain"`

    /* Cookie值 (Optional) */
    Cookie *string `json:"cookie"`
}

/*
 * param domain: 域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddScanAssetConfigRequest(
    domain string,
) *AddScanAssetConfigRequest {

	return &AddScanAssetConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/scanAssetConfigs",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 域名 (Required)
 * param cookie: Cookie值 (Optional)
 */
func NewAddScanAssetConfigRequestWithAllParams(
    domain string,
    cookie *string,
) *AddScanAssetConfigRequest {

    return &AddScanAssetConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/scanAssetConfigs",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        Cookie: cookie,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddScanAssetConfigRequestWithoutParam() *AddScanAssetConfigRequest {

    return &AddScanAssetConfigRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/scanAssetConfigs",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param domain: 域名(Required) */
func (r *AddScanAssetConfigRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param cookie: Cookie值(Optional) */
func (r *AddScanAssetConfigRequest) SetCookie(cookie string) {
    r.Cookie = &cookie
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddScanAssetConfigRequest) GetRegionId() string {
    return ""
}

type AddScanAssetConfigResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddScanAssetConfigResult `json:"result"`
}

type AddScanAssetConfigResult struct {
    ScanAssetConfig csa.ScanAssetConfig `json:"scanAssetConfig"`
    Message string `json:"message"`
}