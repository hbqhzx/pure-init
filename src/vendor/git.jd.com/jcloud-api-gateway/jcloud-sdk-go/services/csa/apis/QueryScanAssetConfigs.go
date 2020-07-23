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
    csa "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/csa/models"
)

type QueryScanAssetConfigsRequest struct {

    core.JDCloudRequest

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为10；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 资产名称查询关键字 (Optional) */
    Keyword *string `json:"keyword"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryScanAssetConfigsRequest(
) *QueryScanAssetConfigsRequest {

	return &QueryScanAssetConfigsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/scanAssetConfigs",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为10；取值范围[10, 100] (Optional)
 * param keyword: 资产名称查询关键字 (Optional)
 */
func NewQueryScanAssetConfigsRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    keyword *string,
) *QueryScanAssetConfigsRequest {

    return &QueryScanAssetConfigsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/scanAssetConfigs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
        Keyword: keyword,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryScanAssetConfigsRequestWithoutParam() *QueryScanAssetConfigsRequest {

    return &QueryScanAssetConfigsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/scanAssetConfigs",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *QueryScanAssetConfigsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为10；取值范围[10, 100](Optional) */
func (r *QueryScanAssetConfigsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param keyword: 资产名称查询关键字(Optional) */
func (r *QueryScanAssetConfigsRequest) SetKeyword(keyword string) {
    r.Keyword = &keyword
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryScanAssetConfigsRequest) GetRegionId() string {
    return ""
}

type QueryScanAssetConfigsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryScanAssetConfigsResult `json:"result"`
}

type QueryScanAssetConfigsResult struct {
    ScanAssetConfigs []csa.ScanAssetConfig `json:"scanAssetConfigs"`
    Total int `json:"total"`
}