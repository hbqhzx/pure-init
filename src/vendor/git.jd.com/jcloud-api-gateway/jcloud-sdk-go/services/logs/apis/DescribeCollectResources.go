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
    logs "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/logs/models"
)

type DescribeCollectResourcesRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 采集配置 UID  */
    CollectInfoUID string `json:"collectInfoUID"`

    /* 当前所在页，默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 页面大小，默认为20；取值范围[1, 100] (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 地域 Id (Required)
 * param collectInfoUID: 采集配置 UID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeCollectResourcesRequest(
    regionId string,
    collectInfoUID string,
) *DescribeCollectResourcesRequest {

	return &DescribeCollectResourcesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/collectinfos/{collectInfoUID}/resources",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CollectInfoUID: collectInfoUID,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param collectInfoUID: 采集配置 UID (Required)
 * param pageNumber: 当前所在页，默认为1 (Optional)
 * param pageSize: 页面大小，默认为20；取值范围[1, 100] (Optional)
 */
func NewDescribeCollectResourcesRequestWithAllParams(
    regionId string,
    collectInfoUID string,
    pageNumber *int,
    pageSize *int,
) *DescribeCollectResourcesRequest {

    return &DescribeCollectResourcesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/collectinfos/{collectInfoUID}/resources",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CollectInfoUID: collectInfoUID,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeCollectResourcesRequestWithoutParam() *DescribeCollectResourcesRequest {

    return &DescribeCollectResourcesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/collectinfos/{collectInfoUID}/resources",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DescribeCollectResourcesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param collectInfoUID: 采集配置 UID(Required) */
func (r *DescribeCollectResourcesRequest) SetCollectInfoUID(collectInfoUID string) {
    r.CollectInfoUID = collectInfoUID
}

/* param pageNumber: 当前所在页，默认为1(Optional) */
func (r *DescribeCollectResourcesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 页面大小，默认为20；取值范围[1, 100](Optional) */
func (r *DescribeCollectResourcesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeCollectResourcesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeCollectResourcesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeCollectResourcesResult `json:"result"`
}

type DescribeCollectResourcesResult struct {
    Data []logs.ResourceEnd `json:"data"`
    NumberPages int64 `json:"numberPages"`
    NumberRecords int64 `json:"numberRecords"`
    PageNumber int64 `json:"pageNumber"`
    PageSize int64 `json:"pageSize"`
}