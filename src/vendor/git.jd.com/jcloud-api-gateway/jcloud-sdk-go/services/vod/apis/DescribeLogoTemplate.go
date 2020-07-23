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

type DescribeLogoTemplateRequest struct {

    core.JDCloudRequest

    /* 模板ID  */
    LogoId int `json:"logoId"`
}

/*
 * param logoId: 模板ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLogoTemplateRequest(
    logoId int,
) *DescribeLogoTemplateRequest {

	return &DescribeLogoTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/logos/{logoId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        LogoId: logoId,
	}
}

/*
 * param logoId: 模板ID (Required)
 */
func NewDescribeLogoTemplateRequestWithAllParams(
    logoId int,
) *DescribeLogoTemplateRequest {

    return &DescribeLogoTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/logos/{logoId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        LogoId: logoId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLogoTemplateRequestWithoutParam() *DescribeLogoTemplateRequest {

    return &DescribeLogoTemplateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/logos/{logoId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param logoId: 模板ID(Required) */
func (r *DescribeLogoTemplateRequest) SetLogoId(logoId int) {
    r.LogoId = logoId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLogoTemplateRequest) GetRegionId() string {
    return ""
}

type DescribeLogoTemplateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLogoTemplateResult `json:"result"`
}

type DescribeLogoTemplateResult struct {
    LogoId int64 `json:"logoId"`
    Name string `json:"name"`
    Position int `json:"position"`
    Width int `json:"width"`
    Height int `json:"height"`
    Unit string `json:"unit"`
    OffsetX int `json:"offsetX"`
    OffsetY int `json:"offsetY"`
    UpdateTime string `json:"updateTime"`
}