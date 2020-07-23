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

type DescribeDDoSGraphRequest struct {

    core.JDCloudRequest

    /* 区域 Id  */
    RegionId string `json:"regionId"`

    /* 开始时间, 最多查最近 60 天, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ  */
    StartTime string `json:"startTime"`

    /* 查询的结束时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ  */
    EndTime string `json:"endTime"`

    /* 高防实例 Id 列表 (Optional) */
    InstanceId []int64 `json:"instanceId"`
}

/*
 * param regionId: 区域 Id (Required)
 * param startTime: 开始时间, 最多查最近 60 天, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Required)
 * param endTime: 查询的结束时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDDoSGraphRequest(
    regionId string,
    startTime string,
    endTime string,
) *DescribeDDoSGraphRequest {

	return &DescribeDDoSGraphRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/charts:DDoSGraph",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        StartTime: startTime,
        EndTime: endTime,
	}
}

/*
 * param regionId: 区域 Id (Required)
 * param startTime: 开始时间, 最多查最近 60 天, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Required)
 * param endTime: 查询的结束时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ (Required)
 * param instanceId: 高防实例 Id 列表 (Optional)
 */
func NewDescribeDDoSGraphRequestWithAllParams(
    regionId string,
    startTime string,
    endTime string,
    instanceId []int64,
) *DescribeDDoSGraphRequest {

    return &DescribeDDoSGraphRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/charts:DDoSGraph",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        StartTime: startTime,
        EndTime: endTime,
        InstanceId: instanceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDDoSGraphRequestWithoutParam() *DescribeDDoSGraphRequest {

    return &DescribeDDoSGraphRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/charts:DDoSGraph",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 Id(Required) */
func (r *DescribeDDoSGraphRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param startTime: 开始时间, 最多查最近 60 天, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ(Required) */
func (r *DescribeDDoSGraphRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 查询的结束时间, UTC 时间, 格式: yyyy-MM-dd'T'HH:mm:ssZ(Required) */
func (r *DescribeDDoSGraphRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

/* param instanceId: 高防实例 Id 列表(Optional) */
func (r *DescribeDDoSGraphRequest) SetInstanceId(instanceId []int64) {
    r.InstanceId = instanceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDDoSGraphRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDDoSGraphResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDDoSGraphResult `json:"result"`
}

type DescribeDDoSGraphResult struct {
    PreProtect []float64 `json:"preProtect"`
    PostProtect []float64 `json:"postProtect"`
    Time []string `json:"time"`
    Unit string `json:"unit"`
}