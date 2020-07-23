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
    sgw "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/sgw/models"
)

type DescribeUsersRequest struct {

    core.JDCloudRequest

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为10；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* pin (Optional) */
    Pin *string `json:"pin"`

    /* 用户类型 (Optional) */
    ConsumptionType *int `json:"consumptionType"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeUsersRequest(
) *DescribeUsersRequest {

	return &DescribeUsersRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/users",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为10；取值范围[10, 100] (Optional)
 * param pin: pin (Optional)
 * param consumptionType: 用户类型 (Optional)
 */
func NewDescribeUsersRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    pin *string,
    consumptionType *int,
) *DescribeUsersRequest {

    return &DescribeUsersRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/users",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
        Pin: pin,
        ConsumptionType: consumptionType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeUsersRequestWithoutParam() *DescribeUsersRequest {

    return &DescribeUsersRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/users",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeUsersRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为10；取值范围[10, 100](Optional) */
func (r *DescribeUsersRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param pin: pin(Optional) */
func (r *DescribeUsersRequest) SetPin(pin string) {
    r.Pin = &pin
}

/* param consumptionType: 用户类型(Optional) */
func (r *DescribeUsersRequest) SetConsumptionType(consumptionType int) {
    r.ConsumptionType = &consumptionType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeUsersRequest) GetRegionId() string {
    return ""
}

type DescribeUsersResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeUsersResult `json:"result"`
}

type DescribeUsersResult struct {
    Data sgw.User `json:"data"`
    Total int `json:"total"`
}