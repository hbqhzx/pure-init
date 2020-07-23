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

type QueryInstancesByAutoscalingGroupIdRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 组uuid  */
    AsGroupUuid string `json:"asGroupUuid"`

    /* PageNumber (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* PageSize (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 地域 Id (Required)
 * param asGroupUuid: 组uuid (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryInstancesByAutoscalingGroupIdRequest(
    regionId string,
    asGroupUuid string,
) *QueryInstancesByAutoscalingGroupIdRequest {

	return &QueryInstancesByAutoscalingGroupIdRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/autoscalingGroups/{asGroupUuid}/instances",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AsGroupUuid: asGroupUuid,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param asGroupUuid: 组uuid (Required)
 * param pageNumber: PageNumber (Optional)
 * param pageSize: PageSize (Optional)
 */
func NewQueryInstancesByAutoscalingGroupIdRequestWithAllParams(
    regionId string,
    asGroupUuid string,
    pageNumber *int,
    pageSize *int,
) *QueryInstancesByAutoscalingGroupIdRequest {

    return &QueryInstancesByAutoscalingGroupIdRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/autoscalingGroups/{asGroupUuid}/instances",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AsGroupUuid: asGroupUuid,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryInstancesByAutoscalingGroupIdRequestWithoutParam() *QueryInstancesByAutoscalingGroupIdRequest {

    return &QueryInstancesByAutoscalingGroupIdRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/autoscalingGroups/{asGroupUuid}/instances",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *QueryInstancesByAutoscalingGroupIdRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param asGroupUuid: 组uuid(Required) */
func (r *QueryInstancesByAutoscalingGroupIdRequest) SetAsGroupUuid(asGroupUuid string) {
    r.AsGroupUuid = asGroupUuid
}

/* param pageNumber: PageNumber(Optional) */
func (r *QueryInstancesByAutoscalingGroupIdRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: PageSize(Optional) */
func (r *QueryInstancesByAutoscalingGroupIdRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryInstancesByAutoscalingGroupIdRequest) GetRegionId() string {
    return r.RegionId
}

type QueryInstancesByAutoscalingGroupIdResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryInstancesByAutoscalingGroupIdResult `json:"result"`
}

type QueryInstancesByAutoscalingGroupIdResult struct {
    AsInstances []autoscaling.AsInstanceResponseVO `json:"asInstances"`
    PageNumber int64 `json:"pageNumber"`
    PageSize int64 `json:"pageSize"`
    TotalCount int64 `json:"totalCount"`
    TotalPages int64 `json:"totalPages"`
}