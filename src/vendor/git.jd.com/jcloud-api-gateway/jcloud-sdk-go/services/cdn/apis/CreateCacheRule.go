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

type CreateCacheRuleRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* 此条配置的权重值, 取值范围为1-10,1最大 (Optional) */
    Weight *int `json:"weight"`

    /* 缓存时间,单位秒 (Optional) */
    Ttl *int64 `json:"ttl"`

    /* 规则内容。其他类型只能以/或者.开头，如/a/b或.jpg (Optional) */
    Contents *string `json:"contents"`

    /* 缓存方式：0、不缓存，1自定义 (Optional) */
    CacheType *int `json:"cacheType"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateCacheRuleRequest(
    domain string,
) *CreateCacheRuleRequest {

	return &CreateCacheRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/cacheRule",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param weight: 此条配置的权重值, 取值范围为1-10,1最大 (Optional)
 * param ttl: 缓存时间,单位秒 (Optional)
 * param contents: 规则内容。其他类型只能以/或者.开头，如/a/b或.jpg (Optional)
 * param cacheType: 缓存方式：0、不缓存，1自定义 (Optional)
 */
func NewCreateCacheRuleRequestWithAllParams(
    domain string,
    weight *int,
    ttl *int64,
    contents *string,
    cacheType *int,
) *CreateCacheRuleRequest {

    return &CreateCacheRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/cacheRule",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        Weight: weight,
        Ttl: ttl,
        Contents: contents,
        CacheType: cacheType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateCacheRuleRequestWithoutParam() *CreateCacheRuleRequest {

    return &CreateCacheRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/cacheRule",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *CreateCacheRuleRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param weight: 此条配置的权重值, 取值范围为1-10,1最大(Optional) */
func (r *CreateCacheRuleRequest) SetWeight(weight int) {
    r.Weight = &weight
}

/* param ttl: 缓存时间,单位秒(Optional) */
func (r *CreateCacheRuleRequest) SetTtl(ttl int64) {
    r.Ttl = &ttl
}

/* param contents: 规则内容。其他类型只能以/或者.开头，如/a/b或.jpg(Optional) */
func (r *CreateCacheRuleRequest) SetContents(contents string) {
    r.Contents = &contents
}

/* param cacheType: 缓存方式：0、不缓存，1自定义(Optional) */
func (r *CreateCacheRuleRequest) SetCacheType(cacheType int) {
    r.CacheType = &cacheType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateCacheRuleRequest) GetRegionId() string {
    return ""
}

type CreateCacheRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateCacheRuleResult `json:"result"`
}

type CreateCacheRuleResult struct {
    ConfigId int64 `json:"configId"`
}