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
    directoryservice "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/directoryservice/models"
)

type DescribeTrustDomainsRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 资源id  */
    DirectoryId string `json:"directoryId"`
}

/*
 * param regionId: 地域ID (Required)
 * param directoryId: 资源id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeTrustDomainsRequest(
    regionId string,
    directoryId string,
) *DescribeTrustDomainsRequest {

	return &DescribeTrustDomainsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/directory/{directoryId}/trust:describeTrustDomains",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DirectoryId: directoryId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param directoryId: 资源id (Required)
 */
func NewDescribeTrustDomainsRequestWithAllParams(
    regionId string,
    directoryId string,
) *DescribeTrustDomainsRequest {

    return &DescribeTrustDomainsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/directory/{directoryId}/trust:describeTrustDomains",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DirectoryId: directoryId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeTrustDomainsRequestWithoutParam() *DescribeTrustDomainsRequest {

    return &DescribeTrustDomainsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/directory/{directoryId}/trust:describeTrustDomains",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeTrustDomainsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param directoryId: 资源id(Required) */
func (r *DescribeTrustDomainsRequest) SetDirectoryId(directoryId string) {
    r.DirectoryId = directoryId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeTrustDomainsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeTrustDomainsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeTrustDomainsResult `json:"result"`
}

type DescribeTrustDomainsResult struct {
    TrustDomains []directoryservice.AdTrustVo `json:"trustDomains"`
}