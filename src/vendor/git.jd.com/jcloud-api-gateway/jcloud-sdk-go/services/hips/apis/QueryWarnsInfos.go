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
    hips "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/hips/models"
)

type QueryWarnsInfosRequest struct {

    core.JDCloudRequest

    /*  (Optional) */
    RegionId *string `json:"regionId"`

    /*  (Optional) */
    Interval *int `json:"interval"`

    /* 传小于等于0为全部获取 (Optional) */
    WarningType *int `json:"warningType"`

    /*  (Optional) */
    ServerName *string `json:"serverName"`

    /*  (Optional) */
    PageNumber *int `json:"pageNumber"`

    /*  (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryWarnsInfosRequest(
) *QueryWarnsInfosRequest {

	return &QueryWarnsInfosRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/warns:queryInfos",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param regionId:  (Optional)
 * param interval:  (Optional)
 * param warningType: 传小于等于0为全部获取 (Optional)
 * param serverName:  (Optional)
 * param pageNumber:  (Optional)
 * param pageSize:  (Optional)
 */
func NewQueryWarnsInfosRequestWithAllParams(
    regionId *string,
    interval *int,
    warningType *int,
    serverName *string,
    pageNumber *int,
    pageSize *int,
) *QueryWarnsInfosRequest {

    return &QueryWarnsInfosRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/warns:queryInfos",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Interval: interval,
        WarningType: warningType,
        ServerName: serverName,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryWarnsInfosRequestWithoutParam() *QueryWarnsInfosRequest {

    return &QueryWarnsInfosRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/warns:queryInfos",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Optional) */
func (r *QueryWarnsInfosRequest) SetRegionId(regionId string) {
    r.RegionId = &regionId
}

/* param interval: (Optional) */
func (r *QueryWarnsInfosRequest) SetInterval(interval int) {
    r.Interval = &interval
}

/* param warningType: 传小于等于0为全部获取(Optional) */
func (r *QueryWarnsInfosRequest) SetWarningType(warningType int) {
    r.WarningType = &warningType
}

/* param serverName: (Optional) */
func (r *QueryWarnsInfosRequest) SetServerName(serverName string) {
    r.ServerName = &serverName
}

/* param pageNumber: (Optional) */
func (r *QueryWarnsInfosRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: (Optional) */
func (r *QueryWarnsInfosRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryWarnsInfosRequest) GetRegionId() string {
    return ""
}

type QueryWarnsInfosResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryWarnsInfosResult `json:"result"`
}

type QueryWarnsInfosResult struct {
    WarnInfos []hips.TotalWarnInfo `json:"warnInfos"`
    TotalCount int `json:"totalCount"`
}