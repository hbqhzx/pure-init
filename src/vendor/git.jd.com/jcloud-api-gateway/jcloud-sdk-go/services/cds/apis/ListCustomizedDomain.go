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

type ListCustomizedDomainRequest struct {

    core.JDCloudRequest

    /* 区域ID  */
    RegionId string `json:"regionId"`

    /* Bucket名称  */
    BucketName string `json:"bucketName"`

    /* 自定义域名状态(0无效；1有效；2被找回；3删除) (Optional) */
    Status *int `json:"status"`
}

/*
 * param regionId: 区域ID (Required)
 * param bucketName: Bucket名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewListCustomizedDomainRequest(
    regionId string,
    bucketName string,
) *ListCustomizedDomainRequest {

	return &ListCustomizedDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/buckets/{bucketName}/customizedDomain",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        BucketName: bucketName,
	}
}

/*
 * param regionId: 区域ID (Required)
 * param bucketName: Bucket名称 (Required)
 * param status: 自定义域名状态(0无效；1有效；2被找回；3删除) (Optional)
 */
func NewListCustomizedDomainRequestWithAllParams(
    regionId string,
    bucketName string,
    status *int,
) *ListCustomizedDomainRequest {

    return &ListCustomizedDomainRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/buckets/{bucketName}/customizedDomain",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        BucketName: bucketName,
        Status: status,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewListCustomizedDomainRequestWithoutParam() *ListCustomizedDomainRequest {

    return &ListCustomizedDomainRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/buckets/{bucketName}/customizedDomain",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域ID(Required) */
func (r *ListCustomizedDomainRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param bucketName: Bucket名称(Required) */
func (r *ListCustomizedDomainRequest) SetBucketName(bucketName string) {
    r.BucketName = bucketName
}

/* param status: 自定义域名状态(0无效；1有效；2被找回；3删除)(Optional) */
func (r *ListCustomizedDomainRequest) SetStatus(status int) {
    r.Status = &status
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ListCustomizedDomainRequest) GetRegionId() string {
    return r.RegionId
}

type ListCustomizedDomainResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ListCustomizedDomainResult `json:"result"`
}

type ListCustomizedDomainResult struct {
}