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
    monitor "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/monitor/models"
)

type DescribeObjsByNsUIDRequest struct {

    core.JDCloudRequest

    /* region  */
    RegionId string `json:"regionId"`

    /* namespaceUID  */
    NamespaceUID string `json:"namespaceUID"`

    /* 当前所在页，默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 页面大小，默认为20；取值范围[1, 100] (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: region (Required)
 * param namespaceUID: namespaceUID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeObjsByNsUIDRequest(
    regionId string,
    namespaceUID string,
) *DescribeObjsByNsUIDRequest {

	return &DescribeObjsByNsUIDRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cmNameSpaces/{namespaceUID}/objs",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        NamespaceUID: namespaceUID,
	}
}

/*
 * param regionId: region (Required)
 * param namespaceUID: namespaceUID (Required)
 * param pageNumber: 当前所在页，默认为1 (Optional)
 * param pageSize: 页面大小，默认为20；取值范围[1, 100] (Optional)
 */
func NewDescribeObjsByNsUIDRequestWithAllParams(
    regionId string,
    namespaceUID string,
    pageNumber *int,
    pageSize *int,
) *DescribeObjsByNsUIDRequest {

    return &DescribeObjsByNsUIDRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cmNameSpaces/{namespaceUID}/objs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        NamespaceUID: namespaceUID,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeObjsByNsUIDRequestWithoutParam() *DescribeObjsByNsUIDRequest {

    return &DescribeObjsByNsUIDRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cmNameSpaces/{namespaceUID}/objs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: region(Required) */
func (r *DescribeObjsByNsUIDRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param namespaceUID: namespaceUID(Required) */
func (r *DescribeObjsByNsUIDRequest) SetNamespaceUID(namespaceUID string) {
    r.NamespaceUID = namespaceUID
}

/* param pageNumber: 当前所在页，默认为1(Optional) */
func (r *DescribeObjsByNsUIDRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 页面大小，默认为20；取值范围[1, 100](Optional) */
func (r *DescribeObjsByNsUIDRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeObjsByNsUIDRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeObjsByNsUIDResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeObjsByNsUIDResult `json:"result"`
}

type DescribeObjsByNsUIDResult struct {
    NumberPages int64 `json:"numberPages"`
    NumberRecords int64 `json:"numberRecords"`
    ObjList []monitor.ObjInfo `json:"objList"`
    PageNumber int64 `json:"pageNumber"`
    PageSize int64 `json:"pageSize"`
}