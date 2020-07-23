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

type QueryPushDomainORAppOrStreamRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* app名，传appName查询流名列表 (Optional) */
    App *string `json:"app"`

    /* 流名模糊查询 (Optional) */
    Stream *string `json:"stream"`

    /* 指定app/流名列表大小，默认50 (Optional) */
    Limit *int `json:"limit"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryPushDomainORAppOrStreamRequest(
    domain string,
) *QueryPushDomainORAppOrStreamRequest {

	return &QueryPushDomainORAppOrStreamRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/liveDomain/{domain}/stream:fuzzyQuery",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param app: app名，传appName查询流名列表 (Optional)
 * param stream: 流名模糊查询 (Optional)
 * param limit: 指定app/流名列表大小，默认50 (Optional)
 */
func NewQueryPushDomainORAppOrStreamRequestWithAllParams(
    domain string,
    app *string,
    stream *string,
    limit *int,
) *QueryPushDomainORAppOrStreamRequest {

    return &QueryPushDomainORAppOrStreamRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveDomain/{domain}/stream:fuzzyQuery",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        App: app,
        Stream: stream,
        Limit: limit,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryPushDomainORAppOrStreamRequestWithoutParam() *QueryPushDomainORAppOrStreamRequest {

    return &QueryPushDomainORAppOrStreamRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveDomain/{domain}/stream:fuzzyQuery",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *QueryPushDomainORAppOrStreamRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param app: app名，传appName查询流名列表(Optional) */
func (r *QueryPushDomainORAppOrStreamRequest) SetApp(app string) {
    r.App = &app
}

/* param stream: 流名模糊查询(Optional) */
func (r *QueryPushDomainORAppOrStreamRequest) SetStream(stream string) {
    r.Stream = &stream
}

/* param limit: 指定app/流名列表大小，默认50(Optional) */
func (r *QueryPushDomainORAppOrStreamRequest) SetLimit(limit int) {
    r.Limit = &limit
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryPushDomainORAppOrStreamRequest) GetRegionId() string {
    return ""
}

type QueryPushDomainORAppOrStreamResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryPushDomainORAppOrStreamResult `json:"result"`
}

type QueryPushDomainORAppOrStreamResult struct {
    Streams []string `json:"streams"`
}