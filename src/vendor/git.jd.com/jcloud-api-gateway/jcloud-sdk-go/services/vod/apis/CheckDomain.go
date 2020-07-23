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

type CheckDomainRequest struct {

    core.JDCloudRequest

    /* 域名  */
    DomainName string `json:"domainName"`
}

/*
 * param domainName: 域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCheckDomainRequest(
    domainName string,
) *CheckDomainRequest {

	return &CheckDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domains/{domainName}:check",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        DomainName: domainName,
	}
}

/*
 * param domainName: 域名 (Required)
 */
func NewCheckDomainRequestWithAllParams(
    domainName string,
) *CheckDomainRequest {

    return &CheckDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains/{domainName}:check",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        DomainName: domainName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCheckDomainRequestWithoutParam() *CheckDomainRequest {

    return &CheckDomainRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains/{domainName}:check",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domainName: 域名(Required) */
func (r *CheckDomainRequest) SetDomainName(domainName string) {
    r.DomainName = domainName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CheckDomainRequest) GetRegionId() string {
    return ""
}

type CheckDomainResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CheckDomainResult `json:"result"`
}

type CheckDomainResult struct {
    CheckStatus bool `json:"checkStatus"`
}