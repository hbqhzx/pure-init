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

type ModifyCacheInstanceAttributeRequest struct {

    core.JDCloudRequest

    /* 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2  */
    RegionId string `json:"regionId"`

    /* 缓存Redis实例ID，是访问实例的唯一标识  */
    CacheInstanceId string `json:"cacheInstanceId"`

    /* 实例的名称，名称只支持数字、字母、英文下划线、中文，且不少于2字符不超过32字符 (Optional) */
    CacheInstanceName *string `json:"cacheInstanceName"`

    /* 实例的描述，不能超过256个字符 (Optional) */
    CacheInstanceDescription *string `json:"cacheInstanceDescription"`
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyCacheInstanceAttributeRequest(
    regionId string,
    cacheInstanceId string,
) *ModifyCacheInstanceAttributeRequest {

	return &ModifyCacheInstanceAttributeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
	}
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识 (Required)
 * param cacheInstanceName: 实例的名称，名称只支持数字、字母、英文下划线、中文，且不少于2字符不超过32字符 (Optional)
 * param cacheInstanceDescription: 实例的描述，不能超过256个字符 (Optional)
 */
func NewModifyCacheInstanceAttributeRequestWithAllParams(
    regionId string,
    cacheInstanceId string,
    cacheInstanceName *string,
    cacheInstanceDescription *string,
) *ModifyCacheInstanceAttributeRequest {

    return &ModifyCacheInstanceAttributeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
        CacheInstanceName: cacheInstanceName,
        CacheInstanceDescription: cacheInstanceDescription,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyCacheInstanceAttributeRequestWithoutParam() *ModifyCacheInstanceAttributeRequest {

    return &ModifyCacheInstanceAttributeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2(Required) */
func (r *ModifyCacheInstanceAttributeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识(Required) */
func (r *ModifyCacheInstanceAttributeRequest) SetCacheInstanceId(cacheInstanceId string) {
    r.CacheInstanceId = cacheInstanceId
}

/* param cacheInstanceName: 实例的名称，名称只支持数字、字母、英文下划线、中文，且不少于2字符不超过32字符(Optional) */
func (r *ModifyCacheInstanceAttributeRequest) SetCacheInstanceName(cacheInstanceName string) {
    r.CacheInstanceName = &cacheInstanceName
}

/* param cacheInstanceDescription: 实例的描述，不能超过256个字符(Optional) */
func (r *ModifyCacheInstanceAttributeRequest) SetCacheInstanceDescription(cacheInstanceDescription string) {
    r.CacheInstanceDescription = &cacheInstanceDescription
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyCacheInstanceAttributeRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyCacheInstanceAttributeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyCacheInstanceAttributeResult `json:"result"`
}

type ModifyCacheInstanceAttributeResult struct {
}