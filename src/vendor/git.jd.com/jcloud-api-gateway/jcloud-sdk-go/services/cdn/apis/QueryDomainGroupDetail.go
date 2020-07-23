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

type QueryDomainGroupDetailRequest struct {

    core.JDCloudRequest

    /* 域名组id  */
    Id int `json:"id"`
}

/*
 * param id: 域名组id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryDomainGroupDetailRequest(
    id int,
) *QueryDomainGroupDetailRequest {

	return &QueryDomainGroupDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domainGroup/{id}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Id: id,
	}
}

/*
 * param id: 域名组id (Required)
 */
func NewQueryDomainGroupDetailRequestWithAllParams(
    id int,
) *QueryDomainGroupDetailRequest {

    return &QueryDomainGroupDetailRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domainGroup/{id}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Id: id,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryDomainGroupDetailRequestWithoutParam() *QueryDomainGroupDetailRequest {

    return &QueryDomainGroupDetailRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domainGroup/{id}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param id: 域名组id(Required) */
func (r *QueryDomainGroupDetailRequest) SetId(id int) {
    r.Id = id
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryDomainGroupDetailRequest) GetRegionId() string {
    return ""
}

type QueryDomainGroupDetailResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryDomainGroupDetailResult `json:"result"`
}

type QueryDomainGroupDetailResult struct {
    Domains []string `json:"domains"`
    PrimaryDomain string `json:"primaryDomain"`
    ShareCache string `json:"shareCache"`
    DomainGroupName string `json:"domainGroupName"`
}