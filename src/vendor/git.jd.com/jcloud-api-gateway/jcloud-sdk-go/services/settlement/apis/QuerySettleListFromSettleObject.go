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
    settlement "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/settlement/models"
)

type QuerySettleListFromSettleObjectRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* 页码：默认为1 (Optional) */
    PageIndex *int `json:"pageIndex"`

    /* 分页大小 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 结算单状态，-1:结算中；4:待结算对象确认；5:结算对象驳回；7:无需结算；12:已完成收付款 (Optional) */
    Status *int `json:"status"`

    /* 账期开始时间 (Optional) */
    StartTime *string `json:"startTime"`

    /* 账期结束时间 (Optional) */
    EndTime *string `json:"endTime"`
}

/*
 * param regionId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQuerySettleListFromSettleObjectRequest(
    regionId string,
) *QuerySettleListFromSettleObjectRequest {

	return &QuerySettleListFromSettleObjectRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/settlement:settleObjectList",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId:  (Required)
 * param pageIndex: 页码：默认为1 (Optional)
 * param pageSize: 分页大小 (Optional)
 * param status: 结算单状态，-1:结算中；4:待结算对象确认；5:结算对象驳回；7:无需结算；12:已完成收付款 (Optional)
 * param startTime: 账期开始时间 (Optional)
 * param endTime: 账期结束时间 (Optional)
 */
func NewQuerySettleListFromSettleObjectRequestWithAllParams(
    regionId string,
    pageIndex *int,
    pageSize *int,
    status *int,
    startTime *string,
    endTime *string,
) *QuerySettleListFromSettleObjectRequest {

    return &QuerySettleListFromSettleObjectRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/settlement:settleObjectList",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageIndex: pageIndex,
        PageSize: pageSize,
        Status: status,
        StartTime: startTime,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQuerySettleListFromSettleObjectRequestWithoutParam() *QuerySettleListFromSettleObjectRequest {

    return &QuerySettleListFromSettleObjectRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/settlement:settleObjectList",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *QuerySettleListFromSettleObjectRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageIndex: 页码：默认为1(Optional) */
func (r *QuerySettleListFromSettleObjectRequest) SetPageIndex(pageIndex int) {
    r.PageIndex = &pageIndex
}

/* param pageSize: 分页大小(Optional) */
func (r *QuerySettleListFromSettleObjectRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param status: 结算单状态，-1:结算中；4:待结算对象确认；5:结算对象驳回；7:无需结算；12:已完成收付款(Optional) */
func (r *QuerySettleListFromSettleObjectRequest) SetStatus(status int) {
    r.Status = &status
}

/* param startTime: 账期开始时间(Optional) */
func (r *QuerySettleListFromSettleObjectRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 账期结束时间(Optional) */
func (r *QuerySettleListFromSettleObjectRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QuerySettleListFromSettleObjectRequest) GetRegionId() string {
    return r.RegionId
}

type QuerySettleListFromSettleObjectResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QuerySettleListFromSettleObjectResult `json:"result"`
}

type QuerySettleListFromSettleObjectResult struct {
    Pagination settlement.Pagination `json:"pagination"`
    Result []settlement.SettlementVo `json:"result"`
}