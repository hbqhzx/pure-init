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

type QueryWarnsServerWeakPasswordsRequest struct {

    core.JDCloudRequest

    /*  (Optional) */
    ServerId *string `json:"serverId"`

    /*  (Optional) */
    Interval *int `json:"interval"`

    /*  (Optional) */
    PageNumber *int `json:"pageNumber"`

    /*  (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryWarnsServerWeakPasswordsRequest(
) *QueryWarnsServerWeakPasswordsRequest {

	return &QueryWarnsServerWeakPasswordsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/warns:queryServerWeakPasswords",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param serverId:  (Optional)
 * param interval:  (Optional)
 * param pageNumber:  (Optional)
 * param pageSize:  (Optional)
 */
func NewQueryWarnsServerWeakPasswordsRequestWithAllParams(
    serverId *string,
    interval *int,
    pageNumber *int,
    pageSize *int,
) *QueryWarnsServerWeakPasswordsRequest {

    return &QueryWarnsServerWeakPasswordsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/warns:queryServerWeakPasswords",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        ServerId: serverId,
        Interval: interval,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryWarnsServerWeakPasswordsRequestWithoutParam() *QueryWarnsServerWeakPasswordsRequest {

    return &QueryWarnsServerWeakPasswordsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/warns:queryServerWeakPasswords",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param serverId: (Optional) */
func (r *QueryWarnsServerWeakPasswordsRequest) SetServerId(serverId string) {
    r.ServerId = &serverId
}

/* param interval: (Optional) */
func (r *QueryWarnsServerWeakPasswordsRequest) SetInterval(interval int) {
    r.Interval = &interval
}

/* param pageNumber: (Optional) */
func (r *QueryWarnsServerWeakPasswordsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: (Optional) */
func (r *QueryWarnsServerWeakPasswordsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryWarnsServerWeakPasswordsRequest) GetRegionId() string {
    return ""
}

type QueryWarnsServerWeakPasswordsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryWarnsServerWeakPasswordsResult `json:"result"`
}

type QueryWarnsServerWeakPasswordsResult struct {
    ServerWeakPasswords []hips.ServerWeakPassword `json:"serverWeakPasswords"`
    TotalCount int `json:"totalCount"`
}