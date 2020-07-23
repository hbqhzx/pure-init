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

type GetCNameRequest struct {

    core.JDCloudRequest

    /* 区域ID  */
    RegionId string `json:"regionId"`

    /* Bucket名称  */
    BucketName string `json:"bucketName"`

    /* 自定义域名  */
    Cname string `json:"cname"`
}

/*
 * param regionId: 区域ID (Required)
 * param bucketName: Bucket名称 (Required)
 * param cname: 自定义域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetCNameRequest(
    regionId string,
    bucketName string,
    cname string,
) *GetCNameRequest {

	return &GetCNameRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/buckets/{bucketName}/cname/{cname}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        BucketName: bucketName,
        Cname: cname,
	}
}

/*
 * param regionId: 区域ID (Required)
 * param bucketName: Bucket名称 (Required)
 * param cname: 自定义域名 (Required)
 */
func NewGetCNameRequestWithAllParams(
    regionId string,
    bucketName string,
    cname string,
) *GetCNameRequest {

    return &GetCNameRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/buckets/{bucketName}/cname/{cname}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        BucketName: bucketName,
        Cname: cname,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetCNameRequestWithoutParam() *GetCNameRequest {

    return &GetCNameRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/buckets/{bucketName}/cname/{cname}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域ID(Required) */
func (r *GetCNameRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param bucketName: Bucket名称(Required) */
func (r *GetCNameRequest) SetBucketName(bucketName string) {
    r.BucketName = bucketName
}

/* param cname: 自定义域名(Required) */
func (r *GetCNameRequest) SetCname(cname string) {
    r.Cname = cname
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetCNameRequest) GetRegionId() string {
    return r.RegionId
}

type GetCNameResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetCNameResult `json:"result"`
}

type GetCNameResult struct {
    Id int `json:"id"`
    Pin string `json:"pin"`
    OriginDomain string `json:"originDomain"`
    Cname string `json:"cname"`
    Status int `json:"status"`
    BucketName string `json:"bucketName"`
    IsCName int `json:"isCName"`
    CreateTime string `json:"createTime"`
    ProtoType int `json:"protoType"`
}