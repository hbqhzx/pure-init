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

type BusinessQueryDetailListRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* 结算单号  */
    SettlementNo string `json:"settlementNo"`

    /* 页数（默认1） (Optional) */
    PageIndex *int `json:"pageIndex"`

    /* 页条数（默认10） (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId:  (Required)
 * param settlementNo: 结算单号 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewBusinessQueryDetailListRequest(
    regionId string,
    settlementNo string,
) *BusinessQueryDetailListRequest {

	return &BusinessQueryDetailListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/settlementDetail:businessQueryDetailList",
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
 * param pageIndex: 页数（默认1） (Optional)
 * param pageSize: 页条数（默认10） (Optional)
 */
func NewBusinessQueryDetailListRequestWithAllParams(
    regionId string,
    settlementNo string,
    pageIndex *int,
    pageSize *int,
) *BusinessQueryDetailListRequest {

    return &BusinessQueryDetailListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/settlementDetail:businessQueryDetailList",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        SettlementNo: settlementNo,
        PageIndex: pageIndex,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewBusinessQueryDetailListRequestWithoutParam() *BusinessQueryDetailListRequest {

    return &BusinessQueryDetailListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/settlementDetail:businessQueryDetailList",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *BusinessQueryDetailListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param settlementNo: 结算单号(Required) */
func (r *BusinessQueryDetailListRequest) SetSettlementNo(settlementNo string) {
    r.SettlementNo = settlementNo
}

/* param pageIndex: 页数（默认1）(Optional) */
func (r *BusinessQueryDetailListRequest) SetPageIndex(pageIndex int) {
    r.PageIndex = &pageIndex
}

/* param pageSize: 页条数（默认10）(Optional) */
func (r *BusinessQueryDetailListRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r BusinessQueryDetailListRequest) GetRegionId() string {
    return r.RegionId
}

type BusinessQueryDetailListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result BusinessQueryDetailListResult `json:"result"`
}

type BusinessQueryDetailListResult struct {
    Pagination settlement.Pagination `json:"pagination"`
    Result []settlement.SettlementDetailVo `json:"result"`
}