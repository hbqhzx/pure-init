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
    ompopenapi "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/ompopenapi/models"
    common "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/common/models"
)

type QueryLogRequest struct {

    core.JDCloudRequest

    /* 页码, 默认为1, 取值范围：[1,∞) (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20，取值范围：[10,100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* startTime - 开始时间，单位毫秒，长度13位，必填。
endTime - 结束时间，单位毫秒，长度13位，必填。
searchField - 查询字段：req_id，uuid,erp,pin,path，response_status,operationId
searchValue - 查询值，单个值，模糊匹配
method - 请求方式
orderby - 排序
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryLogRequest(
) *QueryLogRequest {

	return &QueryLogRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/logs",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNumber: 页码, 默认为1, 取值范围：[1,∞) (Optional)
 * param pageSize: 分页大小，默认为20，取值范围：[10,100] (Optional)
 * param filters: startTime - 开始时间，单位毫秒，长度13位，必填。
endTime - 结束时间，单位毫秒，长度13位，必填。
searchField - 查询字段：req_id，uuid,erp,pin,path，response_status,operationId
searchValue - 查询值，单个值，模糊匹配
method - 请求方式
orderby - 排序
 (Optional)
 */
func NewQueryLogRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
) *QueryLogRequest {

    return &QueryLogRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/logs",
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
func NewQueryLogRequestWithoutParam() *QueryLogRequest {

    return &QueryLogRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/logs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNumber: 页码, 默认为1, 取值范围：[1,∞)(Optional) */
func (r *QueryLogRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小，默认为20，取值范围：[10,100](Optional) */
func (r *QueryLogRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: startTime - 开始时间，单位毫秒，长度13位，必填。
endTime - 结束时间，单位毫秒，长度13位，必填。
searchField - 查询字段：req_id，uuid,erp,pin,path，response_status,operationId
searchValue - 查询值，单个值，模糊匹配
method - 请求方式
orderby - 排序
(Optional) */
func (r *QueryLogRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryLogRequest) GetRegionId() string {
    return ""
}

type QueryLogResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryLogResult `json:"result"`
}

type QueryLogResult struct {
    LogDetail []ompopenapi.LogDetail `json:"logDetail"`
    PageInfo ompopenapi.PageInfo `json:"pageInfo"`
}