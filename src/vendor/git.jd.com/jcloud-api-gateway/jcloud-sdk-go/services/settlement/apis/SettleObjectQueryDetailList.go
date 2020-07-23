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

type SettleObjectQueryDetailListRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* 结算单号  */
    SettlementNo string `json:"settlementNo"`
}

/*
 * param regionId:  (Required)
 * param settlementNo: 结算单号 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSettleObjectQueryDetailListRequest(
    regionId string,
    settlementNo string,
) *SettleObjectQueryDetailListRequest {

	return &SettleObjectQueryDetailListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/settlementDetail:settleObjectQueryDetailList",
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
 * param settlementNo: 结算单号 (Required)
 */
func NewSettleObjectQueryDetailListRequestWithAllParams(
    regionId string,
    settlementNo string,
) *SettleObjectQueryDetailListRequest {

    return &SettleObjectQueryDetailListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/settlementDetail:settleObjectQueryDetailList",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        SettlementNo: settlementNo,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSettleObjectQueryDetailListRequestWithoutParam() *SettleObjectQueryDetailListRequest {

    return &SettleObjectQueryDetailListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/settlementDetail:settleObjectQueryDetailList",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *SettleObjectQueryDetailListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param settlementNo: 结算单号(Required) */
func (r *SettleObjectQueryDetailListRequest) SetSettlementNo(settlementNo string) {
    r.SettlementNo = settlementNo
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SettleObjectQueryDetailListRequest) GetRegionId() string {
    return r.RegionId
}

type SettleObjectQueryDetailListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SettleObjectQueryDetailListResult `json:"result"`
}

type SettleObjectQueryDetailListResult struct {
    Pagination settlement.Pagination `json:"pagination"`
    Result []settlement.SettlementDetailVo `json:"result"`
}