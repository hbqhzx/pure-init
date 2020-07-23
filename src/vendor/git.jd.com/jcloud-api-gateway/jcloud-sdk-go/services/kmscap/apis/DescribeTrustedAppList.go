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
    kmscap "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/kmscap/models"
    common "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/common/models"
)

type DescribeTrustedAppListRequest struct {

    core.JDCloudRequest

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为10；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* appName - 根据应用名称检索
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeTrustedAppListRequest(
) *DescribeTrustedAppListRequest {

	return &DescribeTrustedAppListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/trusted",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为10；取值范围[10, 100] (Optional)
 * param filters: appName - 根据应用名称检索
 (Optional)
 */
func NewDescribeTrustedAppListRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
) *DescribeTrustedAppListRequest {

    return &DescribeTrustedAppListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/trusted",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeTrustedAppListRequestWithoutParam() *DescribeTrustedAppListRequest {

    return &DescribeTrustedAppListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/trusted",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeTrustedAppListRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为10；取值范围[10, 100](Optional) */
func (r *DescribeTrustedAppListRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: appName - 根据应用名称检索
(Optional) */
func (r *DescribeTrustedAppListRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeTrustedAppListRequest) GetRegionId() string {
    return ""
}

type DescribeTrustedAppListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeTrustedAppListResult `json:"result"`
}

type DescribeTrustedAppListResult struct {
    AppList []kmscap.AppItem `json:"appList"`
}