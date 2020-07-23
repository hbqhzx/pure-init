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
    common "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/common/models"
)

type DescribeNotifyConfigsRequest struct {

    core.JDCloudRequest

    /* 页码；默认为1；取值范围[1, 100000] (Optional) */
    PageNum *int `json:"pageNum"`

    /* 分页大小；默认为10；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 转码列表查询过滤条件, 不传递分页参数时默认返回10条 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeNotifyConfigsRequest(
) *DescribeNotifyConfigsRequest {

	return &DescribeNotifyConfigsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/notifyConfigs",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNum: 页码；默认为1；取值范围[1, 100000] (Optional)
 * param pageSize: 分页大小；默认为10；取值范围[10, 100] (Optional)
 * param filters: 转码列表查询过滤条件, 不传递分页参数时默认返回10条 (Optional)
 */
func NewDescribeNotifyConfigsRequestWithAllParams(
    pageNum *int,
    pageSize *int,
    filters []common.Filter,
) *DescribeNotifyConfigsRequest {

    return &DescribeNotifyConfigsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/notifyConfigs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNum: pageNum,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeNotifyConfigsRequestWithoutParam() *DescribeNotifyConfigsRequest {

    return &DescribeNotifyConfigsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/notifyConfigs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNum: 页码；默认为1；取值范围[1, 100000](Optional) */
func (r *DescribeNotifyConfigsRequest) SetPageNum(pageNum int) {
    r.PageNum = &pageNum
}

/* param pageSize: 分页大小；默认为10；取值范围[10, 100](Optional) */
func (r *DescribeNotifyConfigsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: 转码列表查询过滤条件, 不传递分页参数时默认返回10条(Optional) */
func (r *DescribeNotifyConfigsRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeNotifyConfigsRequest) GetRegionId() string {
    return ""
}

type DescribeNotifyConfigsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeNotifyConfigsResult `json:"result"`
}

type DescribeNotifyConfigsResult struct {
    PageNum int `json:"pageNum"`
    PageSize int `json:"pageSize"`
    Total int `json:"total"`
    Pages int `json:"pages"`
    NotifyConfigs []vod.NotifyConfig `json:"notifyConfigs"`
}