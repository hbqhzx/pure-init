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

type BatchCreateRequest struct {

    core.JDCloudRequest

    /*  (Optional) */
    Domains []string `json:"domains"`

    /* 回源类型只能是[ips,domain,oss]中的一种 (Optional) */
    SourceType *string `json:"sourceType"`

    /* 点播域名的类型只能是[vod,download,web]中的一种 (Optional) */
    CdnType *string `json:"cdnType"`

    /* 回源方式,只能是[https,http]中的一种,默认http (Optional) */
    BackSourceType *string `json:"backSourceType"`

    /* 日带宽(Mbps) (Optional) */
    DailyBandWidth *int64 `json:"dailyBandWidth"`

    /* 服务质量,只能是[good,general]中的一种,默认为good (Optional) */
    Quaility *string `json:"quaility"`

    /*  (Optional) */
    MaxFileSize *int64 `json:"maxFileSize"`

    /*  (Optional) */
    MinFileSize *int64 `json:"minFileSize"`

    /*  (Optional) */
    SumFileSize *int64 `json:"sumFileSize"`

    /*  (Optional) */
    AvgFileSize *int64 `json:"avgFileSize"`

    /*  (Optional) */
    DefaultSourceHost *string `json:"defaultSourceHost"`

    /*  (Optional) */
    HttpType *string `json:"httpType"`

    /*  (Optional) */
    IpSource []cdn.IpSourceInfo `json:"ipSource"`

    /*  (Optional) */
    DomainSource []cdn.DomainSourceInfo `json:"domainSource"`

    /*  (Optional) */
    OssSource *string `json:"ossSource"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewBatchCreateRequest(
) *BatchCreateRequest {

	return &BatchCreateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain:batchCreate",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param domains:  (Optional)
 * param sourceType: 回源类型只能是[ips,domain,oss]中的一种 (Optional)
 * param cdnType: 点播域名的类型只能是[vod,download,web]中的一种 (Optional)
 * param backSourceType: 回源方式,只能是[https,http]中的一种,默认http (Optional)
 * param dailyBandWidth: 日带宽(Mbps) (Optional)
 * param quaility: 服务质量,只能是[good,general]中的一种,默认为good (Optional)
 * param maxFileSize:  (Optional)
 * param minFileSize:  (Optional)
 * param sumFileSize:  (Optional)
 * param avgFileSize:  (Optional)
 * param defaultSourceHost:  (Optional)
 * param httpType:  (Optional)
 * param ipSource:  (Optional)
 * param domainSource:  (Optional)
 * param ossSource:  (Optional)
 */
func NewBatchCreateRequestWithAllParams(
    domains []string,
    sourceType *string,
    cdnType *string,
    backSourceType *string,
    dailyBandWidth *int64,
    quaility *string,
    maxFileSize *int64,
    minFileSize *int64,
    sumFileSize *int64,
    avgFileSize *int64,
    defaultSourceHost *string,
    httpType *string,
    ipSource []cdn.IpSourceInfo,
    domainSource []cdn.DomainSourceInfo,
    ossSource *string,
) *BatchCreateRequest {

    return &BatchCreateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain:batchCreate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domains: domains,
        SourceType: sourceType,
        CdnType: cdnType,
        BackSourceType: backSourceType,
        DailyBandWidth: dailyBandWidth,
        Quaility: quaility,
        MaxFileSize: maxFileSize,
        MinFileSize: minFileSize,
        SumFileSize: sumFileSize,
        AvgFileSize: avgFileSize,
        DefaultSourceHost: defaultSourceHost,
        HttpType: httpType,
        IpSource: ipSource,
        DomainSource: domainSource,
        OssSource: ossSource,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewBatchCreateRequestWithoutParam() *BatchCreateRequest {

    return &BatchCreateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain:batchCreate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domains: (Optional) */
func (r *BatchCreateRequest) SetDomains(domains []string) {
    r.Domains = domains
}

/* param sourceType: 回源类型只能是[ips,domain,oss]中的一种(Optional) */
func (r *BatchCreateRequest) SetSourceType(sourceType string) {
    r.SourceType = &sourceType
}

/* param cdnType: 点播域名的类型只能是[vod,download,web]中的一种(Optional) */
func (r *BatchCreateRequest) SetCdnType(cdnType string) {
    r.CdnType = &cdnType
}

/* param backSourceType: 回源方式,只能是[https,http]中的一种,默认http(Optional) */
func (r *BatchCreateRequest) SetBackSourceType(backSourceType string) {
    r.BackSourceType = &backSourceType
}

/* param dailyBandWidth: 日带宽(Mbps)(Optional) */
func (r *BatchCreateRequest) SetDailyBandWidth(dailyBandWidth int64) {
    r.DailyBandWidth = &dailyBandWidth
}

/* param quaility: 服务质量,只能是[good,general]中的一种,默认为good(Optional) */
func (r *BatchCreateRequest) SetQuaility(quaility string) {
    r.Quaility = &quaility
}

/* param maxFileSize: (Optional) */
func (r *BatchCreateRequest) SetMaxFileSize(maxFileSize int64) {
    r.MaxFileSize = &maxFileSize
}

/* param minFileSize: (Optional) */
func (r *BatchCreateRequest) SetMinFileSize(minFileSize int64) {
    r.MinFileSize = &minFileSize
}

/* param sumFileSize: (Optional) */
func (r *BatchCreateRequest) SetSumFileSize(sumFileSize int64) {
    r.SumFileSize = &sumFileSize
}

/* param avgFileSize: (Optional) */
func (r *BatchCreateRequest) SetAvgFileSize(avgFileSize int64) {
    r.AvgFileSize = &avgFileSize
}

/* param defaultSourceHost: (Optional) */
func (r *BatchCreateRequest) SetDefaultSourceHost(defaultSourceHost string) {
    r.DefaultSourceHost = &defaultSourceHost
}

/* param httpType: (Optional) */
func (r *BatchCreateRequest) SetHttpType(httpType string) {
    r.HttpType = &httpType
}

/* param ipSource: (Optional) */
func (r *BatchCreateRequest) SetIpSource(ipSource []cdn.IpSourceInfo) {
    r.IpSource = ipSource
}

/* param domainSource: (Optional) */
func (r *BatchCreateRequest) SetDomainSource(domainSource []cdn.DomainSourceInfo) {
    r.DomainSource = domainSource
}

/* param ossSource: (Optional) */
func (r *BatchCreateRequest) SetOssSource(ossSource string) {
    r.OssSource = &ossSource
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r BatchCreateRequest) GetRegionId() string {
    return ""
}

type BatchCreateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result BatchCreateResult `json:"result"`
}

type BatchCreateResult struct {
}