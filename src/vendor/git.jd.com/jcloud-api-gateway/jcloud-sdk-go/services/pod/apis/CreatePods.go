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
    pod "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/pod/models"
)

type CreatePodsRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* pod 创建参数 (Optional) */
    PodSpec *pod.PodSpec `json:"podSpec"`

    /* 购买实例数量；取值范围：[1,100] (Optional) */
    MaxCount *int `json:"maxCount"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreatePodsRequest(
    regionId string,
) *CreatePodsRequest {

	return &CreatePodsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/pods",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param podSpec: pod 创建参数 (Optional)
 * param maxCount: 购买实例数量；取值范围：[1,100] (Optional)
 */
func NewCreatePodsRequestWithAllParams(
    regionId string,
    podSpec *pod.PodSpec,
    maxCount *int,
) *CreatePodsRequest {

    return &CreatePodsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pods",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PodSpec: podSpec,
        MaxCount: maxCount,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreatePodsRequestWithoutParam() *CreatePodsRequest {

    return &CreatePodsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pods",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreatePodsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param podSpec: pod 创建参数(Optional) */
func (r *CreatePodsRequest) SetPodSpec(podSpec *pod.PodSpec) {
    r.PodSpec = podSpec
}

/* param maxCount: 购买实例数量；取值范围：[1,100](Optional) */
func (r *CreatePodsRequest) SetMaxCount(maxCount int) {
    r.MaxCount = &maxCount
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreatePodsRequest) GetRegionId() string {
    return r.RegionId
}

type CreatePodsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreatePodsResult `json:"result"`
}

type CreatePodsResult struct {
    PodIds []string `json:"podIds"`
}