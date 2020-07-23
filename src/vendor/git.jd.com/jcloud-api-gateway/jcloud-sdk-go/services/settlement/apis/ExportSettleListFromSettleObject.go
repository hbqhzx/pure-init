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

type ExportSettleListFromSettleObjectRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* 结算对象pin (Optional) */
    Pin *string `json:"pin"`

    /* 结算单号 (Optional) */
    SettlementNo *string `json:"settlementNo"`

    /* 业务系统 (Optional) */
    SystemCode *string `json:"systemCode"`

    /* 结算对象类型 (Optional) */
    SettleObjectType *string `json:"settleObjectType"`

    /* 结算单状态 (Optional) */
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
func NewExportSettleListFromSettleObjectRequest(
    regionId string,
) *ExportSettleListFromSettleObjectRequest {

	return &ExportSettleListFromSettleObjectRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/settlement:settleObjectExport",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId:  (Required)
 * param pin: 结算对象pin (Optional)
 * param settlementNo: 结算单号 (Optional)
 * param systemCode: 业务系统 (Optional)
 * param settleObjectType: 结算对象类型 (Optional)
 * param status: 结算单状态 (Optional)
 * param startTime: 账期开始时间 (Optional)
 * param endTime: 账期结束时间 (Optional)
 */
func NewExportSettleListFromSettleObjectRequestWithAllParams(
    regionId string,
    pin *string,
    settlementNo *string,
    systemCode *string,
    settleObjectType *string,
    status *int,
    startTime *string,
    endTime *string,
) *ExportSettleListFromSettleObjectRequest {

    return &ExportSettleListFromSettleObjectRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/settlement:settleObjectExport",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Pin: pin,
        SettlementNo: settlementNo,
        SystemCode: systemCode,
        SettleObjectType: settleObjectType,
        Status: status,
        StartTime: startTime,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewExportSettleListFromSettleObjectRequestWithoutParam() *ExportSettleListFromSettleObjectRequest {

    return &ExportSettleListFromSettleObjectRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/settlement:settleObjectExport",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *ExportSettleListFromSettleObjectRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pin: 结算对象pin(Optional) */
func (r *ExportSettleListFromSettleObjectRequest) SetPin(pin string) {
    r.Pin = &pin
}

/* param settlementNo: 结算单号(Optional) */
func (r *ExportSettleListFromSettleObjectRequest) SetSettlementNo(settlementNo string) {
    r.SettlementNo = &settlementNo
}

/* param systemCode: 业务系统(Optional) */
func (r *ExportSettleListFromSettleObjectRequest) SetSystemCode(systemCode string) {
    r.SystemCode = &systemCode
}

/* param settleObjectType: 结算对象类型(Optional) */
func (r *ExportSettleListFromSettleObjectRequest) SetSettleObjectType(settleObjectType string) {
    r.SettleObjectType = &settleObjectType
}

/* param status: 结算单状态(Optional) */
func (r *ExportSettleListFromSettleObjectRequest) SetStatus(status int) {
    r.Status = &status
}

/* param startTime: 账期开始时间(Optional) */
func (r *ExportSettleListFromSettleObjectRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 账期结束时间(Optional) */
func (r *ExportSettleListFromSettleObjectRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ExportSettleListFromSettleObjectRequest) GetRegionId() string {
    return r.RegionId
}

type ExportSettleListFromSettleObjectResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ExportSettleListFromSettleObjectResult `json:"result"`
}

type ExportSettleListFromSettleObjectResult struct {
    Uri string `json:"uri"`
}