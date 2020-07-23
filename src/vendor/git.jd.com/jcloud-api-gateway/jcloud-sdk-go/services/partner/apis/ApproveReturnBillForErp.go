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

type ApproveReturnBillForErpRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* 返还单ID  */
    RetrunBillId string `json:"retrunBillId"`

    /* 审批状态（1：通过 2：驳回）  */
    Status int `json:"status"`

    /* 原因 (Optional) */
    Reason *string `json:"reason"`
}

/*
 * param regionId:  (Required)
 * param retrunBillId: 返还单ID (Required)
 * param status: 审批状态（1：通过 2：驳回） (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewApproveReturnBillForErpRequest(
    regionId string,
    retrunBillId string,
    status int,
) *ApproveReturnBillForErpRequest {

	return &ApproveReturnBillForErpRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/returnBill/approveReturnBillForErp",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        RetrunBillId: retrunBillId,
        Status: status,
	}
}

/*
 * param regionId:  (Required)
 * param retrunBillId: 返还单ID (Required)
 * param status: 审批状态（1：通过 2：驳回） (Required)
 * param reason: 原因 (Optional)
 */
func NewApproveReturnBillForErpRequestWithAllParams(
    regionId string,
    retrunBillId string,
    status int,
    reason *string,
) *ApproveReturnBillForErpRequest {

    return &ApproveReturnBillForErpRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/returnBill/approveReturnBillForErp",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        RetrunBillId: retrunBillId,
        Status: status,
        Reason: reason,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewApproveReturnBillForErpRequestWithoutParam() *ApproveReturnBillForErpRequest {

    return &ApproveReturnBillForErpRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/returnBill/approveReturnBillForErp",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *ApproveReturnBillForErpRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param retrunBillId: 返还单ID(Required) */
func (r *ApproveReturnBillForErpRequest) SetRetrunBillId(retrunBillId string) {
    r.RetrunBillId = retrunBillId
}

/* param status: 审批状态（1：通过 2：驳回）(Required) */
func (r *ApproveReturnBillForErpRequest) SetStatus(status int) {
    r.Status = status
}

/* param reason: 原因(Optional) */
func (r *ApproveReturnBillForErpRequest) SetReason(reason string) {
    r.Reason = &reason
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ApproveReturnBillForErpRequest) GetRegionId() string {
    return r.RegionId
}

type ApproveReturnBillForErpResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ApproveReturnBillForErpResult `json:"result"`
}

type ApproveReturnBillForErpResult struct {
    Result string `json:"result"`
    Msg string `json:"msg"`
}