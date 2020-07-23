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
    live "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/live/models"
)

type DescribeLiveTimeshiftConfigsRequest struct {

    core.JDCloudRequest

    /* 页码；默认为1；取值范围[1, 100000] (Optional) */
    PageNum *int `json:"pageNum"`

    /* 分页大小；默认为10；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 直播的推流域名  */
    PlayDomain string `json:"playDomain"`
}

/*
 * param playDomain: 直播的推流域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLiveTimeshiftConfigsRequest(
    playDomain string,
) *DescribeLiveTimeshiftConfigsRequest {

	return &DescribeLiveTimeshiftConfigsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/liveTimeshift:configs",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        PlayDomain: playDomain,
	}
}

/*
 * param pageNum: 页码；默认为1；取值范围[1, 100000] (Optional)
 * param pageSize: 分页大小；默认为10；取值范围[10, 100] (Optional)
 * param playDomain: 直播的推流域名 (Required)
 */
func NewDescribeLiveTimeshiftConfigsRequestWithAllParams(
    pageNum *int,
    pageSize *int,
    playDomain string,
) *DescribeLiveTimeshiftConfigsRequest {

    return &DescribeLiveTimeshiftConfigsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveTimeshift:configs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNum: pageNum,
        PageSize: pageSize,
        PlayDomain: playDomain,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLiveTimeshiftConfigsRequestWithoutParam() *DescribeLiveTimeshiftConfigsRequest {

    return &DescribeLiveTimeshiftConfigsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveTimeshift:configs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNum: 页码；默认为1；取值范围[1, 100000](Optional) */
func (r *DescribeLiveTimeshiftConfigsRequest) SetPageNum(pageNum int) {
    r.PageNum = &pageNum
}

/* param pageSize: 分页大小；默认为10；取值范围[10, 100](Optional) */
func (r *DescribeLiveTimeshiftConfigsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param playDomain: 直播的推流域名(Required) */
func (r *DescribeLiveTimeshiftConfigsRequest) SetPlayDomain(playDomain string) {
    r.PlayDomain = playDomain
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLiveTimeshiftConfigsRequest) GetRegionId() string {
    return ""
}

type DescribeLiveTimeshiftConfigsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLiveTimeshiftConfigsResult `json:"result"`
}

type DescribeLiveTimeshiftConfigsResult struct {
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalCount int `json:"totalCount"`
    TimeshiftConfigs []live.TimeshiftConfig `json:"timeshiftConfigs"`
}