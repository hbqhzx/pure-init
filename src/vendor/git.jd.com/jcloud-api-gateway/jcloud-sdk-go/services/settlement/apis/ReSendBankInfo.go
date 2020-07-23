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

type ReSendBankInfoRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /*   */
    SettlementNo string `json:"settlementNo"`
}

/*
 * param regionId:  (Required)
 * param settlementNo:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewReSendBankInfoRequest(
    regionId string,
    settlementNo string,
) *ReSendBankInfoRequest {

	return &ReSendBankInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/settlement/{settlementNo}:reSendBankInfo",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        SettlementNo: settlementNo,
	}
}

/*
 * param regionId:  (Required)
 * param settlementNo:  (Required)
 */
func NewReSendBankInfoRequestWithAllParams(
    regionId string,
    settlementNo string,
) *ReSendBankInfoRequest {

    return &ReSendBankInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/settlement/{settlementNo}:reSendBankInfo",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        SettlementNo: settlementNo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewReSendBankInfoRequestWithoutParam() *ReSendBankInfoRequest {

    return &ReSendBankInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/settlement/{settlementNo}:reSendBankInfo",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *ReSendBankInfoRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param settlementNo: (Required) */
func (r *ReSendBankInfoRequest) SetSettlementNo(settlementNo string) {
    r.SettlementNo = settlementNo
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ReSendBankInfoRequest) GetRegionId() string {
    return r.RegionId
}

type ReSendBankInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ReSendBankInfoResult `json:"result"`
}

type ReSendBankInfoResult struct {
    RetCode int `json:"retCode"`
    RetMsg string `json:"retMsg"`
}