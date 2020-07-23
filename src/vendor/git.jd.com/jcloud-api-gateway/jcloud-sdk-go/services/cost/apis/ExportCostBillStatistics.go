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

type ExportCostBillStatisticsRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* 分组类型 1:产品 2：消费类型 3：地域 0:全部  */
    GroupType int `json:"groupType"`

    /* 起始时间  */
    StartTime string `json:"startTime"`

    /* 结束时间  */
    EndTime string `json:"endTime"`
}

/*
 * param regionId:  (Required)
 * param groupType: 分组类型 1:产品 2：消费类型 3：地域 0:全部 (Required)
 * param startTime: 起始时间 (Required)
 * param endTime: 结束时间 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewExportCostBillStatisticsRequest(
    regionId string,
    groupType int,
    startTime string,
    endTime string,
) *ExportCostBillStatisticsRequest {

	return &ExportCostBillStatisticsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/exportCostBillStatistics",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        GroupType: groupType,
        StartTime: startTime,
        EndTime: endTime,
	}
}

/*
 * param regionId:  (Required)
 * param groupType: 分组类型 1:产品 2：消费类型 3：地域 0:全部 (Required)
 * param startTime: 起始时间 (Required)
 * param endTime: 结束时间 (Required)
 */
func NewExportCostBillStatisticsRequestWithAllParams(
    regionId string,
    groupType int,
    startTime string,
    endTime string,
) *ExportCostBillStatisticsRequest {

    return &ExportCostBillStatisticsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/exportCostBillStatistics",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        GroupType: groupType,
        StartTime: startTime,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewExportCostBillStatisticsRequestWithoutParam() *ExportCostBillStatisticsRequest {

    return &ExportCostBillStatisticsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/exportCostBillStatistics",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *ExportCostBillStatisticsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param groupType: 分组类型 1:产品 2：消费类型 3：地域 0:全部(Required) */
func (r *ExportCostBillStatisticsRequest) SetGroupType(groupType int) {
    r.GroupType = groupType
}

/* param startTime: 起始时间(Required) */
func (r *ExportCostBillStatisticsRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 结束时间(Required) */
func (r *ExportCostBillStatisticsRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ExportCostBillStatisticsRequest) GetRegionId() string {
    return r.RegionId
}

type ExportCostBillStatisticsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ExportCostBillStatisticsResult `json:"result"`
}

type ExportCostBillStatisticsResult struct {
    Result string `json:"result"`
}