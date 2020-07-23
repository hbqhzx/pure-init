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

type ExportSubDistributorAchievmentRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* 按月查询开始时间  */
    BeginTime string `json:"beginTime"`

    /* 按月查询结束时间  */
    EndTime string `json:"endTime"`

    /* 服务商名称 (Optional) */
    DistributorName *string `json:"distributorName"`
}

/*
 * param regionId:  (Required)
 * param beginTime: 按月查询开始时间 (Required)
 * param endTime: 按月查询结束时间 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewExportSubDistributorAchievmentRequest(
    regionId string,
    beginTime string,
    endTime string,
) *ExportSubDistributorAchievmentRequest {

	return &ExportSubDistributorAchievmentRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/achievementManage/exportSubDistributorAchievment",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        BeginTime: beginTime,
        EndTime: endTime,
	}
}

/*
 * param regionId:  (Required)
 * param beginTime: 按月查询开始时间 (Required)
 * param endTime: 按月查询结束时间 (Required)
 * param distributorName: 服务商名称 (Optional)
 */
func NewExportSubDistributorAchievmentRequestWithAllParams(
    regionId string,
    beginTime string,
    endTime string,
    distributorName *string,
) *ExportSubDistributorAchievmentRequest {

    return &ExportSubDistributorAchievmentRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/achievementManage/exportSubDistributorAchievment",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        BeginTime: beginTime,
        EndTime: endTime,
        DistributorName: distributorName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewExportSubDistributorAchievmentRequestWithoutParam() *ExportSubDistributorAchievmentRequest {

    return &ExportSubDistributorAchievmentRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/achievementManage/exportSubDistributorAchievment",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *ExportSubDistributorAchievmentRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param beginTime: 按月查询开始时间(Required) */
func (r *ExportSubDistributorAchievmentRequest) SetBeginTime(beginTime string) {
    r.BeginTime = beginTime
}

/* param endTime: 按月查询结束时间(Required) */
func (r *ExportSubDistributorAchievmentRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

/* param distributorName: 服务商名称(Optional) */
func (r *ExportSubDistributorAchievmentRequest) SetDistributorName(distributorName string) {
    r.DistributorName = &distributorName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ExportSubDistributorAchievmentRequest) GetRegionId() string {
    return r.RegionId
}

type ExportSubDistributorAchievmentResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ExportSubDistributorAchievmentResult `json:"result"`
}

type ExportSubDistributorAchievmentResult struct {
    Uri string `json:"uri"`
}