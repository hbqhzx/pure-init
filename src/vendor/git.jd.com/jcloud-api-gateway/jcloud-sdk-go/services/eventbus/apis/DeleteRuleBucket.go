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

type DeleteRuleBucketRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 订阅事件uuid  */
    Uuid string `json:"uuid"`
}

/*
 * param regionId: 地域 Id (Required)
 * param uuid: 订阅事件uuid (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteRuleBucketRequest(
    regionId string,
    uuid string,
) *DeleteRuleBucketRequest {

	return &DeleteRuleBucketRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/ruleBucket",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Uuid: uuid,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param uuid: 订阅事件uuid (Required)
 */
func NewDeleteRuleBucketRequestWithAllParams(
    regionId string,
    uuid string,
) *DeleteRuleBucketRequest {

    return &DeleteRuleBucketRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ruleBucket",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Uuid: uuid,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteRuleBucketRequestWithoutParam() *DeleteRuleBucketRequest {

    return &DeleteRuleBucketRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ruleBucket",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DeleteRuleBucketRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param uuid: 订阅事件uuid(Required) */
func (r *DeleteRuleBucketRequest) SetUuid(uuid string) {
    r.Uuid = uuid
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteRuleBucketRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteRuleBucketResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteRuleBucketResult `json:"result"`
}

type DeleteRuleBucketResult struct {
    Success bool `json:"success"`
    Uuid string `json:"uuid"`
}