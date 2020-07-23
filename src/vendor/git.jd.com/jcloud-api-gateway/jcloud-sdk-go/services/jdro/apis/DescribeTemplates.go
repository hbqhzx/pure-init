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
    jdro "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/jdro/models"
)

type DescribeTemplatesRequest struct {

    core.JDCloudRequest

    /* 地域 ID  */
    RegionId string `json:"regionId"`

    /* 当前所在页，默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 页面大小，默认为20；取值范围[1, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 示例模板名称 (Optional) */
    Name *string `json:"name"`

    /* 模板类型，可取值(customer（用户），sample（示例）), 默认为sample (Optional) */
    Type *string `json:"type"`
}

/*
 * param regionId: 地域 ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeTemplatesRequest(
    regionId string,
) *DescribeTemplatesRequest {

	return &DescribeTemplatesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/templates",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域 ID (Required)
 * param pageNumber: 当前所在页，默认为1 (Optional)
 * param pageSize: 页面大小，默认为20；取值范围[1, 100] (Optional)
 * param name: 示例模板名称 (Optional)
 * param type_: 模板类型，可取值(customer（用户），sample（示例）), 默认为sample (Optional)
 */
func NewDescribeTemplatesRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    name *string,
    type_ *string,
) *DescribeTemplatesRequest {

    return &DescribeTemplatesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/templates",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Name: name,
        Type: type_,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeTemplatesRequestWithoutParam() *DescribeTemplatesRequest {

    return &DescribeTemplatesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/templates",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 ID(Required) */
func (r *DescribeTemplatesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 当前所在页，默认为1(Optional) */
func (r *DescribeTemplatesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 页面大小，默认为20；取值范围[1, 100](Optional) */
func (r *DescribeTemplatesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param name: 示例模板名称(Optional) */
func (r *DescribeTemplatesRequest) SetName(name string) {
    r.Name = &name
}

/* param type_: 模板类型，可取值(customer（用户），sample（示例）), 默认为sample(Optional) */
func (r *DescribeTemplatesRequest) SetType(type_ string) {
    r.Type = &type_
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeTemplatesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeTemplatesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeTemplatesResult `json:"result"`
}

type DescribeTemplatesResult struct {
    List []jdro.TemplateOut `json:"list"`
    TotalCount int64 `json:"totalCount"`
}