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
    vod "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/vod/models"
)

type DescribeCategorysRequest struct {

    core.JDCloudRequest

    /* 列表结构返回 (Optional) */
    Fmt *string `json:"fmt"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeCategorysRequest(
) *DescribeCategorysRequest {

	return &DescribeCategorysRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/categorys",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param fmt: 列表结构返回 (Optional)
 */
func NewDescribeCategorysRequestWithAllParams(
    fmt *string,
) *DescribeCategorysRequest {

    return &DescribeCategorysRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/categorys",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Fmt: fmt,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeCategorysRequestWithoutParam() *DescribeCategorysRequest {

    return &DescribeCategorysRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/categorys",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param fmt: 列表结构返回(Optional) */
func (r *DescribeCategorysRequest) SetFmt(fmt string) {
    r.Fmt = &fmt
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeCategorysRequest) GetRegionId() string {
    return ""
}

type DescribeCategorysResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeCategorysResult `json:"result"`
}

type DescribeCategorysResult struct {
    Categorys []vod.CategoryTree `json:"categorys"`
}