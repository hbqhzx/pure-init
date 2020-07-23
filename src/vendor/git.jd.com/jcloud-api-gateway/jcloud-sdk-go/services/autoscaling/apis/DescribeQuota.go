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

type DescribeQuotaRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 类型 group/config/schedule/dynamic  */
    Type string `json:"type"`

    /* 当type=schedule/dynamic (Optional) */
    GroupUuid *string `json:"groupUuid"`
}

/*
 * param regionId: 地域 Id (Required)
 * param type_: 类型 group/config/schedule/dynamic (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeQuotaRequest(
    regionId string,
    type_ string,
) *DescribeQuotaRequest {

	return &DescribeQuotaRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/quota",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Type: type_,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param type_: 类型 group/config/schedule/dynamic (Required)
 * param groupUuid: 当type=schedule/dynamic (Optional)
 */
func NewDescribeQuotaRequestWithAllParams(
    regionId string,
    type_ string,
    groupUuid *string,
) *DescribeQuotaRequest {

    return &DescribeQuotaRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/quota",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Type: type_,
        GroupUuid: groupUuid,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeQuotaRequestWithoutParam() *DescribeQuotaRequest {

    return &DescribeQuotaRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/quota",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DescribeQuotaRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param type_: 类型 group/config/schedule/dynamic(Required) */
func (r *DescribeQuotaRequest) SetType(type_ string) {
    r.Type = type_
}

/* param groupUuid: 当type=schedule/dynamic(Optional) */
func (r *DescribeQuotaRequest) SetGroupUuid(groupUuid string) {
    r.GroupUuid = &groupUuid
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeQuotaRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeQuotaResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeQuotaResult `json:"result"`
}

type DescribeQuotaResult struct {
    Rest int64 `json:"rest"`
    Total int64 `json:"total"`
}