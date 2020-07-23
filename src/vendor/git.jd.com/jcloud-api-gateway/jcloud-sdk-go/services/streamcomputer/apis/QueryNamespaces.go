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
    streamcomputer "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/streamcomputer/models"
)

type QueryNamespacesRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /*  (Optional) */
    Keyword *string `json:"keyword"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryNamespacesRequest(
    regionId string,
) *QueryNamespacesRequest {

	return &QueryNamespacesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/namespaces",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param keyword:  (Optional)
 */
func NewQueryNamespacesRequestWithAllParams(
    regionId string,
    keyword *string,
) *QueryNamespacesRequest {

    return &QueryNamespacesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/namespaces",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Keyword: keyword,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryNamespacesRequestWithoutParam() *QueryNamespacesRequest {

    return &QueryNamespacesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/namespaces",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *QueryNamespacesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param keyword: (Optional) */
func (r *QueryNamespacesRequest) SetKeyword(keyword string) {
    r.Keyword = &keyword
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryNamespacesRequest) GetRegionId() string {
    return r.RegionId
}

type QueryNamespacesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryNamespacesResult `json:"result"`
}

type QueryNamespacesResult struct {
    Namespaces []streamcomputer.Namespace `json:"namespaces"`
}