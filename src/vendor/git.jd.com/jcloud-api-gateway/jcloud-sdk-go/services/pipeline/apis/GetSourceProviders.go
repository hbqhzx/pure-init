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
    pipeline "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/pipeline/models"
)

type GetSourceProvidersRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetSourceProvidersRequest(
    regionId string,
) *GetSourceProvidersRequest {

	return &GetSourceProvidersRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/options/sourceProviders",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 */
func NewGetSourceProvidersRequestWithAllParams(
    regionId string,
) *GetSourceProvidersRequest {

    return &GetSourceProvidersRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/options/sourceProviders",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetSourceProvidersRequestWithoutParam() *GetSourceProvidersRequest {

    return &GetSourceProvidersRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/options/sourceProviders",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *GetSourceProvidersRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetSourceProvidersRequest) GetRegionId() string {
    return r.RegionId
}

type GetSourceProvidersResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetSourceProvidersResult `json:"result"`
}

type GetSourceProvidersResult struct {
    TotalCount int `json:"totalCount"`
    Providers []pipeline.NameLabelPair `json:"providers"`
}