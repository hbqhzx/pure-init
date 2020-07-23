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
    cdn "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/cdn/models"
)

type QueryDomainStreamHistoryRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* pageNumber (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* pageSize (Optional) */
    PageSize *int `json:"pageSize"`

    /* 禁播类型，不传查询全部禁播类型 (Optional) */
    Type *string `json:"type"`

    /* app名，传空代表查询全部 (Optional) */
    App *string `json:"app"`

    /* 流名,传空代表查询全部 (Optional) */
    Stream *string `json:"stream"`

    /* 开始禁播时间，格式：yyyy-MM-dd HH:mm  */
    StartTime string `json:"startTime"`

    /* 结束禁播时间，格式：yyyy-MM-dd HH:mm  */
    EndTime string `json:"endTime"`
}

/*
 * param domain: 用户域名 (Required)
 * param startTime: 开始禁播时间，格式：yyyy-MM-dd HH:mm (Required)
 * param endTime: 结束禁播时间，格式：yyyy-MM-dd HH:mm (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryDomainStreamHistoryRequest(
    domain string,
    startTime string,
    endTime string,
) *QueryDomainStreamHistoryRequest {

	return &QueryDomainStreamHistoryRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/liveDomain/{domain}/stream:forbiddenHistory",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
        StartTime: startTime,
        EndTime: endTime,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param pageNumber: pageNumber (Optional)
 * param pageSize: pageSize (Optional)
 * param type_: 禁播类型，不传查询全部禁播类型 (Optional)
 * param app: app名，传空代表查询全部 (Optional)
 * param stream: 流名,传空代表查询全部 (Optional)
 * param startTime: 开始禁播时间，格式：yyyy-MM-dd HH:mm (Required)
 * param endTime: 结束禁播时间，格式：yyyy-MM-dd HH:mm (Required)
 */
func NewQueryDomainStreamHistoryRequestWithAllParams(
    domain string,
    pageNumber *int,
    pageSize *int,
    type_ *string,
    app *string,
    stream *string,
    startTime string,
    endTime string,
) *QueryDomainStreamHistoryRequest {

    return &QueryDomainStreamHistoryRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveDomain/{domain}/stream:forbiddenHistory",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Type: type_,
        App: app,
        Stream: stream,
        StartTime: startTime,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryDomainStreamHistoryRequestWithoutParam() *QueryDomainStreamHistoryRequest {

    return &QueryDomainStreamHistoryRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveDomain/{domain}/stream:forbiddenHistory",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *QueryDomainStreamHistoryRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param pageNumber: pageNumber(Optional) */
func (r *QueryDomainStreamHistoryRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: pageSize(Optional) */
func (r *QueryDomainStreamHistoryRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param type_: 禁播类型，不传查询全部禁播类型(Optional) */
func (r *QueryDomainStreamHistoryRequest) SetType(type_ string) {
    r.Type = &type_
}

/* param app: app名，传空代表查询全部(Optional) */
func (r *QueryDomainStreamHistoryRequest) SetApp(app string) {
    r.App = &app
}

/* param stream: 流名,传空代表查询全部(Optional) */
func (r *QueryDomainStreamHistoryRequest) SetStream(stream string) {
    r.Stream = &stream
}

/* param startTime: 开始禁播时间，格式：yyyy-MM-dd HH:mm(Required) */
func (r *QueryDomainStreamHistoryRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 结束禁播时间，格式：yyyy-MM-dd HH:mm(Required) */
func (r *QueryDomainStreamHistoryRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryDomainStreamHistoryRequest) GetRegionId() string {
    return ""
}

type QueryDomainStreamHistoryResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryDomainStreamHistoryResult `json:"result"`
}

type QueryDomainStreamHistoryResult struct {
    Total int `json:"total"`
    StreamList []cdn.ForbiddenStreamHistoryItem `json:"streamList"`
}