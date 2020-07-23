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
    asset "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/asset/models"
)

type GetWithdrawalsRecordsByPageOpenAPIRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /*  (Optional) */
    Pin *string `json:"pin"`

    /*  (Optional) */
    BeginTime *string `json:"beginTime"`

    /*  (Optional) */
    EndTime *string `json:"endTime"`

    /*  (Optional) */
    PageNum *int `json:"pageNum"`

    /*  (Optional) */
    PageSize *int `json:"pageSize"`

    /*  (Optional) */
    Sort *int `json:"sort"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetWithdrawalsRecordsByPageOpenAPIRequest(
    regionId string,
) *GetWithdrawalsRecordsByPageOpenAPIRequest {

	return &GetWithdrawalsRecordsByPageOpenAPIRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/asset:getWithdrawalsRecordsByPage",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pin:  (Optional)
 * param beginTime:  (Optional)
 * param endTime:  (Optional)
 * param pageNum:  (Optional)
 * param pageSize:  (Optional)
 * param sort:  (Optional)
 */
func NewGetWithdrawalsRecordsByPageOpenAPIRequestWithAllParams(
    regionId string,
    pin *string,
    beginTime *string,
    endTime *string,
    pageNum *int,
    pageSize *int,
    sort *int,
) *GetWithdrawalsRecordsByPageOpenAPIRequest {

    return &GetWithdrawalsRecordsByPageOpenAPIRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/asset:getWithdrawalsRecordsByPage",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Pin: pin,
        BeginTime: beginTime,
        EndTime: endTime,
        PageNum: pageNum,
        PageSize: pageSize,
        Sort: sort,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetWithdrawalsRecordsByPageOpenAPIRequestWithoutParam() *GetWithdrawalsRecordsByPageOpenAPIRequest {

    return &GetWithdrawalsRecordsByPageOpenAPIRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/asset:getWithdrawalsRecordsByPage",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *GetWithdrawalsRecordsByPageOpenAPIRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pin: (Optional) */
func (r *GetWithdrawalsRecordsByPageOpenAPIRequest) SetPin(pin string) {
    r.Pin = &pin
}

/* param beginTime: (Optional) */
func (r *GetWithdrawalsRecordsByPageOpenAPIRequest) SetBeginTime(beginTime string) {
    r.BeginTime = &beginTime
}

/* param endTime: (Optional) */
func (r *GetWithdrawalsRecordsByPageOpenAPIRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param pageNum: (Optional) */
func (r *GetWithdrawalsRecordsByPageOpenAPIRequest) SetPageNum(pageNum int) {
    r.PageNum = &pageNum
}

/* param pageSize: (Optional) */
func (r *GetWithdrawalsRecordsByPageOpenAPIRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param sort: (Optional) */
func (r *GetWithdrawalsRecordsByPageOpenAPIRequest) SetSort(sort int) {
    r.Sort = &sort
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetWithdrawalsRecordsByPageOpenAPIRequest) GetRegionId() string {
    return r.RegionId
}

type GetWithdrawalsRecordsByPageOpenAPIResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetWithdrawalsRecordsByPageOpenAPIResult `json:"result"`
}

type GetWithdrawalsRecordsByPageOpenAPIResult struct {
    TotalPage int `json:"totalPage"`
    TotalCount int `json:"totalCount"`
    ResultList []asset.WithdrawalsVo `json:"resultList"`
}