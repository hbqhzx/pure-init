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

type DeleteCdnDomainHearderRequest struct {

    core.JDCloudRequest

    /* 域名  */
    DomainName string `json:"domainName"`

    /* 域名头编号  */
    HeaderId int `json:"headerId"`
}

/*
 * param domainName: 域名 (Required)
 * param headerId: 域名头编号 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteCdnDomainHearderRequest(
    domainName string,
    headerId int,
) *DeleteCdnDomainHearderRequest {

	return &DeleteCdnDomainHearderRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domains/{domainName}/cdns/{headerId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        DomainName: domainName,
        HeaderId: headerId,
	}
}

/*
 * param domainName: 域名 (Required)
 * param headerId: 域名头编号 (Required)
 */
func NewDeleteCdnDomainHearderRequestWithAllParams(
    domainName string,
    headerId int,
) *DeleteCdnDomainHearderRequest {

    return &DeleteCdnDomainHearderRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains/{domainName}/cdns/{headerId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        DomainName: domainName,
        HeaderId: headerId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteCdnDomainHearderRequestWithoutParam() *DeleteCdnDomainHearderRequest {

    return &DeleteCdnDomainHearderRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains/{domainName}/cdns/{headerId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domainName: 域名(Required) */
func (r *DeleteCdnDomainHearderRequest) SetDomainName(domainName string) {
    r.DomainName = domainName
}

/* param headerId: 域名头编号(Required) */
func (r *DeleteCdnDomainHearderRequest) SetHeaderId(headerId int) {
    r.HeaderId = headerId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteCdnDomainHearderRequest) GetRegionId() string {
    return ""
}

type DeleteCdnDomainHearderResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteCdnDomainHearderResult `json:"result"`
}

type DeleteCdnDomainHearderResult struct {
    HeaderId int64 `json:"headerId"`
}