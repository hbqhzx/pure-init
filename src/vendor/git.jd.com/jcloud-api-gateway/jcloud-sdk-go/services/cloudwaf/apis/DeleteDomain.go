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
    cloudwaf "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/cloudwaf/models"
)

type DeleteDomainRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 域名  */
    Domain string `json:"domain"`

    /* 请求  */
    Req *cloudwaf.CommonReq `json:"req"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domain: 域名 (Required)
 * param req: 请求 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteDomainRequest(
    regionId string,
    domain string,
    req *cloudwaf.CommonReq,
) *DeleteDomainRequest {

	return &DeleteDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/domain/{domain}/delete",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Domain: domain,
        Req: req,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domain: 域名 (Required)
 * param req: 请求 (Required)
 */
func NewDeleteDomainRequestWithAllParams(
    regionId string,
    domain string,
    req *cloudwaf.CommonReq,
) *DeleteDomainRequest {

    return &DeleteDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domain}/delete",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Domain: domain,
        Req: req,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteDomainRequestWithoutParam() *DeleteDomainRequest {

    return &DeleteDomainRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domain}/delete",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *DeleteDomainRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param domain: 域名(Required) */
func (r *DeleteDomainRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param req: 请求(Required) */
func (r *DeleteDomainRequest) SetReq(req *cloudwaf.CommonReq) {
    r.Req = req
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteDomainRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteDomainResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteDomainResult `json:"result"`
}

type DeleteDomainResult struct {
}