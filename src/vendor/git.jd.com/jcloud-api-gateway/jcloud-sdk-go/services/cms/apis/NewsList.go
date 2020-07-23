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
    cms "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/cms/models"
)

type NewsListRequest struct {

    core.JDCloudRequest

    /* pageSize (Optional) */
    PageSize *int `json:"pageSize"`

    /* pageNum (Optional) */
    PageNum *int `json:"pageNum"`

    /* hot (Optional) */
    Hot *int `json:"hot"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewNewsListRequest(
) *NewsListRequest {

	return &NewsListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/portal/api/news/list",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageSize: pageSize (Optional)
 * param pageNum: pageNum (Optional)
 * param hot: hot (Optional)
 */
func NewNewsListRequestWithAllParams(
    pageSize *int,
    pageNum *int,
    hot *int,
) *NewsListRequest {

    return &NewsListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/portal/api/news/list",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageSize: pageSize,
        PageNum: pageNum,
        Hot: hot,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewNewsListRequestWithoutParam() *NewsListRequest {

    return &NewsListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/portal/api/news/list",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageSize: pageSize(Optional) */
func (r *NewsListRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param pageNum: pageNum(Optional) */
func (r *NewsListRequest) SetPageNum(pageNum int) {
    r.PageNum = &pageNum
}

/* param hot: hot(Optional) */
func (r *NewsListRequest) SetHot(hot int) {
    r.Hot = &hot
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r NewsListRequest) GetRegionId() string {
    return ""
}

type NewsListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result NewsListResult `json:"result"`
}

type NewsListResult struct {
    PageInfo cms.PageInfo `json:"pageInfo"`
    List []cms.News `json:"list"`
}