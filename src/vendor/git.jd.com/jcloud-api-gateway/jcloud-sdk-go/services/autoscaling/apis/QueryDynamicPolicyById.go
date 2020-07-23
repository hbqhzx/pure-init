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
    autoscaling "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/autoscaling/models"
)

type QueryDynamicPolicyByIdRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 组uuid  */
    AsGroupUuid string `json:"asGroupUuid"`

    /* 动态策略uuid  */
    DynamicPolicyUuid string `json:"dynamicPolicyUuid"`
}

/*
 * param regionId: 地域 Id (Required)
 * param asGroupUuid: 组uuid (Required)
 * param dynamicPolicyUuid: 动态策略uuid (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryDynamicPolicyByIdRequest(
    regionId string,
    asGroupUuid string,
    dynamicPolicyUuid string,
) *QueryDynamicPolicyByIdRequest {

	return &QueryDynamicPolicyByIdRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/autoscalingGroups/{asGroupUuid}/dynamicPolicies/{dynamicPolicyUuid}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AsGroupUuid: asGroupUuid,
        DynamicPolicyUuid: dynamicPolicyUuid,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param asGroupUuid: 组uuid (Required)
 * param dynamicPolicyUuid: 动态策略uuid (Required)
 */
func NewQueryDynamicPolicyByIdRequestWithAllParams(
    regionId string,
    asGroupUuid string,
    dynamicPolicyUuid string,
) *QueryDynamicPolicyByIdRequest {

    return &QueryDynamicPolicyByIdRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/autoscalingGroups/{asGroupUuid}/dynamicPolicies/{dynamicPolicyUuid}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AsGroupUuid: asGroupUuid,
        DynamicPolicyUuid: dynamicPolicyUuid,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryDynamicPolicyByIdRequestWithoutParam() *QueryDynamicPolicyByIdRequest {

    return &QueryDynamicPolicyByIdRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/autoscalingGroups/{asGroupUuid}/dynamicPolicies/{dynamicPolicyUuid}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *QueryDynamicPolicyByIdRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param asGroupUuid: 组uuid(Required) */
func (r *QueryDynamicPolicyByIdRequest) SetAsGroupUuid(asGroupUuid string) {
    r.AsGroupUuid = asGroupUuid
}

/* param dynamicPolicyUuid: 动态策略uuid(Required) */
func (r *QueryDynamicPolicyByIdRequest) SetDynamicPolicyUuid(dynamicPolicyUuid string) {
    r.DynamicPolicyUuid = dynamicPolicyUuid
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryDynamicPolicyByIdRequest) GetRegionId() string {
    return r.RegionId
}

type QueryDynamicPolicyByIdResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryDynamicPolicyByIdResult `json:"result"`
}

type QueryDynamicPolicyByIdResult struct {
    AsDynamicPolicy autoscaling.AsDynamicPolicyVo `json:"asDynamicPolicy"`
}