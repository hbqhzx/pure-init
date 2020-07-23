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
    kmscap "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/kmscap/models"
)

type DescribeCustomerListRequest struct {

    core.JDCloudRequest

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为10；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeCustomerListRequest(
) *DescribeCustomerListRequest {

	return &DescribeCustomerListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/customer",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为10；取值范围[10, 100] (Optional)
 */
func NewDescribeCustomerListRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
) *DescribeCustomerListRequest {

    return &DescribeCustomerListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/customer",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeCustomerListRequestWithoutParam() *DescribeCustomerListRequest {

    return &DescribeCustomerListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/customer",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeCustomerListRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为10；取值范围[10, 100](Optional) */
func (r *DescribeCustomerListRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeCustomerListRequest) GetRegionId() string {
    return ""
}

type DescribeCustomerListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeCustomerListResult `json:"result"`
}

type DescribeCustomerListResult struct {
    CustomerList []kmscap.CustomerItem `json:"customerList"`
}