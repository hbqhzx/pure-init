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

type GetSubDomainsRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 域名ID，请使用getDomains接口获取。  */
    DomainId string `json:"domainId"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用getDomains接口获取。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetSubDomainsRequest(
    regionId string,
    domainId string,
) *GetSubDomainsRequest {

	return &GetSubDomainsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/domain/{domainId}/getSubDomains",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DomainId: domainId,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID，请使用getDomains接口获取。 (Required)
 */
func NewGetSubDomainsRequestWithAllParams(
    regionId string,
    domainId string,
) *GetSubDomainsRequest {

    return &GetSubDomainsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/getSubDomains",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DomainId: domainId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetSubDomainsRequestWithoutParam() *GetSubDomainsRequest {

    return &GetSubDomainsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/getSubDomains",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *GetSubDomainsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param domainId: 域名ID，请使用getDomains接口获取。(Required) */
func (r *GetSubDomainsRequest) SetDomainId(domainId string) {
    r.DomainId = domainId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetSubDomainsRequest) GetRegionId() string {
    return r.RegionId
}

type GetSubDomainsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetSubDomainsResult `json:"result"`
}

type GetSubDomainsResult struct {
    Data []string `json:"data"`
}